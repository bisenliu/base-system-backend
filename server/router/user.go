package router

import (
	v1 "base-system-backend/api/v1"
	"base-system-backend/constants"
	"base-system-backend/middleware"
	"github.com/gin-gonic/gin"
)

type UserRouter struct{}

// InitUserRouter
//  @Description: 用户路由
//  @receiver u
//  @param Router routerGroup对象

func (UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	userApi := v1.ApiGroupApp.UserApi
	// 创建用户
	Router.POST("", middleware.PrivilegeVerify(constants.AccountCreate), userApi.UserCreateApi)
	// 登陆
	Router.POST("login/", userApi.UserLoginApi)
	// 登出
	Router.POST("logout/", userApi.UserLogoutApi)

	// 用户列表
	Router.GET("list/", middleware.PrivilegeVerify(constants.AccountList), userApi.UserListApi)
	// 查询|修改当前登陆用户信息
	detailRouterGroup := Router.Group("detail/")
	{
		detailRouterGroup.GET("", userApi.UserDetailApi)
		detailRouterGroup.PUT("", userApi.UserUpdateApi)
	}
	// 修改当前登陆用户密码
	Router.PATCH("password/", userApi.UserChangePwdApi)
	//上传头像
	Router.PATCH("avatar/", userApi.UserUploadAvatarApi)
	//重置指定账号密码
	Router.PUT(":user_id/password/", middleware.PrivilegeVerify(constants.ResetPwdOther), userApi.UserResetPwdByIdApi)
	// 修改指定账户状态
	Router.PUT(":user_id/status/", middleware.PrivilegeVerify(constants.ChangeStatusOther), userApi.UserStatusChangeByIdApi)
	// 查询|编辑指定账户信息
	RURouterGroup := Router.Group(":user_id/")
	{
		RURouterGroup.GET("", middleware.PrivilegeVerify(constants.AccountDetailOther), userApi.UserDetailByIdApi)
		RURouterGroup.PUT("", middleware.PrivilegeVerify(constants.AccountUpdateOther), userApi.UserUpdateByIdApi)
	}

}
