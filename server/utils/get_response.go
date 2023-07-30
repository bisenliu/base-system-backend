package utils

import (
	"base-system-backend/constants"
	"base-system-backend/global"
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
	"strings"
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
	parts := strings.Split(strings.Join(strings.Split(funcName, "-fm"), ""), "/")
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

func GetResponseData(c *gin.Context) (success bool, statusInfo interface{}) {
	var rsp map[string]interface{}
	method := c.Request.Method
	blw := &bodyLogWriter{body: bytes.NewBuffer([]byte{}), ResponseWriter: c.Writer}
	c.Writer = blw
	c.Next()

	contentType := c.Writer.Header().Get("Content-Type")
	success = true
	if contentType != constants.ExcelContentType && blw != nil {
		if err := json.Unmarshal(blw.body.Bytes(), &rsp); err != nil {
			global.LOG.Error("body convert json failed:%w ", zap.Error(err))
			c.Abort()
			return
		}
	}
	status, ok := rsp["status"].(float64)
	if ok && status == 0 && method == "GET" {
		c.Abort()
		return
	}
	success = status == 0
	statusInfo = rsp["status_info"]
	return
}
