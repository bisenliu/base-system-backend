package validate

import (
	"base-system-backend/constants/code"
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

func RequestParamsVerify(c *gin.Context, params interface{}) bool {
	// 绑定请求参数
	if err := c.ShouldBindBodyWith(&params, binding.JSON); err != nil {
		handleError(c, params, err)
		return false
	}
	return true
}

func handleError(c *gin.Context, params interface{}, err error) {
	global.LOG.Error("invalid request params: ", zap.Error(err))
	switch err := err.(type) {
	case validator.ValidationErrors:
		validateError := err
		response.Error(c, code.InvalidParams, removeTopStruct(validateError.Translate(global.TRANS)), err.Error())
	case *json.UnmarshalTypeError:
		unmarshalTypeError := err
		fieldName := getFieldJSONName(params, err)
		errorMap := map[string]string{
			unmarshalTypeError.Field: buildErrorMessage(fieldName, unmarshalTypeError.Type.String()),
		}
		response.Error(c, code.InvalidParams, errorMap, unmarshalTypeError.Error())
	default:
		response.Error(c, code.InvalidParams, err, err.Error())
	}
}

func getFieldJSONName(params interface{}, unmarshalTypeError *json.UnmarshalTypeError) string {
	var rawFailed string
	typeOfCat := reflect.TypeOf(params)
	//如果是指针类型则转换为非指针
	if typeOfCat.Kind() == reflect.Ptr {
		typeOfCat = typeOfCat.Elem()
	}
	for i := 0; i < typeOfCat.NumField(); i++ {
		// 获取每个成员的结构体字段类型
		fieldType := typeOfCat.Field(i)
		if unmarshalTypeError.Field == strings.Split(fieldType.Tag.Get("json"), ",")[0] {
			rawFailed = strings.Split(fieldType.Tag.Get("json"), ",")[0]
			if fieldType.Tag.Get("label") != "" {
				rawFailed = fieldType.Tag.Get("label")
			} else {
				rawFailed = unmarshalTypeError.Field
			}
		}
	}
	return rawFailed
}

func buildErrorMessage(fieldName, fieldType string) string {
	switch fieldType {
	case "int", "int64", "[]int", "[]int64", "enums.Gender":
		return fmt.Sprintf("%s字段不合法，应为整型", fieldName)
	case "string":
		return fmt.Sprintf("%s字段不合法，应为字符串类型", fieldName)
	default:
		return fmt.Sprintf("%s字段不合法，应为%s类型", fieldName, fieldType)
	}
}
