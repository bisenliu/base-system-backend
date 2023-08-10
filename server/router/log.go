package router

import (
	v1 "base-system-backend/api/v1"
	"base-system-backend/constants"
	"base-system-backend/middleware"
	"github.com/gin-gonic/gin"
)

type LogRouter struct{}

// InitLogRouter
//  @Description: 操作日志路由
//  @receiver LogRouter
//  @param Router routerGroup对象

func (LogRouter) InitLogRouter(Router *gin.RouterGroup) {
	logApi := v1.ApiGroupApp.LogApi
	Router.GET("operate/list/", middleware.PrivilegeVerify(constants.OperateLogList), logApi.OperateLogListApi)
	Router.GET("operate/download/", middleware.PrivilegeVerify(constants.OperateLogDownload), logApi.OperateLogDownloadApi)
}
