package code

import (
	"fmt"
	"strconv"
)

type StatusCode string

const (
	Success        StatusCode = "0"
	InvalidLogin   StatusCode = "001"
	InvalidParams  StatusCode = "002"
	QueryFailed    StatusCode = "003"
	SaveFailed     StatusCode = "004"
	UpdateFailed   StatusCode = "005"
	DeleteFailed   StatusCode = "006"
	NotPermissions StatusCode = "007"

	UnknownExc StatusCode = "999"
)

var statusCodeMapping = map[StatusCode]string{
	Success:       "Success",
	InvalidParams: "无效的参数",
	UnknownExc:    "未知错误",
}

func (receiver StatusCode) Msg() string {
	msg, ok := statusCodeMapping[receiver]
	if !ok {
		msg = statusCodeMapping[UnknownExc]
	}
	return msg
}

func GetStatusCodeByModelCode(urlPrefix string, statusCode StatusCode) int {

	modelCode, ok := ModelMapping[urlPrefix]
	if !ok {
		modelCode = Unknown
	}
	if statusCode == InvalidLogin {
		modelCode = User
	}
	code, _ := strconv.Atoi(fmt.Sprintf("%s", modelCode.Code()) + fmt.Sprintf("%s", statusCode))
	return code
}
