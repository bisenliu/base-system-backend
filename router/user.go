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
	// 创建用户
	Router.POST("create/", userApi.UserCreateApi)
	// 查询|修改当前登录用户信息
	detailRouterGroup := Router.Group("detail/")
	{
		detailRouterGroup.GET("", userApi.UserDetailApi)
		detailRouterGroup.PUT("", userApi.UserUpdateApi)
	}
	// 修改当前登录用户密码
	Router.PATCH("change_pwd/", userApi.UserChangePwdApi)
	//上传头像
	Router.PATCH("avatar/", userApi.UserUploadAvatarApi)
	//重置指定账号密码
	Router.PUT("reset_pwd/:user_id/", userApi.UserResetPwdByIdApi)
	// 修改指定账户状态
	Router.PUT("change_status/:user_id/", userApi.UserStatusChangeByIdApi)
	// 查询|编辑指定账户信息
	RURouterGroup := Router.Group(":user_id/")
	{
		RURouterGroup.GET("", userApi.UserDetailByIdApi)
		//RURouterGroup.PUT("", userApi.UserUpdateById)
	}

}
