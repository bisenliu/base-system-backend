package request

import (
	"base-system-backend/enums/gender"
	"base-system-backend/enums/login"
	"base-system-backend/enums/user"
	"base-system-backend/model/common/field"
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

type UserCreate struct {
	Id        int64          `json:"id"`
	Account   string         `json:"account" binding:"required,max=20" label:"账号"`
	Password  string         `json:"password" binding:"required" label:"密码"`
	SecretKey string         `json:"secret_key" label:"API秘钥"`
	Phone     string         `json:"phone" binding:"max=11" label:"手机号"`
	Email     string         `json:"email" binding:"omitempty,email" label:"邮箱"`
	Name      string         `json:"name" binding:"max=20" label:"姓名"`
	IdCard    string         `json:"id_card" binding:"max=18" label:"身份证号码"`
	Gender    gender.Gender  `json:"gender" binding:"enum" label:"性别"`
	Status    user.AccStatus `json:"status" binding:"enum" label:"账号状态"`
	RoleIds   *[]int64       `json:"role_ids" gorm:"-" label:"角色ID列表"`
	UserFullNameAndShortName
	field.CUTime
}

type UserFullNameAndShortName struct {
	FullName  string `json:"full_name" binding:"" label:"姓名全拼"`
	ShortName string `json:"short_name" binding:"" label:"姓名简拼"`
}

type UserUpdate struct {
	Phone string `json:"phone" binding:"max=11" label:"手机号"`
	UserUpdateById
}

type UserUpdateById struct {
	Name    string         `json:"name" binding:"required,max=20" label:"姓名"`
	IdCard  string         `json:"id_card" binding:"max=18" label:"身份证号码"`
	Email   string         `json:"email" binding:"omitempty,email" label:"邮箱"`
	Gender  gender.Gender  `json:"gender" binding:"enum" label:"性别"`
	Status  user.AccStatus `json:"status" binding:"enum" label:"账号状态"`
	RoleIds *[]int64       `json:"role_ids" gorm:"-" label:"角色ID列表"`
	UserFullNameAndShortName
}

type UserChangePwdBase struct {
	Type        user.ChangePwdType `json:"type" binding:"enum" label:"修改密码类型"`
	OldPassword *string            `json:"old_password" binding:"" label:"旧密码"`
	NewPassword *string            `json:"new_password" binding:"" label:"新密码"`
}

type PwdChangeByPwd struct {
	OldPassword string `json:"old_password" binding:"required,max=70" label:"旧密码"`
	NewPassword string `json:"new_password" binding:"required,max=70" label:"新密码"`
}

type PwdChangeById struct {
	Password string `json:"password" binding:"required,max=70" label:"密码"`
}

type StatusChangeById struct {
	Status user.AccStatus `json:"status" binding:"enum" label:"账号状态"`
}
