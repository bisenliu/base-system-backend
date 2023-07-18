package request

type RolePrivilegeUpdate struct {
	PrivilegeKeys []string `json:"privilege_keys" binding:"required" label:"权限key列表"`
}
