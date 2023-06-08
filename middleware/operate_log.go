package middleware

import (
	"base-system-backend/utils"
	"github.com/gin-gonic/gin"
)

func OperateLogMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		success, detailByte := utils.GetResponseData(c)
		utils.CreateOperateLog(c, success, detailByte)
	}
}
