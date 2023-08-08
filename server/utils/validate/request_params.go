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

// RequestParamsVerify
//  @Description: http 请求参数校验
//  @param c 上下文信息
//  @param params 请求参数
//  @return bool 是否校验成功

func RequestParamsVerify(c *gin.Context, params interface{}) bool {
	// 绑定请求参数
	if err := c.ShouldBindBodyWith(&params, binding.JSON); err != nil {
		handleError(c, params, err)
		return false
	}
	return true
}

// handleError
//  @Description: 请求参数校验失败异常判断(根据不同错误类型去处理异常)
//  @param c 上下文信息
//  @param params 请求参数
//  @param err 参数校验失败异常

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

// getFieldJSONName
//  @Description: 处理 json.UnmarshalTypeError 异常,获取 struct 字段别名,获取不到则为该字段本身
//  @param params 请求参数
//  @param unmarshalTypeError 反序列化失败异常
//  @return string struct 名称/字段别名

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

// buildErrorMessage
//  @Description: 根据不同错误类型返回对应的错误信息
//  @param fieldName 字段名称
//  @param fieldType 字段类型
//  @return string 错误信息

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
