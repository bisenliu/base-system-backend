package common

import (
	"bytes"
	"github.com/mozillazg/go-pinyin"
	"strings"
)

// ConvertCnToLetter
//  @Description: 姓名全/简拼转换
//  @param cn 姓名
//  @return fullName 全拼
//  @return shortName 简拼

func ConvertCnToLetter(cn string) (fullName string, shortName string) {
	a := pinyin.NewArgs()
	a.Fallback = func(r rune, a pinyin.Args) []string {
		return []string{string(r)}
	}
	rows := pinyin.Pinyin(cn, a)
	for i := 0; i < len(rows); i++ {
		if len(rows[i]) != 0 {
			str := rows[i][0]
			pi := str[0:1]
			fullName += strings.ToLower(str)
			shortName += string(bytes.ToLower([]byte(pi)))
		}
	}
	return
}
