package router

import (
	v1 "base-system-backend/api/v1"
	"github.com/gin-gonic/gin"
)

type RoleRouter struct{}

func (RoleRouter) InitRoleRouter(Router *gin.RouterGroup) {
	roleApi := v1.ApiGroupApp.RoleApi
	// 角色列表
	Router.GET("list/", roleApi.RoleListApi)
	//// 角色创建
	//Router.POST("", roleApi.RoleCreate)
	//// 角色CRUD
	//roleCRUDRouterGroup := Router.Group(":role_id/")
	//{
	//	// 角色详情
	//	roleCRUDRouterGroup.GET("", roleApi.RoleDetail)
	//	// 角色修改
	//	roleCRUDRouterGroup.PUT("", roleApi.RoleUpdate)
	//	// 角色删除
	//	roleCRUDRouterGroup.DELETE("", roleApi.RoleDelete)
	//}
}
