package router

import (
	v1 "base-system-backend/api/v1"
	"base-system-backend/enums"
	"base-system-backend/middleware"
	"github.com/gin-gonic/gin"
)

type LogRouter struct{}

func (LogRouter) InitLogRouter(Router *gin.RouterGroup) {
	logApi := v1.ApiGroupApp.LogApi
	Router.GET("operate/list/", middleware.PrivilegeVerify(enums.OperateLogList), logApi.OperateLogListApi)
	Router.GET("operate/download/", middleware.PrivilegeVerify(enums.OperateLogDownload), logApi.OperateLogDownloadApi)
}
