package utils

import (
	"base-system-backend/constants/errmsg"
	"base-system-backend/constants/table"
	"base-system-backend/global"
	"base-system-backend/model/privilege/response"
	"base-system-backend/utils/common"
	"encoding/json"
	"fmt"
)

func GetPrivilegeKeysByUserId(userID interface{}) (privilegeKeys []string, userRoleIds []int64, err error, debugInfo interface{}) {
	var (
		rolePrivilege []response.RoleIdPrivilege
		privilege     []string
	)
	err = global.DB.Table(table.Role).
		Select("privilege_keys", "role_id").
		Joins(fmt.Sprintf("LEFT join %s ur ON %s.id = ur.role_id", table.UserRole, table.Role)).
		Where("ur.user_id =?", userID).Find(&rolePrivilege).Error
	for _, value := range rolePrivilege {
		if err = json.Unmarshal(value.PrivilegeKeys, &privilege); err != nil {
			return nil, nil, fmt.Errorf("角色权限%w", errmsg.JsonConvertFiled), err.Error()
		}
		privilegeKeys = append(privilegeKeys, privilege...)
		userRoleIds = append(userRoleIds, value.RoleId)
	}
	//去重
	privilegeKeys = common.RemoveDuplication(privilegeKeys)
	return

}

func GetRolePrivilegeKeysByRoleId(roleId interface{}) (privilegeKeys []string, err error, debugInfo interface{}) {
	// 获取权限 key 列表
	var userPrivilegeKeys []string
	if err = global.DB.Table(table.Role).
		Select("privilege_keys").
		Where("id = ?", roleId).
		Find(&userPrivilegeKeys).Error; err != nil {
		return nil, fmt.Errorf("角色%w", errmsg.QueryFailed), err.Error()
	}
	// 合并
	if privilegeKeys, err, debugInfo = common.Merge(userPrivilegeKeys); err != nil {
		return nil, err, debugInfo
	}
	//去重
	privilegeKeys = common.RemoveDuplication(privilegeKeys)
	return
}

func RecursionGetChildPrivilege(privilege []*response.PrivilegeList, parentID int64) (res []*response.PrivilegeList) {
	for _, p := range privilege {
		if p.ParentId == 0 {
			continue
		}
		if p.ParentId == parentID {
			children := RecursionGetChildPrivilege(privilege, p.Id)
			p.ChildList = children
			res = append(res, p)
		}
	}
	return res
}
