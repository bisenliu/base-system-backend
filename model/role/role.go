package role

import (
	"base-system-backend/enums"
	"base-system-backend/enums/table"
	"gorm.io/datatypes"
)

type Role struct {
	Id            int64          `gorm:"column:id;primaryKey;autoIncrement;notNull;comment:Id"`
	Name          string         `gorm:"column:name;notNull;unique;size:16;comment:角色名称"`
	IsSystem      enums.BoolSign `gorm:"column:is_system;notNull;default:0;comment:是否系统默认角色"`
	PrivilegeKeys datatypes.JSON `gorm:"column:privilege_keys;comment:角色key列表"`
}

func (receiver Role) TableName() string {
	return table.Role
}
