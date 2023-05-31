package service

import (
	"base-system-backend/enums"
	"base-system-backend/enums/errmsg"
	"base-system-backend/enums/table"
	"base-system-backend/global"
	"base-system-backend/model/privilege/request"
	"base-system-backend/model/privilege/response"
	"base-system-backend/model/role"
	"base-system-backend/utils"
	"encoding/json"
	"fmt"
)

type PrivilegeService struct{}

func (PrivilegeService) PrivilegeListService() (res []*response.PrivilegeList, err error, debugInfo interface{}) {
	var privileges []*response.PrivilegeList
	if err = global.DB.Table(table.Privilege).Find(&privileges).Error; err != nil {
		return nil, fmt.Errorf("权限信息%w", errmsg.QueryFailed), err.Error()
	}
	res = utils.RecursionGetChildPrivilege(privileges, 1)
	return
}

func (PrivilegeService) RolePrivilegeUpdateService(roleId string, params *request.RolePrivilegeUpdate) (err error, debugInfo interface{}) {
	r := new(role.Role)
	if err = global.DB.Table(table.Role).Where("id =?", roleId).First(&r).Error; err != nil {
		return fmt.Errorf("角色%w", errmsg.QueryFailed), err.Error()
	}
	//不能更新系统角色
	if r.IsSystem == enums.True {
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
