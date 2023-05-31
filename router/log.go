package router

import (
	v1 "base-system-backend/api/v1"
	"github.com/gin-gonic/gin"
)

type LogRouter struct{}

func (LogRouter) InitLogRouter(Router *gin.RouterGroup) {
	logApi := v1.ApiGroupApp.LogApi
	Router.GET("operate/list/", logApi.OperateLogListApi)
	//Router.GET("operate/download/", logApi.OperateLogDownloadApi)
}
