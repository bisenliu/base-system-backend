package v1

import (
	"base-system-backend/enums/code"
	"base-system-backend/model/common/response"
	"base-system-backend/model/log/request"
	"base-system-backend/utils/validate"
	"github.com/gin-gonic/gin"
)

type LogApi struct{}

func (LogApi) OperateLogListApi(c *gin.Context) {
	params := new(request.OperateLogFilter)
	if ok := validate.QueryParamsVerify(c, params); !ok {
		return
	}
	operateLogList, err, debugInfo := logService.OperateLogListService(c, params)
	if err != nil {
		response.Error(c, code.QueryFailed, err, debugInfo)
	}
	response.OK(c, operateLogList)
}

func (LogApi) OperateLogDownloadApi(c *gin.Context) {
	params := new(request.OperateLogFilter)
	if ok := validate.QueryParamsVerify(c, params); !ok {
		return
	}
	content, err, debugInfo := logService.OperateLogDownloadService(c, params)
	if err != nil {
		response.Error(c, code.QueryFailed, err, debugInfo)
	}
	response.File(c, content, "操作日志")
}
