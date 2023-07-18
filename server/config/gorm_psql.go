package config

type Pgsql struct {
	Host         string `mapstructure:"host" json:"host" yaml:"host"`                               // 服务器地址:端口
	Port         string `mapstructure:"port" json:"port" yaml:"port"`                               //:端口
	Dbname       string `mapstructure:"db_name" json:"db_name" yaml:"db_name"`                      // 数据库名
	Username     string `mapstructure:"username" json:"username" yaml:"username"`                   // 数据库用户名
	Password     string `mapstructure:"password" json:"password" yaml:"password"`                   // 数据库密码
	MaxIdleConns int    `mapstructure:"max-idle_conns" json:"max-idle_conns" yaml:"max-idle_conns"` // 空闲中的最大连接数
	MaxOpenConns int    `mapstructure:"max-open_conns" json:"max-open_conns" yaml:"max-open_conns"` // 打开到数据库的最大连接数

}

func (p *Pgsql) Dsn() string {
	return "host=" + p.Host + " user=" + p.Username + " password=" + p.Password + " dbname=" + p.Dbname + " port=" + p.Port
}
