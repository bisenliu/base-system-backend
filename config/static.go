package config

type Static struct {
	Base   string   `mapstructure:"base" json:"base" yaml:"base"`       // 静态文件根目录
	Avatar []string `mapstructure:"avatar" json:"avatar" yaml:"avatar"` // 头像文件目录

}
