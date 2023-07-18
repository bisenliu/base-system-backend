package config

type Key struct {
	SecretKey string `mapstructure:"secret_key" json:"secret_key" yaml:"secret_key"` // 项目秘钥
	AesKey    string `mapstructure:"aes_key" json:"aes_key" yaml:"aes_key"`          // aes加密key
}
