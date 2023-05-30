package service

import (
	"base-system-backend/enums"
	"base-system-backend/enums/errmsg"
	"base-system-backend/enums/table"
	"base-system-backend/global"
	"base-system-backend/model/role"
	"base-system-backend/model/role/request"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
)

type RoleService struct{}

func (RoleService) RoleListService(c *gin.Context) (roleList []role.Role, err error, debugInfo interface{}) {
	roleName := c.Query("name")
	tx := global.DB.Table(table.Role)
	if roleName != "" {
		tx = tx.Where("name LIKE ?", fmt.Sprintf("%%%s%%", roleName))
	}
	if err = tx.Order("id").Find(&roleList).Error; err != nil {
		return nil, fmt.Errorf("角色%w", errmsg.QueryFailed), err.Error()
	}
	return
}

func (RoleService) RoleCreateService(params *request.RoleCreate) (err error, debugInfo interface{}) {
	if err = global.DB.Table(table.Role).Create(params).Error; err != nil {
		return fmt.Errorf("角色%w", errmsg.SaveFailed), err.Error()
	}
	return
}

func (RoleService) RoleDetailService(roleId string) (roleDetail *role.Role, err error, debugInfo interface{}) {
	if err = global.DB.Table(table.Role).Where("id = ?", roleId).First(&roleDetail).Error; err != nil {
		return nil, fmt.Errorf("角色%w", errmsg.QueryFailed), err.Error()
	}
	return
}

func (RoleService) RoleUpdateService(roleId string, params *request.RoleUpdate) (err error, debugInfo interface{}) {
	var r role.Role
	if err = global.DB.Table(table.Role).Where("id = ?", roleId).First(&r).Error; err != nil {
		return fmt.Errorf("角色%w", errmsg.QueryFailed), err.Error()
	}
	if r.IsSystem == enums.True {
		return fmt.Errorf(errmsg.NotPrivilege.Error(), "修改系统默认角色"), nil
	}
	var marshal []byte
	//过滤无效权限key
	if params.PrivilegeKeys != nil {
		var privilegeKeys []string
		if err = global.DB.Table(table.Privilege).
			Where("key in ?", params.PrivilegeKeys).
			Select("key").Find(&privilegeKeys).Error; err != nil {
			return fmt.Errorf("权限key%w", errmsg.QueryFailed), err.Error()
		}
		marshal, err = json.Marshal(privilegeKeys)
		if err != nil {
			return errmsg.JsonConvertFiled, err.Error()
		}
	}
	if err = global.DB.Model(&r).Updates(role.Role{Name: params.Name, PrivilegeKeys: marshal}).Error; err != nil {
		return fmt.Errorf("角色%w", errmsg.UpdateFailed), err.Error()
	}
	return
}
