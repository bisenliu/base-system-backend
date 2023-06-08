package middleware

import (
	"base-system-backend/enums"
	"base-system-backend/enums/code"
	"base-system-backend/enums/table"
	"base-system-backend/global"
	"base-system-backend/model/common/field"
	"base-system-backend/model/log/request"
	"base-system-backend/utils"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"runtime"
	"strings"
	"time"
)

type bodyLogWriter struct {
	//嵌入gin框架ResponseWriter
	gin.ResponseWriter
	body *bytes.Buffer //我们记录用的response
}

// Write 写入响应体数据
func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)                  //我们记录一份
	return w.ResponseWriter.Write(b) //真正写入响应
}

// 辅助函数，通过函数名称获取注释
func getFunctionComments(funcName string) string {
	goFiles, err := getAllGoFiles()
	if err != nil {
		fmt.Println("获取 Go 文件失败:", err)
		return ""
	}

	for _, filePath := range goFiles {
		fset := token.NewFileSet()
		file, err := parser.ParseFile(fset, filePath, nil, parser.ParseComments)
		if err != nil {
			fmt.Printf("解析文件 %s 失败: %s\n", filePath, err)
			continue
		}

		for _, decl := range file.Decls {
			if fd, ok := decl.(*ast.FuncDecl); ok {
				//fmt.Println(fd.Name.Name, funcName)
				if isFunctionMatch(fd.Name.Name, funcName) && fd.Doc != nil {
					return fd.Doc.Text()
				}
			}
		}
	}

	return ""
}

// 辅助函数，检查函数名是否匹配
func isFunctionMatch(declName, funcName string) bool {
	parts := strings.Split(funcName, "/")
	if len(parts) > 0 {
		funcName = parts[len(parts)-1]
	}

	// 提取函数名部分
	parts = strings.Split(funcName, ".")
	if len(parts) > 0 {
		funcName = parts[len(parts)-1]
	}

	return declName == funcName
}

// 辅助函数，通过处理函数获取函数名称
func getFunctionName(handler gin.HandlerFunc) string {
	handlerType := reflect.TypeOf(handler)
	if handlerType.Kind() == reflect.Func {
		handlerValue := reflect.ValueOf(handler)
		handlerName := runtime.FuncForPC(handlerValue.Pointer()).Name()

		parts := strings.Split(strings.Join(strings.Split(handlerName, "/-fm"), ""), "/")
		if len(parts) > 0 {
			funcName := parts[len(parts)-1]

			parts = strings.Split(funcName, ".")
			if len(parts) > 0 {
				funcName = parts[len(parts)-1]
			}

			return strings.TrimSuffix(funcName, "-fm")
		}
	}
	return ""
}

// 辅助函数，获取所有 Go 文件
func getAllGoFiles() ([]string, error) {
	var goFiles []string

	err := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		if filepath.Ext(path) == ".go" {
			goFiles = append(goFiles, path)
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return goFiles, nil
}

func OperateLogMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		var (
			rsp        map[string]interface{}
			success    enums.BoolSign
			detailByte []byte
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
		accessUrl := strings.Join([]string{scheme, "://", c.Request.Host, c.Request.RequestURI}, "")
		method := c.Request.Method
		userAgent := c.Request.UserAgent()
		userInstance, err, _ := utils.GetCurrentUser(c)
		if err != nil {
			userId = nil
		} else {
			userId = &userInstance.Id
		}
		blw := &bodyLogWriter{body: bytes.NewBuffer([]byte{}), ResponseWriter: c.Writer}
		c.Writer = blw
		c.Next()
		err = json.Unmarshal(blw.body.Bytes(), &rsp)
		if err != nil {
			global.LOG.Error("body convert json failed:%w ", zap.Error(err))
			c.Abort()
			return
		}
		status := rsp["status"]
		if value, ok := status.(float64); ok {
			if value == float64(0) {
				if method == "GET" {
					c.Abort()
					return
				}
				success = enums.True
			} else {
				success = enums.False
			}
		}
		detail := make(map[string]interface{})
		if errInfo, ok := rsp["status_info"].(map[string]interface{}); ok {
			if errInfo["message"] != nil {
				detail["message"] = errInfo["message"]
			}
		}
		if detail != nil {
			detailByte, _ = json.Marshal(detail)
		} else {
			detailByte = nil
		}
		prefix := utils.GetUrlPrefix(c.Request.RequestURI)
		model := code.Model{}.Choices(prefix).Desc()
		global.DB.Table(table.OperateLog).Create(&request.OperateLogCreate{
			UserId:     userId,
			ActionName: actionName,
			Module:     model,
			AccessUrl:  accessUrl,
			RequestIp:  c.ClientIP(),

			UserAgent: userAgent,

			Success:    success,
			Detail:     detailByte,
			AccessTime: field.CustomTime(time.Now()),
		})

	}
}
