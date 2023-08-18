package constants

import constant "github.com/TestsLing/aj-captcha-go/const"

type CaptchaType int

// 账号状态枚举
const (
	BlockPuzzle CaptchaType = iota
	ClickWord
)

// IsValid
//  @Description: 滑块类型校验,配合自定义 validator
//  @receiver c 接收者
//  @return bool 是否校验通过

func (c CaptchaType) IsValid() bool {
	switch c {
	case BlockPuzzle, ClickWord:
		return true
	}
	return false
}

// String
//  @Description: 返回对应的滑块类型 key
//  @receiver c 接收者
//  @return string 滑块 key

func (c CaptchaType) String() string {
	switch c {
	case ClickWord:
		return constant.ClickWordCaptcha
	default:
		return constant.BlockPuzzleCaptcha
	}
}
