package service

import (
	"base-system-backend/constants/errmsg"
	"base-system-backend/constants/table"
	"base-system-backend/global"
	"base-system-backend/model/role"
	"base-system-backend/model/role/request"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
)

type RoleService struct{}

// RoleListService
//  @Description: 角色列表 service
//	@receiver RoleService
//  @param c 上下文信息
//  @return roleList 角色列表
//  @return err 查询失败异常
//  @return debugInfo 错误调试信息

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

// RoleCreateService
//  @Description: 创建角色 service
//	@receiver RoleService
//  @param params 角色信息
//  @return err 创建失败异常
//  @return debugInfo 错误调试信息

func (RoleService) RoleCreateService(params *request.RoleCreate) (err error, debugInfo interface{}) {
	if err = global.DB.Table(table.Role).Create(params).Error; err != nil {
		return fmt.Errorf("角色%w", errmsg.SaveFailed), err.Error()
	}
	return
}

// RoleDetailService
//  @Description: 角色详情 service
//	@receiver RoleService
//  @param roleId 角色 ID
//  @return roleDetail 角色信息
//  @return err 查询失败异常
//  @return debugInfo 错误调试信息

func (RoleService) RoleDetailService(roleId string) (roleDetail *role.Role, err error, debugInfo interface{}) {
	if err = global.DB.Table(table.Role).Where("id = ?", roleId).First(&roleDetail).Error; err != nil {
		return nil, fmt.Errorf("角色%w", errmsg.QueryFailed), err.Error()
	}
	return
}

// RoleUpdateService
//  @Description: 更新角色信息 service
//	@receiver RoleService
//  @param roleId 角色 ID
//  @param params 角色信息
//  @return err 更新/查询失败异常
//  @return debugInfo 错误调试信息

func (RoleService) RoleUpdateService(roleId string, params *request.RoleUpdate) (err error, debugInfo interface{}) {
	var r role.Role
	if err = global.DB.Table(table.Role).Where("id = ?", roleId).First(&r).Error; err != nil {
		return fmt.Errorf("角色%w", errmsg.QueryFailed), err.Error()
	}
	if r.IsSystem == true {
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

// RoleDeleteService
//  @Description: 删除指定角色 service
//	@receiver RoleService
//  @param roleId 角色 ID
//  @return err 查询/删除失败异常
//  @return debugInfo 错误调试信息

func (RoleService) RoleDeleteService(roleId string) (err error, debugInfo interface{}) {
	var r role.Role
	if err = global.DB.Table(table.Role).Where("id = ?", roleId).First(&r).Error; err != nil {
		return fmt.Errorf("角色%w", errmsg.QueryFailed), err.Error()
	}
	if r.IsSystem == true {
		return fmt.Errorf(errmsg.NotPrivilege.Error(), "删除系统默认角色"), nil
	}
	if err = global.DB.Delete(&r).Error; err != nil {
		return fmt.Errorf("角色%w", errmsg.DeleteFailed), err.Error()
	}
	return
}
