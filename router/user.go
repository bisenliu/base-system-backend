package router

import (
	v1 "base-system-backend/api/v1"
	"github.com/gin-gonic/gin"
)

type UserRouter struct{}

func (u UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	userApi := v1.ApiGroupApp.UserApi
	//创建用户

	// 用户列表
	Router.GET("list/", userApi.UserListApi)
	// 登录
	Router.POST("login/", userApi.UserLoginApi)
	Router.POST("logout/", userApi.UserLogoutApi)
}
