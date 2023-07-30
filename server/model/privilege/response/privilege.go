package response

import (
	"base-system-backend/model/privilege"
	"gorm.io/datatypes"
)

type PrivilegeList struct {
	privilege.Privilege
	ChildList []*PrivilegeList `json:"child_list" gorm:"-"`
}

type RoleIdPrivilege struct {
	PrivilegeKeys datatypes.JSON `json:"privilege_keys"`
	RoleId        int64          `json:"role_id"`
}
