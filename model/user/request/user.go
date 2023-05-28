package request

import (
	"base-system-backend/enums/login"
	"base-system-backend/enums/user"
	"base-system-backend/model/common/request"
)

// UserFilter
//
//	@Description: 用户列表过滤参数
type UserFilter struct {
	Name   string         `json:"name" form:"name" binding:"" label:"用户名"`
	Status user.AccStatus `json:"status" form:"status" binding:"enum" label:"用户状态"`
	request.PageInfo
}

// UserLoginBase
//
//	@Description: 登录请求参数
type UserLoginBase struct {
	LoginType *login.LoginType `json:"login_type"  binding:"required,enum" label:"登录类型"`
	Account   *string          `json:"account"  binding:"" label:"账号"`
	Password  *string          `json:"password"  binding:"" label:"密码"`
	Phone     *string          `json:"phone"  binding:"" label:"手机号"`
	Code      *string          `json:"code"  binding:"" label:"验证码"`
	Slider    *struct{}        `json:"slider"  binding:"" label:"滑块轨迹信息"`
}

// UserAccountLogin
//
//	@Description: 账号密码登录请求参数
type UserAccountLogin struct {
	Account  string `json:"account"  binding:"required,max=20" label:"账号"`
	Password string `json:"password"  binding:"required" label:"密码"`
}
