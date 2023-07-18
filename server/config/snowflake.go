package config

type SnowFlake struct {
	StartTime string `mapstructure:"start_time" json:"start_time" 'yaml:"start_time"` // 开始时间(不能比当前时间大)
	MachineId int64  `mapstructure:"machine_id" json:"machine_id" 'yaml:"machine_id"` // 当前机器编号
}
