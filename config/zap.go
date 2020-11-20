package config

//构造 zap 日志结构体, 配置文件为 config.yaml
type ZapConf struct {
	Level         string `yaml:"level"`
	Format        string `yaml:"format"`
	Prefix        string `yaml:"prefix"`
	Director      string `yaml:"director"`
	LinkName      string `yaml:"link-name"`
	ShowLine      bool   `yaml:"show-line"`
	EncodeLevel   string `yaml:"encode-level"`
	StacktraceKey string `yaml:"stacktrace-key"`
	LogInConsole  bool   `yaml:"log-in-console"`
}
