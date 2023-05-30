package service

type ServicesGroup struct {
	UserService
	//PrivilegeService
	RoleService
	//LogService
}

var ServicesGroupApp = new(ServicesGroup)
