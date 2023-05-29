package router

import (
	v1 "base-system-backend/api/v1"
	"github.com/gin-gonic/gin"
)

type UserRouter struct{}

func (u UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	userApi := v1.ApiGroupApp.UserApi
	// 登录
	Router.POST("login/", userApi.UserLoginApi)
	// 登出
	Router.POST("logout/", userApi.UserLogoutApi)

	// 用户列表
	Router.GET("list/", userApi.UserListApi)
	//创建用户
	Router.POST("create/", userApi.UserCreateApi)
}
