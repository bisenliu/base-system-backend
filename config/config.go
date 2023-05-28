package config

type Service struct {
	System    System    `mapstructure:"system" json:"system" yaml:"system"`
	Key       Key       `mapstructure:"key" json:"key" yaml:"key"`
	Token     Token     `mapstructure:"token" json:"token" vaml:"token"`
	Static    Static    `mapstructure:"static" json:"static" yaml:"static"`
	Redis     Redis     `mapstructure:"redis" json:"redis" yaml:"redis"`
	Avatar    Avatar    `mapstructure:"avatar" json:"avatar" yaml:"avatar"`
	Zap       Zap       `mapstructure:"zap" json:"zap" yaml:"zap"`
	Pgsql     Pgsql     `mapstructure:"pgsql" json:"pgsql" yaml:"pgsql"`
	SnowFlake SnowFlake `mapstructure:"snow_flake" json:"snow_flake" yaml:"snow_flake"`
}
