package validate

import (
	"base-system-backend/constants/errmsg"
	"base-system-backend/global"
	"fmt"
	"mime/multipart"
	"net/http"
	"strings"
)

//  常用图片格式

var imageExtensions = map[string]bool{
	"jpg":  true,
	"jpeg": true,
	"png":  true,
	"gif":  true,
	"bmp":  true,
	"webp": true,
}

// ImageVerify
//  @Description: 上传图片校验
//  @param fileHeader 文件头
//  @return err 校验失败异常
//  @return debugInfo 错误调试信息

func ImageVerify(fileHeader *multipart.FileHeader) (err error, debugInfo interface{}) {
	if fileHeader.Size > global.CONFIG.Avatar.Size*1024*1024 {
		return fmt.Errorf(errmsg.FileSizeRange.Error(), int(global.CONFIG.Avatar.Size)), nil
	}
	buff := make([]byte, 512)
	file, err := fileHeader.Open()
	if err != nil {
		return fmt.Errorf("头像文件%w", errmsg.Invalid), err.Error()
	}
	if _, err = file.Read(buff); err != nil {
		return fmt.Errorf("头像文件%w", errmsg.ReadFailed), err.Error()
	}
	filetype := http.DetectContentType(buff)
	ext := strings.ToLower(strings.TrimPrefix(filetype, "image/"))
	if !imageExtensions[ext] {
		return fmt.Errorf("头像文件%w", errmsg.Invalid), nil
	}
	return
}
