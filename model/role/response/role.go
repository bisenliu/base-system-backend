package response

import (
	"base-system-backend/model/role"
)

//
//  RoleList
//  @Description: 角色列表

type RoleList struct {
	Results []role.Role `json:"results" form:"results"`
}
