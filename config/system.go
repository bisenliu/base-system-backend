package config

type System struct {
	Env       string `mapstructure:"env" json:"env" yaml:"env"`                      // 系统环境变量
	Version   string `mapstructure:"version" json:"version" yaml:"version"`          // 项目版本号
	StartTime string `mapstructure:"start_time" json:"start_time" yaml:"start_time"` // 项目开始时间
	Port      int64  `mapstructure:"port" json:"port" yaml:"port"`                   // 项目端口
}
