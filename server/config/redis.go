package config

type Redis struct {
	Prefix      string `mapstructure:"prefix" ison:"prefix" yaml:"prefix"`                      // redis前缀
	Host        string `mapstructure:"host" ison:"host" yaml:"host"`                            //服务器地址：端口
	Port        string `mapstructure:"port" ison:"port" yaml:"port"`                            //端口
	Password    string `mapstructure:"password" json:"password" yaml:"password"`                //密码
	DefaultDb   int    `mapstructure:"default db" json:"default db" yaml:"default_db"`          //默认缓存
	TokenDb     int    `mapstructure:"token_db" json:"token_db" yaml:"token_db"`                // token缓存
	VerifyCodDb int    `mapstructure:"verify_cod_db" json:"verify_cod_db" yaml:"verify_cod_db"` //滑块验证缓存
}
