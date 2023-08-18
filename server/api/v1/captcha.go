package v1

import (
	"base-system-backend/constants/code"
	"base-system-backend/global"
	"base-system-backend/model/captcha/request"
	rsp "base-system-backend/model/captcha/response"
	"base-system-backend/model/common/response"
	"base-system-backend/utils/validate"
	"github.com/gin-gonic/gin"
)

type CaptchaApi struct{}

// CaptchaGetApi
// @Summary 获取滑块信息
// @Description 获取滑块信息
// @Tags CaptchaApi
// @Accept application/json
// @Produce application/json
// @Param object body request.CaptchaType true "滑块类型(文字/图片)"
// @Security ApiKeyAuth
// @Success 200 {object} response.CaptchaInfo
// @Router /captcha/get/ [post]
func (CaptchaApi) CaptchaGetApi(c *gin.Context) {
	params := new(request.CaptchaType)
	if !validate.RequestParamsVerify(c, params) {
		return
	}
	// 根据参数类型获取不同服务即可
	data, err := global.CAPTCHA.GetService(params.Type.String()).Get()
	if err != nil {
		response.Error(c, code.InvalidParams, err, err.Error())
		return
	}
	captcha := &rsp.CaptchaInfo{
		JigsawImageBase64:   data["jigsawImageBase64"],
		OriginalImageBase64: data["originalImageBase64"],
		Token:               data["token"],
		SecretKey:           data["secretKey"],
	}

	response.OK(c, captcha)
}

// CaptchaCheckApi
// @Summary 校验滑块轨迹/文字
// @Description 校验滑块轨迹/文字
// @Tags CaptchaApi
// @Accept application/json
// @Produce application/json
// @Param object body request.CaptchaParam true "校验滑块轨迹/文字"
// @Security ApiKeyAuth
// @Success 200 {object} response.Data
// @Router /captcha/check/ [post]
func (CaptchaApi) CaptchaCheckApi(c *gin.Context) {
	params := new(request.CaptchaParam)
	if !validate.RequestParamsVerify(c, params) {
		return
	}
	ser := global.CAPTCHA.GetService(params.Type.String())

	if err := ser.Check(params.Token, params.PointJson); err != nil {
		response.Error(c, code.InvalidParams, err, err.Error())
		return
	}
	response.OK(c, nil)
}
