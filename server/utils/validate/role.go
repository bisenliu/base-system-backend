package validate

import (
	"base-system-backend/constants/errmsg"
	"base-system-backend/constants/table"
	"base-system-backend/global"
	"base-system-backend/model/user"
	"fmt"
)

// RoleIdsVerify
//  @Description: 角色 ID Slice 校验
//  @param roleIds 角色 ID Slice
//  @return err 查询失败异常
//  @return debugInfo 错误调试信息

func RoleIdsVerify(roleIds *[]int64) (err error, debugInfo interface{}) {
	var roleIdsCount int64
	if err = global.DB.Table(table.Role).
		Where("id in ?", *roleIds).Count(&roleIdsCount).Error; err != nil {
		return fmt.Errorf("角色%w", errmsg.QueryFailed), err.Error()
	}
	if int64(len(*roleIds)) != roleIdsCount {

		return fmt.Errorf("角色ID%w", errmsg.Invalid), nil
	}
	return
}

// BindRoleVerify
//  @Description: 用户绑定角色校验
//  @param userId 用户 ID
//  @param roleIds 角色 ID Slice
//  @return userRoles 绑定实例 Slice
//  @return err 角色查询失败异常
//  @return debugInfo 错误调试信息

func BindRoleVerify(userId int64, roleIds *[]int64) (userRoles []user.UserRole, err error, debugInfo interface{}) {
	if err, debugInfo = RoleIdsVerify(roleIds); err != nil {
		return nil, err, debugInfo
	}
	for _, roleId := range *roleIds {
		userRoles = append(userRoles, user.UserRole{
			UserId: userId,
			RoleId: roleId,
		})
	}
	return
}
