package utils

import (
	"fmt"
	"strings"
)

// GetUrlPrefix
//  @Description: 获取 url 前缀
//  @param url url
//  @return string 前缀信息

func GetUrlPrefix(url string) string {
	urlSlice := strings.Split(fmt.Sprintf("%s", url), "/")
	// 根据 api 不同,调整切片索引
	return urlSlice[2]
}
