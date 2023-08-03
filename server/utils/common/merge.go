package common

import (
	"base-system-backend/constants/errmsg"
	"encoding/json"
	"fmt"
)

// Merge
//  @Description: 合并 Slice
//  @param arr Slice
//  @return newArr 新 Slice
//  @return err 转换失败异常
//  @return debugInfo 错误调试信息

func Merge(arr []string) (newArr []string, err error, debugInfo interface{}) {
	for _, keys := range arr {
		var jsonArr []string
		if err = json.Unmarshal([]byte(keys), &jsonArr); err != nil {
			return nil, errmsg.JsonConvertFiled, err.Error()
		}
		newArr = append(newArr, jsonArr...)
	}
	fmt.Println(newArr)
	return
}
