package router

import (
	v1 "base-system-backend/api/v1"
	"github.com/gin-gonic/gin"
)

type RoleRouter struct{}

func (RoleRouter) InitRoleRouter(Router *gin.RouterGroup) {
	roleApi := v1.ApiGroupApp.RoleApi
	// 角色创建
	Router.POST("", roleApi.RoleCreateApi)
	// 角色列表
	Router.GET("list/", roleApi.RoleListApi)
	// 角色CRUD
	roleCRUDRouterGroup := Router.Group(":role_id/")
	{
		// 角色详情
		roleCRUDRouterGroup.GET("", roleApi.RoleDetailApi)
		// 角色修改
		roleCRUDRouterGroup.PUT("", roleApi.RoleUpdate)
		//// 角色删除
		//roleCRUDRouterGroup.DELETE("", roleApi.RoleDelete)
	}
}