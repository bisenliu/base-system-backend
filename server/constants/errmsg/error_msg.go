package errmsg

import "errors"

// 自定义异常信息
var (
	QueryFailed   = errors.New("查询失败")
	SaveFailed    = errors.New("保存失败")
	UpdateFailed  = errors.New("更新失败")
	DeleteFailed  = errors.New("删除失败")
	ReadFailed    = errors.New("读取失败")
	ParseFailed   = errors.New("解析失败")
	NotFound      = errors.New("不存在")
	Invalid       = errors.New("不合法")
	Incorrect     = errors.New("不正确")
	Required      = errors.New("不能为空")
	FileSizeRange = errors.New("文件大小不能超过 %d M")
	NotPrivilege  = errors.New("您没有(%s)的权限")

	LoginInvalid     = errors.New("未登陆或登陆状态已失效")
	AccPwdInvalid    = errors.New("账号或密码不正确,请重新输入")
	AccStop          = errors.New("账号被停用,请联系管理员启用")
	LoginOutLimit    = errors.New("账号或密码已错误%d次，请%d分钟后重试")
	ResetPwdFailed   = errors.New("无法重置状态为%s的账号密码")
	OnlyStopOrEnable = errors.New("只能停用或启用账号")

	JsonConvertFiled = errors.New("json转换失败")
	TimeCalcFiled    = errors.New("时间计算失败")
	DecryptFailed    = errors.New("字段解密失败")
	RequestLimit     = errors.New("您已达到系统的最大并发请求数,请稍后重试")
)
