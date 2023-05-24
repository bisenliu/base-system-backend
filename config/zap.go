package config

type Zap struct {
	Level      string `mapstructure:"level" json:"level" yaml:"level"`                   // 级别
	Director   string `mapstructure:"director" json:"director"  yaml:"director"`         // 日志文件夹
	MaxAge     int    `mapstructure:"max_age" json:"max_age" yaml:"max_age"`             // 日志留存时间
	MaxSize    int    `mapstructure:"max_size" json:"max_size" yaml:"max_size"`          // 单个日志最大容量
	MaxBackups int    `mapstructure:"max_backups" json:"max_backups" yaml:"max_backups"` // 最大备份数量
}
