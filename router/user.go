package router

import (
	v1 "base-system-backend/api/v1"
	"base-system-backend/enums"
	"base-system-backend/middleware"
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
	Router.GET("list/", middleware.PrivilegeVerify(enums.AccountList), userApi.UserListApi)
	// 创建用户
	Router.POST("create/", middleware.PrivilegeVerify(enums.AccountCreate), userApi.UserCreateApi)
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
	Router.PUT("reset_pwd/:user_id/", middleware.PrivilegeVerify(enums.ResetPwdOther), userApi.UserResetPwdByIdApi)
	// 修改指定账户状态
	Router.PUT("change_status/:user_id/", middleware.PrivilegeVerify(enums.ChangeStatusOther), userApi.UserStatusChangeByIdApi)
	// 查询|编辑指定账户信息
	RURouterGroup := Router.Group(":user_id/")
	{
		RURouterGroup.GET("", middleware.PrivilegeVerify(enums.AccountDetailOther), userApi.UserDetailByIdApi)
		RURouterGroup.PUT("", middleware.PrivilegeVerify(enums.AccountUpdateOther), userApi.UserUpdateByIdApi)
	}

}
