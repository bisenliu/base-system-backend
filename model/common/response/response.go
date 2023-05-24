package response

import (
	"base-system-backend/enums/code"
	"base-system-backend/enums/errmsg"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

type Data struct {
	Status     int         `json:"status"`
	StatusInfo statusInfo  `json:"status_info"`
	Data       interface{} `json:"data"`
}

type statusInfo struct {
	Message   interface{} `json:"message"`
	Detail    interface{} `json:"detail"`
	DebugInfo interface{} `json:"debug_info"`
}

func OK(c *gin.Context, data interface{}) {
	status, _ := strconv.Atoi(fmt.Sprint(code.Success))
	c.JSON(http.StatusOK, &Data{Status: status, Data: data})
}

func Error(c *gin.Context, statusCode code.StatusCode, errorInfo interface{}, debugInfo interface{}) {
	// 参数错误
	errMsg, detail := getErrorInfo(statusCode, errorInfo)
	// 组装状态码

	status := generateStatusCode(c, statusCode)
	// 是否登录失败错误
	var data interface{}
	if value, ok := debugInfo.(map[string]interface{}); ok {
		if value["next_time"] != nil {
			data = debugInfo
			debugInfo = nil
			errMsg = fmt.Errorf(errmsg.LoginOutLimit.Error(), value["failed_num"], value["login_time"]).Error()
		}
	}
	c.JSON(http.StatusOK, &Data{
		Status: status,
		StatusInfo: statusInfo{
			Message:   errMsg,
			Detail:    detail,
			DebugInfo: debugInfo,
		},
		Data: data,
	})
}

func getErrorInfo(statusCode code.StatusCode, errorInfo interface{}) (errMsg interface{}, detail interface{}) {
	if statusCode == code.InvalidParams {
		detail = errorInfo
		if errMap, ok := errorInfo.(map[string]string); ok {
			for _, v := range errMap {
				errMsg = v
				break
			}
		} else {
			errMsg = statusCode.Msg()
		}

	} else {
		if errorInfo != nil {
			errMsg = fmt.Sprintf("%s", errorInfo)
		} else {
			errMsg = statusCode.Msg()
		}
		detail = nil
	}
	return
}

func generateStatusCode(c *gin.Context, statusCode code.StatusCode) int {
	urlSlice := strings.Split(fmt.Sprintf("%s", c.Request.URL), "/")
	prefix := urlSlice[2]
	return code.GetStatusCodeByModelCode(prefix, statusCode)
}
