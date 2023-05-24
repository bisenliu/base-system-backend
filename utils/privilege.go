package utils

import (
	"base-system-backend/enums/errmsg"
	"base-system-backend/enums/table"
	"base-system-backend/global"
	"base-system-backend/utils/common"
	"fmt"
)

func GetPrivilegeKeysByUserId(userID int64) (privilegeKeys []string, userRoleIds []int64, err error, debugInfo interface{}) {
	// 获取用户角色 ID 列表
	userRoleIds, err, debugInfo = GetUserRoleIdsByUserId(userID)
	if err != nil {
		return nil, nil, err, debugInfo
	}
	// 获取角色权限 key 列表(根据角色 Id 列表)
	privilegeKeys, err, debugInfo = GetRolePrivilegeKeysByRoleId(userRoleIds)
	if err != nil {
		return nil, nil, err, debugInfo
	}
	return
}

func GetRolePrivilegeKeysByRoleId(userRoleIds []int64) (privilegeKeys []string, err error, debugInfo interface{}) {
	// 获取权限 key 列表
	var userPrivilegeKeys []string
	if err = global.DB.Table(table.Role).
		Select("privilege_keys").
		Where("id in ?", userRoleIds).
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
