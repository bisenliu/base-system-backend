package validate

import (
	"strings"
)

// removeTopStruct
//  @Description: 去除提示信息中的结构体名称
//  @param fields 结构体字段
//  @return map[string]string 修改后的结构体

func removeTopStruct(fields map[string]string) map[string]string {
	res := map[string]string{}
	for field, err := range fields {
		//fmt.Println("field, err: ", field, err)
		fieldSlice := strings.Split(err, "|")
		if len(fieldSlice) == 2 {
			res[fieldSlice[0]] = fieldSlice[1]
		} else {
			res[field[strings.Index(field, ".")+1:]] = err
		}
	}
	return res
}
