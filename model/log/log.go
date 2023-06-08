package log

import (
	"base-system-backend/enums"
	"base-system-backend/enums/table"
	"base-system-backend/model/common/field"
	"gorm.io/datatypes"
)

type OperateLog struct {
	Id         int64            `gorm:"column:id;primaryKey;autoIncrement;notNull;comment:操作日志ID"`
	UserId     int64            `gorm:"column:user_id;comment:请求者的用户ID"`
	ActionName string           `gorm:"column:action_name;notNull;size:100;comment:行为名称"`
	Module     string           `gorm:"column:module;notNull;size:20;comment:模块名称"`
	AccessUrl  string           `gorm:"column:access_url;notNull;size:200;comment:访问Url"`
	RequestIp  string           `gorm:"column:request_ip;notNull;size:100;comment:访问时的ip"`
	UserAgent  string           `gorm:"column:user_agent;notNull;size:200;comment:请求者的UserAgent"`
	AccessTime field.CustomTime `gorm:"column:access_time;autoCreateTime;comment:访问时间"`
	Success    enums.BoolSign   `gorm:"column:success;notNull;default:0;comment:操作是否成功"`
	Detail     datatypes.JSON   `gorm:"column:detail;comment:失败原因"`
}

func (OperateLog) TableName() string {
	return table.OperateLog
}
