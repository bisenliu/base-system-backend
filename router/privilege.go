package router

import (
	v1 "base-system-backend/api/v1"
	"github.com/gin-gonic/gin"
)

type PrivilegeRouter struct{}

func (PrivilegeRouter) InitPrivilegeRouter(Router *gin.RouterGroup) {
	privilegeApi := v1.ApiGroupApp.PrivilegeApi
	Router.GET("list/", privilegeApi.PrivilegeListApi)
	Router.PUT(":role_id/", privilegeApi.RolePrivilegeUpdateApi)
}
