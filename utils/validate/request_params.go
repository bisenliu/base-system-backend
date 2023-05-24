package validate

import (
	"base-system-backend/enums/code"
	"base-system-backend/global"
	"base-system-backend/model/common/response"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"reflect"
	"strings"
)

func RequestParamsVerify(c *gin.Context, params interface{}) (res bool) {
	// 绑定请求参数
	if err := c.ShouldBindBodyWith(&params, binding.JSON); err != nil {
		global.LOG.Error("invalid request params: ", zap.Error(err))
		if validateError, ok := err.(validator.ValidationErrors); ok {
			response.Error(c, code.InvalidParams, removeTopStruct(validateError.Translate(global.TRANS)), err.Error())
			return
		}
		if unmarshalTypeError, ok := err.(*json.UnmarshalTypeError); ok {
			typeOfCat := reflect.TypeOf(params)
			//如果是指针类型则转换为非指针
			if typeOfCat.Kind() == reflect.Ptr {
				typeOfCat = typeOfCat.Elem()
			}
			for i := 0; i < typeOfCat.NumField(); i++ {
				// 获取每个成员的结构体字段类型
				fieldType := typeOfCat.Field(i)
				var (
					rawFailed string
					value     string
				)
				if unmarshalTypeError.Field == strings.Split(fieldType.Tag.Get("json"), ",")[0] {
					rawFailed = strings.Split(fieldType.Tag.Get("json"), ",")[0]
					value = fieldMapping(unmarshalTypeError)
					if fieldType.Tag.Get("label") != "" {
						rawFailed = fieldType.Tag.Get("label")
					}
					msg := strings.Join([]string{fmt.Sprintf("%s", rawFailed), "字段不合法,", "应为", value}, "")
					response.Error(c, code.InvalidParams, map[string]interface{}{
						unmarshalTypeError.Field: msg,
					}, unmarshalTypeError.Error())
					return
				}
			}
		} else {
			response.Error(c, code.InvalidParams, err, err.Error())
			return
		}
	}
	return true
}

func fieldMapping(unmarshalTypeError *json.UnmarshalTypeError) string {
	switch unmarshalTypeError.Type.String() {
	case "int", "int64", "[]int", "[]int64", "enums.Gender":
		return "整型"
	case "string":
		return "字符串类型"
	default:
		return unmarshalTypeError.Type.String()
	}
}
