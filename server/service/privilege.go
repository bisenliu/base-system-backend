package service

import (
	"base-system-backend/constants/errmsg"
	"base-system-backend/constants/table"
	"base-system-backend/global"
	"base-system-backend/model/privilege/request"
	"base-system-backend/model/privilege/response"
	"base-system-backend/model/role"
	"base-system-backend/utils"
	"base-system-backend/utils/common"
	"encoding/json"
	"fmt"
	"gorm.io/datatypes"
)

type PrivilegeService struct{}

// PrivilegeListService
//  @Description: 权限列表 service
//	@receiver PrivilegeService
//  @param userId 用户 ID
//  @param roleId 角色 ID
//  @return res 权限列表
//  @return err 查询/序列化失败异常
//  @return debugInfo 错误调试信息

func (PrivilegeService) PrivilegeListService(userId, roleId string) (res interface{}, err error, debugInfo interface{}) {
	var (
		data          []string
		privilegeKey  []string
		privilegeKeys []datatypes.JSON
		privilegeList []*response.PrivilegeList
	)

	switch {
	//合并查询
	case userId != "" && roleId != "":
		if err = global.DB.Debug().Table(table.Role).
			Joins(fmt.Sprintf("LEFT join %s ur ON %s.id = ur.role_id", table.UserRole, table.Role)).
			Where("ur.user_id = ?", userId).
			Where("ur.role_id = ?", roleId).
			Pluck("privilege_keys", &privilegeKeys).Error; err != nil {
			return
		}
		for _, value := range privilegeKeys {
			if err = json.Unmarshal(value, &data); err != nil {
				return nil, fmt.Errorf("角色权限%w", errmsg.JsonConvertFiled), err.Error()
			}
			privilegeKey = append(privilegeKey, data...)
		}
		res = common.RemoveDuplication(privilegeKey)
		return
		// userId 查询
	case userId != "":
		res, _, err, debugInfo = utils.GetPrivilegeKeysByUserId(userId)
		if err != nil {
			return nil, err, debugInfo
		}
		return
		// roleId 查询
	case roleId != "":
		res, err, debugInfo = utils.GetRolePrivilegeKeysByRoleId(roleId)
		if err != nil {
			return nil, err, debugInfo
		}
		return
		//无查询条件
	default:

		if err = global.DB.Table(table.Privilege).Find(&privilegeList).Error; err != nil {
			return nil, fmt.Errorf("权限信息%w", errmsg.QueryFailed), err.Error()
		}
	}
	res = utils.RecursionGetChildPrivilege(privilegeList, 1)
	return
}

// RolePrivilegeUpdateService
//  @Description: 更新角色权限 service
//	@receiver PrivilegeService
//  @param roleId 角色 ID
//  @param params 权限信息
//  @return err 权限key无效/查询/序列化失败异常
//  @return debugInfo 错误调试信息

func (PrivilegeService) RolePrivilegeUpdateService(roleId string, params *request.RolePrivilegeUpdate) (err error, debugInfo interface{}) {
	r := new(role.Role)
	if err = global.DB.Table(table.Role).Where("id =?", roleId).First(&r).Error; err != nil {
		return fmt.Errorf("角色%w", errmsg.QueryFailed), err.Error()
	}
	//不能更新系统角色
	if r.IsSystem == true {
		return fmt.Errorf(errmsg.NotPrivilege.Error(), "修改系统角色权限Key"), nil
	}
	//过滤无效的权限key
	var privilegeKeys []string
	if err = global.DB.Table(table.Privilege).
		Select("key").Where("key in ?", params.PrivilegeKeys).Find(&privilegeKeys).Error; err != nil {
		return fmt.Errorf("权限key%w", errmsg.QueryFailed), err.Error()
	}
	marshal, err := json.Marshal(privilegeKeys)
	if err != nil {
		return errmsg.JsonConvertFiled, err.Error()
	}
	if err = global.DB.Model(&r).Update("privilege_keys", marshal).Error; err != nil {
		return fmt.Errorf("角色权限%w", errmsg.UpdateFailed), err.Error()
	}
	return
}
