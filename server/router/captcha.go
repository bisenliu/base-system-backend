package router

import (
	v1 "base-system-backend/api/v1"
	"github.com/gin-gonic/gin"
)

type CaptchaRouter struct{}

// InitCaptchaRouter
//  @Description: 滑块路由
//  @receiver CaptchaRouter
//  @param Router routerGroup对象

func (CaptchaRouter) InitCaptchaRouter(Router *gin.RouterGroup) {
	captchaApi := v1.ApiGroupApp.CaptchaApi
	Router.POST("/get/", captchaApi.CaptchaGetApi)
	Router.POST("/check/", captchaApi.CaptchaCheckApi)
}
