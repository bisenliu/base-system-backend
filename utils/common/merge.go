package common

import (
	"base-system-backend/enums/errmsg"
	"encoding/json"
)

func Merge(arr []string) (newArr []string, err error, debugInfo interface{}) {
	for _, keys := range arr {
		var jsonArr []string
		if err = json.Unmarshal([]byte(keys), &jsonArr); err != nil {
			return nil, errmsg.JsonConvertFiled, err.Error()
		}
		newArr = append(newArr, jsonArr...)
	}
	return
}
