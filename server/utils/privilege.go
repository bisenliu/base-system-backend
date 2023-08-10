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

// GetPrivilegeKeysByUserId
//  @Description: 获取用户拥有权限
//  @param userID 用户 ID
//  @return privilegeKeys 权限 key slice
//  @return userRoleIds 角色 ID slice
//  @return err 查询/序列化失败异常
//  @return debugInfo 错误调试信息

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

// GetRolePrivilegeKeysByRoleId
//  @Description: 获取角色拥有权限
//  @param roleId 角色 ID
//  @return privilegeKeys 角色权限 slice
//  @return err 查询失败异常
//  @return debugInfo 错误调试信息

func GetRolePrivilegeKeysByRoleId(roleId interface{}) (privilegeKeys []string, err error, debugInfo interface{}) {
	// 获取权限 key 列表
	var userPrivilegeKeys []string
	if err = global.DB.Table(table.Role).
		Select("privilege_keys").
		Where("id = ?", roleId).
		Find(&userPrivilegeKeys).Error; err != nil {
		return nil, fmt.Errorf("角色%w", errmsg.QueryFailed), err.Error()
	}
	return
}

// RecursionGetChildPrivilege
//  @Description: 递归组装权限信息
//  @param privilege 权限 slice
//  @param parentID 父级 ID
//  @return res 组装完成的权限 slice

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
