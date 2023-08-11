package middleware

import (
	"base-system-backend/constants/code"
	"base-system-backend/constants/errmsg"
	"base-system-backend/constants/user"
	"base-system-backend/model/common/response"
	"base-system-backend/utils"
	"base-system-backend/utils/cache"
	"base-system-backend/utils/jwt"
	"fmt"
	"github.com/gin-gonic/gin"
)

var notNeedAuthPath = []string{
	"/v1/common/version/",
	"/v1/user/login/",
}

// JWTAuthMiddleware
//
//	@Description: 基于JWT的认证中间件
//	@return func(c *gin.Context) 上下文信息
func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {

		statusInfo := map[string]interface{}{
			"message": errmsg.LoginInvalid.Error(),
		}

		// 获取请求uri
		requestURI := c.Request.RequestURI

		// 检查请求是否不需要认证
		for _, path := range notNeedAuthPath {
			if requestURI == path {
				c.Next()
				return
			}
		}
		authHeader := c.Request.Header.Get("Identification")
		if authHeader == "" {
			// todo 如果没有 Identification，先校验ak/sk
			debugInfo := fmt.Errorf("token%w", errmsg.Required).Error()
			statusInfo["debug_info"] = debugInfo
			utils.CreateOperateLog(c, false, statusInfo)
			response.Error(c, code.InvalidLogin, errmsg.LoginInvalid, debugInfo)
			c.Abort()
			return
		}
		// parts[1]是获取到的tokenString，我们使用之前定义好的解析JWT的函数来解析它
		mc, err := jwt.ParseToken(authHeader)
		if err != nil {
			debugInfo := fmt.Errorf("token%w", errmsg.ParseFailed).Error()
			statusInfo["debug_info"] = debugInfo
			utils.CreateOperateLog(c, false, statusInfo)
			response.Error(c, code.InvalidLogin, errmsg.LoginInvalid, debugInfo)
			c.Abort()
			return

		}
		// 获取redis token
		cacheToken := cache.GetToken(mc.UserId)
		if cacheToken != authHeader {
			debugInfo := fmt.Errorf("token%w", errmsg.Invalid).Error()
			statusInfo["debug_info"] = debugInfo
			utils.CreateOperateLog(c, false, statusInfo)
			response.Error(c, code.InvalidLogin, errmsg.LoginInvalid, debugInfo)
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
