package privilege

import (
	"base-system-backend/constants/table"
	"gorm.io/datatypes"
)

type Privilege struct {
	Id         int64          `json:"id" gorm:"column:id;primaryKey;autoIncrement;notNull;comment:Id"`
	ParentId   int64          `json:"parent_id" gorm:"column:parent_id;comment:父级Id"`
	Title      string         `json:"title" gorm:"column:title;notNull;size:20;comment:系统权限名称"`
	Key        string         `json:"key" gorm:"column:key;size:50;comment:系统权限key"`
	Dependency datatypes.JSON `json:"dependency" gorm:"column:dependency;comment:依赖权限key"`
}

func (receiver Privilege) TableName() string {
	return table.Privilege
}
