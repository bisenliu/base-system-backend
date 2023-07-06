package response

import (
	"base-system-backend/model/common/field"
	"base-system-backend/model/common/response"
	"gorm.io/datatypes"
)

type OperateLogErrMessage struct {
	Message string `json:"message"`
}

type OperateLogDetail struct {
	Id          int64             `json:"id"`
	ActionName  string            `json:"action_name"`
	Module      string            `json:"module"`
	AccessUrl   string            `json:"access_url"`
	RequestIp   string            `json:"request_ip"`
	UserAgent   string            `json:"user_agent"`
	UserId      *int64            `json:"user_id"`
	UserName    string            `json:"user_name"`
	UserAccount string            `json:"user_account"`
	AccessTime  *field.CustomTime `json:"access_time"`
	Success     bool              `json:"success"`
	Detail      *datatypes.JSON   `json:"-"`
	OperateLogErrMessage
}

type OperateLogList struct {
	response.PageInfo
	Results []OperateLogDetail `json:"results" form:"results"` //数据
}

type OperateLogDownload struct {
	Id          int64  `json:"id"`
	ActionName  string `json:"action_name"`
	Module      string `json:"module"`
	AccessUrl   string `json:"access_url"`
	RequestIp   string `json:"request_ip"`
	UserAgent   string `json:"user_agent"`
	UserId      string `json:"user_id"`
	UserName    string `json:"user_name"`
	UserAccount string `json:"user_account"`
	AccessTime  string `json:"access_time"`
	Success     string `json:"success"`
	Message     string `json:"message"`
}
