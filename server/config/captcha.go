package config

type Captcha struct {
	Expire    int    `mapstructure:"expire" json:"expire" yaml:"expire"`             // 过期时间
	WaterSeal string `mapstructure:"water_seal" json:"water_seal" yaml:"water_seal"` // 滑块图片水印
}
