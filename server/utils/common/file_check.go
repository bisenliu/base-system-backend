package common

import (
	"base-system-backend/constants/errmsg"
	"fmt"
	"os"
)

// FileCheck
//  @Description: 路径不存在则创建
//  @param path 路径
//  @return err 创建失败异常
//  @return debugInfo 错误调试信息

func FileCheck(path string) (err error, debugInfo interface{}) {
	_, err = os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			if err = os.MkdirAll(path, os.ModePerm); err != nil {
				return fmt.Errorf("文件%w", errmsg.SaveFailed), err.Error()
			}
		} else {
			return fmt.Errorf("文件%w", errmsg.ReadFailed), err.Error()
		}
	}
	return
}
