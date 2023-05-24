package errmsg

import "errors"

var (
	QueryFailed = errors.New("查询失败")

	LoginOutLimit = errors.New("账号或密码已错误%d次，请%d分钟后重试")

	JsonConvertFiled = errors.New("Json转换失败")
)
