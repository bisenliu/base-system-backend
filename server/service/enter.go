package service

type ServicesGroup struct {
	UserService
	PrivilegeService
	RoleService
	LogService
	CaptchaService
}

var ServicesGroupApp = new(ServicesGroup)
