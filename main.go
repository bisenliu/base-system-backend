package main

import (
	"base-system-backend/core"
	"base-system-backend/global"
	"base-system-backend/initialize"
	"go.uber.org/zap"
)

// @title 基础后端框架
// @version 1.0
// @description 应用于快速搭建后端
// @termsOfService http://swagger.io/terms/

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host http://127.0.0.1:8001
// @BasePath /v1/
func main() {
	// 环境变量
	global.ENV = core.Env()
	// 初始化viper
	global.VP = core.Viper()
	//初始化日志库
	global.LOG = core.Zap()
	zap.ReplaceGlobals(global.LOG)
	// 初始化gorm连接
	global.DB = initialize.GormPgSql()
	if global.DB != nil {
		// 初始化表
		initialize.RegisterTables()
		//程序结束前关闭数据库链援
		defer initialize.CloseDB()
	}
	// 初始化redis
	initialize.Redis()
	defer initialize.CloseRedis()
	// 初始化雪花算法
	global.Node = initialize.SnowFlake()
	// 初始化翻译器
	if err := initialize.InitTrans("zh"); err != nil {
		global.LOG.Error("init trans failed:", zap.Error(err))
		return
	}
	core.RunServer()
}
