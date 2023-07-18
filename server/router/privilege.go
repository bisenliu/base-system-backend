package router

import (
	v1 "base-system-backend/api/v1"
	"base-system-backend/constants"
	"base-system-backend/middleware"
	"github.com/gin-gonic/gin"
)

type PrivilegeRouter struct{}

func (PrivilegeRouter) InitPrivilegeRouter(Router *gin.RouterGroup) {
	privilegeApi := v1.ApiGroupApp.PrivilegeApi
	Router.GET("list/", middleware.PrivilegeVerify(constants.PrivilegeList), privilegeApi.PrivilegeListApi)
	Router.PUT(":role_id/", middleware.PrivilegeVerify(constants.PrivilegeSet), privilegeApi.RolePrivilegeUpdateApi)
}