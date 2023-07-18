package utils

import (
	"fmt"
	"strings"
)

func GetUrlPrefix(url string) string {
	urlSlice := strings.Split(fmt.Sprintf("%s", url), "/")
	return urlSlice[2]
}
