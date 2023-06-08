package middleware

import (
	"base-system-backend/enums/code"
	"base-system-backend/enums/errmsg"
	"base-system-backend/enums/user"
	"base-system-backend/model/common/response"
	"base-system-backend/utils/cache"
	"base-system-backend/utils/jwt"
	"fmt"
	"github.com/gin-gonic/gin"
	"reflect"
	"runtime"
)

// JWTAuthMiddleware 基于JWT的认证中间件
func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		// 获取当前请求的处理函数
		handler := c.Handler()

		// 使用反射获取函数名称
		funcName := runtime.FuncForPC(reflect.ValueOf(handler).Pointer()).Name()

		// 获取函数的注释
		comments := getFunctionComments(funcName)

		// 打印注释
		fmt.Println("API 注释：", comments)

		// 获取请求uri
		requestURI := c.Request.RequestURI
		// 不需要认证的接口
		notNeedAuthPath := []string{
			"/v1/common/version/",
			"/v1/user/login/",
			"/v1/user/slider_verify/",
		}
		for _, path := range notNeedAuthPath {
			if requestURI == path {
				c.Next()
				return
			}
		}

		authHeader := c.Request.Header.Get("Identification")
		if authHeader == "" {
			// todo 如果没有 Identification，先校验ak/sk
			response.Error(c, code.InvalidLogin, errmsg.LoginInvalid, "token not found")
			c.Abort()
			return
		}
		// parts[1]是获取到的tokenString，我们使用之前定义好的解析JWT的函数来解析它
		mc, err := jwt.ParseToken(authHeader)
		if err != nil {
			response.Error(c, code.InvalidLogin, errmsg.LoginInvalid, "token parse failed")
			c.Abort()
			return
		}
		// 获取redis token
		cacheToken := cache.GetToken(mc.UserId)
		if cacheToken == "" {
			response.Error(c, code.InvalidLogin, errmsg.LoginInvalid, "cache token not found")
			c.Abort()
			return
		}
		if cacheToken != authHeader {
			response.Error(c, code.InvalidLogin, errmsg.LoginInvalid, "invalid token")
			c.Abort()
			return
		}
		// 将当前请求的username信息保存到请求的上下文c上
		c.Set(user.CtxUserIdKey, mc.UserId)
		// 刷新token
		cache.FlushToken(mc.UserId)
		// 后续的处理函数可以用过c.Get(user.CtxUserIdKey)来获取当前请求的用户信息
		c.Next()
	}
}
