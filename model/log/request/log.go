package request

import (
	"base-system-backend/model/common/field"
	"base-system-backend/model/common/request"
	"gorm.io/datatypes"
)

type OperateLogFilter struct {
	ActionName      string `json:"action_name" form:"action_name" label:"行为名称"`
	Module          string `json:"module" form:"module" label:"模块名称"`
	RequestIp       string `json:"request_ip" form:"request_ip" label:"访问时的ip"`
	UserId          *int64 `json:"user_id" form:"user_id" label:"请求者的用户ID"`
	StartAccessTime *int64 `json:"start_access_time" form:"start_access_time" label:"访问开始时间"`
	EndAccessTime   *int64 `json:"end_access_time" form:"end_access_time" label:"访问结束时间"`
	Success         bool   `json:"success" form:"success" binding:"enum" label:"操作是否成功"`
	request.PageInfo
}

type OperateLogCreate struct {
	UserId     *int64           `json:"user_id"`
	ActionName string           `json:"action_name"`
	Module     string           `json:"module"`
	AccessUrl  string           `json:"access_url"`
	RequestIp  string           `json:"request_ip"`
	UserAgent  string           `json:"user_agent"`
	AccessTime field.CustomTime `json:"access_time"`
	Success    bool             `json:"success"`
	Detail     datatypes.JSON   `json:"detail"`
}
