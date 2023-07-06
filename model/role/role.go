package role

import (
	"base-system-backend/enums/table"
	"gorm.io/datatypes"
)

type Role struct {
	Id            int64          `json:"id" gorm:"column:id;primaryKey;autoIncrement;notNull;comment:Id"`
	Name          string         `json:"name" gorm:"column:name;notNull;unique;size:16;comment:角色名称"`
	IsSystem      bool           `json:"is_system" gorm:"column:is_system;notNull;default:false;comment:是否系统默认角色"`
	PrivilegeKeys datatypes.JSON `json:"privilege_keys" gorm:"column:privilege_keys;comment:角色key列表"`
}

func (receiver Role) TableName() string {
	return table.Role
}
