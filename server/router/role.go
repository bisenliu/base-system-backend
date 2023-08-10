package router

import (
	v1 "base-system-backend/api/v1"
	"base-system-backend/constants"
	"base-system-backend/middleware"
	"github.com/gin-gonic/gin"
)

type RoleRouter struct{}

// InitRoleRouter
//  @Description: 角色路由
//  @receiver RoleRouter
//  @param Router routerGroup对象

func (RoleRouter) InitRoleRouter(Router *gin.RouterGroup) {
	roleApi := v1.ApiGroupApp.RoleApi
	// 角色创建
	Router.POST("", middleware.PrivilegeVerify(constants.RoleCreate), roleApi.RoleCreateApi)
	// 角色列表
	Router.GET("list/", roleApi.RoleListApi)
	// 角色CRUD
	roleCRUDRouterGroup := Router.Group(":role_id/")
	{
		// 角色详情
		roleCRUDRouterGroup.GET("", roleApi.RoleDetailApi)
		// 角色修改
		roleCRUDRouterGroup.PUT("", middleware.PrivilegeVerify(constants.RoleUpdate), roleApi.RoleUpdateApi)
		// 角色删除
		roleCRUDRouterGroup.DELETE("", middleware.PrivilegeVerify(constants.RoleDelete), roleApi.RoleDeleteApi)
	}
}
