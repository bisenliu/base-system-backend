package config

type Avatar struct {
	Size int64 `mapstructure:"size" json:"size" yaml:"size"` // 头像大小限制(MB)
}
