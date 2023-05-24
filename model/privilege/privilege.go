package privilege

import (
	"base-system-backend/enums/table"
	"gorm.io/datatypes"
)

type Privilege struct {
	Id         int64          `gorm:"column:id;primaryKey;autoIncrement;notNull;comment:Id"`
	Title      int            `gorm:"column:title;notNull;size:20;comment:系统权限名称"`
	Key        int            `gorm:"column:key;size:50;comment:系统权限key"`
	Dependency datatypes.JSON `gorm:"column:dependency;comment:依赖权限key"`
	ParentId   string         `gorm:"column:parent_id;comment:父级Id"`
}

func (receiver Privilege) TableName() string {
	return table.Privilege
}
