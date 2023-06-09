package middleware

import (
	"base-system-backend/enums"
	"base-system-backend/enums/code"
	"base-system-backend/enums/errmsg"
	"base-system-backend/enums/user"
	"base-system-backend/model/common/response"
	"base-system-backend/utils"
	"base-system-backend/utils/cache"
	"base-system-backend/utils/jwt"
	"encoding/json"
	"github.com/gin-gonic/gin"
)

// JWTAuthMiddleware 基于JWT的认证中间件
func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {

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
		var debugInfo string
		authHeader := c.Request.Header.Get("Identification")
		if authHeader == "" {
			// todo 如果没有 Identification，先校验ak/sk
			debugInfo = "token not found"
		}
		// parts[1]是获取到的tokenString，我们使用之前定义好的解析JWT的函数来解析它
		mc, err := jwt.ParseToken(authHeader)
		if err != nil {
			debugInfo = "token parse failed"

		}
		// 获取redis token
		cacheToken := cache.GetToken(mc.UserId)
		if cacheToken == "" {
			debugInfo = "cache token not found"
		}
		if cacheToken != authHeader {
			debugInfo = "invalid token"
		}
		if debugInfo != "" {
			response.Error(c, code.InvalidLogin, errmsg.LoginInvalid, debugInfo)
			detailByte, _ := json.Marshal(map[string]string{"message": errmsg.LoginInvalid.Error()})
			utils.CreateOperateLog(c, enums.False, detailByte)
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
