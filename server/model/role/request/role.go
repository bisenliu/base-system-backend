package request

type RoleCreate struct {
	Id   int64  `json:"id"`
	Name string `json:"name" binding:"required,max=20" label:"角色名"`
}

type RoleUpdate struct {
	Name          string   `json:"name" binding:"required,max=20" label:"角色名"`
	PrivilegeKeys []string `json:"privilege_keys" binding:"" label:"角色权限key列表"`
}

// RoleFilter
//
//	@Description: 角色列表过滤参数
type RoleFilter struct {
	Name string `json:"name" form:"name" binding:"" label:"角色名"`
}
