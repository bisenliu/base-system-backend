package user

type (
	AccStatus int
)

// 账号状态枚举
const (
	AccNormal AccStatus = iota
	AccFreeze
	AccChangePwd
	AccStop
)

// IsValid
//  @Description: 账号枚举校验,配合自定义 validator
//  @receiver s 接收者
//  @return bool 是否校验通过

func (s AccStatus) IsValid() bool {
	switch s {
	case AccNormal, AccFreeze, AccChangePwd, AccStop:
		return true
	}
	return false
}

// Choices
//  @Description: 枚举转换
//  @receiver s 接收者
//  @return string 对应的描述

func (s AccStatus) Choices() string {
	switch s {
	case AccNormal:
		return "正常"
	case AccFreeze:
		return "冻结"
	case AccChangePwd:
		return "需要修改密码"
	case AccStop:
		return "需要修改密码"
	default:
		return "Unknown"
	}
}
