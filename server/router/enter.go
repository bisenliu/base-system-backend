package router

type GroupRouter struct {
	User          UserRouter
	Privilege     PrivilegeRouter
	Role          RoleRouter
	Log           LogRouter
	CaptchaRouter CaptchaRouter
}

var RouterGroupApp = new(GroupRouter)
