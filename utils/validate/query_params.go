package validate

import (
	"base-system-backend/enums/code"
	"base-system-backend/global"
	"base-system-backend/model/common/response"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

func QueryParamsVerify(c *gin.Context, params interface{}) (res bool) {
	// 绑定查询参数
	if err := c.ShouldBindQuery(params); err != nil {
		global.LOG.Error("invalid query params: ", zap.Error(err))
		if validateError, ok := err.(validator.ValidationErrors); ok {
			response.Error(c, code.InvalidParams, removeTopStruct(validateError.Translate(global.TRANS)), err.Error())
			return
		}
		response.Error(c, code.InvalidParams, nil, err.Error())
		return
	}
	return true
}
