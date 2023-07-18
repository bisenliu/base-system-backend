package core

import "flag"

func Env() (env string, systemInit bool) {
	flag.StringVar(&env, "env", "local", "请输入环境变量")
	flag.BoolVar(&systemInit, "system_init", false, "是否需要初始化默认数据")
	flag.Parse()
	return
}
