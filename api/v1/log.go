package v1

import (
	"base-system-backend/enums/code"
	"base-system-backend/model/common/response"
	"base-system-backend/model/log/request"
	"base-system-backend/utils/validate"
	"github.com/gin-gonic/gin"
)

type LogApi struct{}

// OperateLogListApi
// @Summary 操作日志列表
// @Description 操作日志列表
// @Tags LogApi
// @Accept application/json
// @Produce application/json
// @Param Identification header string true "Token 令牌"
// @Param object query request.OperateLogFilter false "查询参数"
// @Security ApiKeyAuth
// @Success 200 {object} response.Data{data=response.OperateLogList}
// @Router /log/operate/list/ [get]
func (LogApi) OperateLogListApi(c *gin.Context) {
	params := new(request.OperateLogFilter)
	if !validate.QueryParamsVerify(c, params) {
		return
	}
	operateLogList, err, debugInfo := logService.OperateLogListService(c, params)
	if err != nil {
		response.Error(c, code.QueryFailed, err, debugInfo)
		return
	}
	response.OK(c, operateLogList)
}

// OperateLogDownloadApi
// @Summary 操作日志下载
// @Description 操作日志下载
// @Tags LogApi
// @Accept application/json
// @Produce application/vnd.openxmlformats-officedocument.spreadsheetml.sheet
// @Param Identification header string true "Token 令牌"
// @Param object query request.OperateLogFilter false "查询参数"
// @Security ApiKeyAuth
// @Router /log/operate/download/ [get]
func (LogApi) OperateLogDownloadApi(c *gin.Context) {
	params := new(request.OperateLogFilter)
	if !validate.QueryParamsVerify(c, params) {
		return
	}
	content, err, debugInfo := logService.OperateLogDownloadService(c, params)
	if err != nil {
		response.Error(c, code.QueryFailed, err, debugInfo)
		return
	}
	response.File(c, content, "操作日志")
}
