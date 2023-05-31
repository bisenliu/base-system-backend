package router

type GroupRouter struct {
	User UserRouter
	//Privilege PrivilegeRouter
	Role RoleRouter
	Log  LogRouter
}

var RouterGroupApp = new(GroupRouter)
