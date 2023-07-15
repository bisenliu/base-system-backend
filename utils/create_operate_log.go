package utils

import (
	"base-system-backend/enums/code"
	"base-system-backend/enums/table"
	"base-system-backend/global"
	"base-system-backend/model/common/field"
	"base-system-backend/model/log/request"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"reflect"
	"regexp"
	"runtime"
	"time"
)

func CreateOperateLog(c *gin.Context, success bool, detailByte []byte) {
	var (
		actionName string
		userId     *int64
	)
	//获取当前请求的处理函数
	handler := c.Handler()
	//使用反射获取函数名称
	funcName := runtime.FuncForPC(reflect.ValueOf(handler).Pointer()).Name()
	//获取函数的注释
	comments := getFunctionComments(funcName)
	re := regexp.MustCompile(`@Description (\S+)`)
	match := re.FindStringSubmatch(comments)
	if len(match) > 1 {
		actionName = match[1]
	} else {
		actionName = funcName
	}
	scheme := "http"
	if c.Request.TLS != nil {
		scheme = "https"
	}
	accessUrl := fmt.Sprintf("%s://%s%s", scheme, c.Request.Host, c.Request.RequestURI)
	userAgent := c.Request.UserAgent()
	userInstance, err, _ := GetCurrentUser(c)
	if err != nil {
		userId = nil
	} else {
		userId = &userInstance.Id
	}
	prefix := GetUrlPrefix(c.Request.RequestURI)
	model := code.Model{}.Choices(prefix).Desc()
	if err = global.DB.Table(table.OperateLog).Create(&request.OperateLogCreate{
		UserId:     userId,
		ActionName: actionName,
		Module:     model,
		AccessUrl:  accessUrl,
		RequestIp:  c.ClientIP(),
		UserAgent:  userAgent,
		Success:    success,
		Detail:     detailByte,
		AccessTime: field.CustomTime(time.Now()),
	}).Error; err != nil {
		global.LOG.Error("create operate log info failed: %s", zap.Error(err))
	}

	logFields := []zap.Field{
		zap.Bool("status", success),
		zap.String("method", c.Request.Method),
		zap.String("query", c.Request.URL.RawQuery),
		zap.String("ip", c.ClientIP()),
		zap.String("user-agent", c.Request.UserAgent()),
	}
	if len(detailByte) > 0 {
		logFields = append(logFields, zap.String("errors", string(detailByte)))
	}
	if success {
		global.LOG.Info(c.Request.URL.Path, logFields...)
	} else {
		global.LOG.Error(c.Request.URL.Path, logFields...)
	}
}
