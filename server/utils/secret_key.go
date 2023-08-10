package utils

import (
	"crypto/rand"
	"encoding/base64"
)

// GenerateSecretKey
//  @Description: 生成64位 key
//  @return string key
//  @return error 读取失败异常

func GenerateSecretKey() (string, error) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), nil
}
