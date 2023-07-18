package utils

import (
	"base-system-backend/constants/errmsg"
	"base-system-backend/constants/table"
	"base-system-backend/global"
	"fmt"
)

func GetUserRoleIdsByUserId(userId int64) (userRoleIds []int64, err error, debugInfo interface{}) {
	if err = global.DB.Table(table.UserRole).
		Select("role_id").Where("user_id = ?", userId).Find(&userRoleIds).Error; err != nil {
		return nil, fmt.Errorf("用户角色%w", errmsg.QueryFailed), err.Error()
	}
	return
}
