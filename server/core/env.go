package core

import "flag"

// Env
//  @Description: 初始化环境变量
//  @return env 环境变量
//  @return systemInit 是否需要初始化相关表(用户,角色,权限)

func Env() (env string, systemInit bool) {
	flag.StringVar(&env, "env", "local", "请输入环境变量")
	flag.BoolVar(&systemInit, "system_init", false, "是否需要初始化默认数据")
	flag.Parse()
	return
}
