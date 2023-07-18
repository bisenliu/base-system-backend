package field

import (
	"base-system-backend/constants/errmsg"
	"base-system-backend/global"
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"go.uber.org/zap"
	"strings"
)

type PlainEncrypt string
type SplitEncrypt string

// Value 普通字段加密入库
func (p PlainEncrypt) Value() (driver.Value, error) {
	if p != "" {
		res, err := encryptAES([]byte(p), []byte(global.CONFIG.Key.AesKey))
		if err != nil {
			global.LOG.Error("plain fields encrypt failed: ", zap.Error(err))
		}
		newres, _ := json.Marshal(res)
		return newres, err
	}
	return nil, nil
}

// Scan 普通字段解密
func (p *PlainEncrypt) Scan(value interface{}) error {
	str, ok := value.(string)
	if !ok {
		global.LOG.Error("plain fields type error")
		return errmsg.DecryptFailed
	}
	if value != "" {
		var byteValue []byte
		if err := json.Unmarshal([]byte(str), &byteValue); err != nil {
			global.LOG.Error("split fields marshal failed: ", zap.Error(err))
			return err
		}
		res, err := decryptAES(byteValue, []byte(global.CONFIG.Key.AesKey))
		if err != nil {
			global.LOG.Error("plain fields decrypt failed: ", zap.Error(err))
		} else {
			*p = PlainEncrypt(res)
		}
		return err
	}
	return nil
}

func (p SplitEncrypt) Value() (driver.Value, error) {
	//切分字符串到切片中
	resArray := strings.Split(string(p), "")
	if len(resArray) > 0 {
		var valueArray [][]byte
		for _, v := range resArray {
			res, err := encryptAES([]byte(v), []byte(global.CONFIG.Key.AesKey))
			if err != nil {
				global.LOG.Error("split fields encrypt failed: ", zap.Error(err))
				break
			}
			valueArray = append(valueArray, res)
		}
		//重新组装加密数据
		var jsonArray []string
		for _, v := range valueArray {
			res, err := json.Marshal(v)
			if err != nil {
				global.LOG.Error("split fields marshal failed: ", zap.Error(err))
				break
			}
			jsonArray = append(jsonArray, string(res))
		}
		res := strings.Join(jsonArray, "")
		return res, nil
	}
	return nil, nil
}

func (p *SplitEncrypt) Scan(value interface{}) error {
	str, ok := value.(string)
	if !ok {
		global.LOG.Error("split fields type error")
		return errmsg.DecryptFailed
	}
	if value != "" {
		//切分
		strArray := strings.Split(str, `="`)
		fmt.Println("strArray:", strArray)
		var resArray []string
		for _, v := range strArray {
			if v != "" {
				var byteValue []byte
				if err := json.Unmarshal([]byte(v+`="`), &byteValue); err != nil {
					global.LOG.Error("split fields marshal failed: ", zap.Error(err))
					return err
				}
				//解密并存放到切片
				res, err := decryptAES(byteValue, []byte(global.CONFIG.Key.AesKey))
				if err != nil {
					global.LOG.Error("split fields decrypt failed: ", zap.Error(err))
					return err
				}
				resArray = append(resArray, string(res))
			}
		}
		//合并字符串切片
		*p = SplitEncrypt(strings.Join(resArray, ""))
		return nil
	}
	return nil
}

// 填充数据
func padding(src []byte, blockSize int) []byte {
	padNum := blockSize - len(src)%blockSize
	pad := bytes.Repeat([]byte{byte(padNum)}, padNum)
	return append(src, pad...)
}

// 去掉填充数据
func unPadding(src []byte) []byte {
	n := len(src)
	unPadNum := int(src[n-1])
	return src[:n-unPadNum]
}

// EncryptAES 加密
func encryptAES(src []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	src = padding(src, block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, key)
	blockMode.CryptBlocks(src, src)
	return src, nil
}

// DecryptAES 解密
func decryptAES(src []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, fmt.Errorf("字段解密失败：%w", err)
	}
	blockMode := cipher.NewCBCDecrypter(block, key)
	blockMode.CryptBlocks(src, src)
	src = unPadding(src)
	return src, nil
}
