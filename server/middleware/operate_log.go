package middleware

import (
	"base-system-backend/utils"
	"github.com/gin-gonic/gin"
)

// OperateLogMiddleware
//
//	@Description: 操作日志中间件
//	@return func(c *gin.Context) 上下文信息
func OperateLogMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		success, statusInfo := utils.GetResponseData(c)
		utils.CreateOperateLog(c, success, statusInfo)
	}
}
