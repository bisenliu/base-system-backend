package request

import "base-system-backend/constants"

type CaptchaType struct {
	Type constants.CaptchaType `json:"captcha_type" binding:"enum" label:"验证码类型"`
}

type CaptchaParam struct {
	Token     string `json:"token" binding:"required" label:"token"`
	PointJson string `json:"point_json" binding:"required" label:"滑块轨迹"`
	CaptchaType
}
