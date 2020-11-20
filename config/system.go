package config

type SystemConf struct {
	Env           string `yaml:"env"`
	Addr          int    `yaml:"addr"`
	DbType        string `yaml:"db-type"`
	OssType       string `yaml:"oss-type"`
	UseMultipoint bool   `yaml:"use-multipoint"`
}
