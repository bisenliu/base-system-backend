package request

type RoleCreate struct {
	Id   int64  `json:"id"`
	Name string `json:"name" binding:"required,max=20" label:"角色名"`
}

type RoleUpdate struct {
	Name          string   `json:"name" binding:"required,max=20" label:"角色名"`
	PrivilegeKeys []string `json:"privilege_keys" binding:"" label:"角色权限key列表"`
}
