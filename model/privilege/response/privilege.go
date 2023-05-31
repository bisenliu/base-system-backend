package response

import "base-system-backend/model/privilege"

type PrivilegeList struct {
	privilege.Privilege
	ChildList []*PrivilegeList `json:"child_list" gorm:"-"`
}
