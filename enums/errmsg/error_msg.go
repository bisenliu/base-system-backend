package errmsg

import "errors"

var (
	QueryFailed   = errors.New("查询失败")
	SaveFailed    = errors.New("保存失败")
	UpdateFailed  = errors.New("更新失败")
	DeleteFailed  = errors.New("删除失败")
	ReadFailed    = errors.New("读取失败")
	NotFound      = errors.New("不存在")
	Exists        = errors.New("已存在")
	Invalid       = errors.New("不合法")
	Incorrect     = errors.New("不正确")
	Required      = errors.New("不能为空")
	FileSizeRange = errors.New("文件大小不能超过 %d M")
	NotPrivilege  = errors.New("您没有(%s)的权限")

	LoginInvalid   = errors.New("未登录或登录状态已失效")
	AccPwdInvalid  = errors.New("账号或密码不正确,请重新输入")
	AccStop        = errors.New("账号被停用,请联系管理员启用")
	LoginOutLimit  = errors.New("账号或密码已错误%d次，请%d分钟后重试")
	ResetPwdFailed = errors.New("无法重置状态为%s的账号密码")

	JsonConvertFiled = errors.New("json转换失败")
	TimeCalcFiled    = errors.New("时间计算失败")
)
