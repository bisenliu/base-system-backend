package config

type Token struct {
	ExpiredTime float64 `mapstructure:"expired_time" yaml:"expired_time" json:"expired_time"` // 过期时间
}
