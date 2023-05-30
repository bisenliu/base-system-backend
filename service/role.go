package service

import (
	"base-system-backend/enums/errmsg"
	"base-system-backend/enums/table"
	"base-system-backend/global"
	"base-system-backend/model/role"
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
