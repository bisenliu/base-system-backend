package v1

import "base-system-backend/service"

type ApiGroup struct {
	UserApi UserApi
	//PrivilegeApi PrivilegeApi
	RoleApi RoleApi
	LogApi  LogApi
}

var ApiGroupApp = new(ApiGroup)

var (
	// 用户模块服务层入口
	userService = service.ServicesGroupApp.UserService
	//// 权限模块服务层入口
	//privilegeService=service.ServicesGroupApp.PrivilegeService
	// 角色模块服务层入口
	roleService = service.ServicesGroupApp.RoleService
	// 日志模块服务层入口
	logService = service.ServicesGroupApp.LogService
)
