package config

// 构造配置结构体，配置文件为 config.yaml
type Server struct {
	Zap      ZapConf      `yaml:"zap"`
	Postgres PostgresConf `yaml:"postgres"`
	Mysql MysqlConf `yaml:"mysql"`
	System   SystemConf   `yaml:"system"`
	Jwt      JwtConf      `yaml:"jwt"`
}
