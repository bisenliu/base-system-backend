package core

import "flag"

func Env() (env string) {
	flag.StringVar(&env, "env", "local", "请输入环境变量")
	flag.Parse()
	return
}
