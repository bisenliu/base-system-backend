package main

import (
	"base-system-backend/core"
	"base-system-backend/global"
	"base-system-backend/initialize"
	"embed"
	"fmt"
	"go.uber.org/zap"
)

//go:embed initialize/internal/privilege.json
//go:embed version.txt
var f embed.FS

// @title 基础后端框架
// @version 1.0
// @description 应用于快速搭建后端
// @termsOfService http://swagger.io/terms/

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host http://127.0.0.1:8001
// @BasePath /v1/
func main() {
	// build 时把静态文件打包到二进制中 []uint64
	global.FS = f
	// 环境变量
	global.ENV, global.SystemInit = core.Env()
	// 初始化viper
	global.VP = core.Viper()
	//初始化日志库
	global.LOG = core.Zap()
	zap.ReplaceGlobals(global.LOG)
	// 初始化雪花算法
	global.Node = initialize.SnowFlake()
	// 初始化gorm连接
	global.DB = initialize.GormPgSql()
	// 初始化表
	initialize.RegisterTables()
	//程序结束前关闭数据库链援
	defer initialize.CloseDB()
	initialize.DefaultDataInit()
	// 初始化redis
	initialize.Redis()
	defer initialize.CloseRedis()
	// 初始化翻译器
	if err := initialize.InitTrans("zh"); err != nil {
		panic(fmt.Errorf("init trans failed: %s", err.Error()))
	}
	core.RunServer()
}
