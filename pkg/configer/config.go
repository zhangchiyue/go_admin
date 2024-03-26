package configer

import (
	"adx-admin/pkg/microlog"
	"time"
)

type Config struct {
	LogConfig         LogConfig         `mapstructure:"log_config"`
	HttpServerConfig  HttpServerConfig  `mapstructure:"http_server_config"`
	RedisServerConfig RedisServerConfig `mapstructure:"redis_server_config"`
	RedShiftConfig    RedshiftConfig    `mapstructure:"redshift_config"`
}

type RedshiftConfig struct {
	Dsn string `mapstructure:"dsn"`
}

type HttpServerConfig struct {
	HttpPort     int `mapstructure:"http_port"`
	ReadTimeout  int `mapstructure:"read_timeout"`
	WriteTimeout int `mapstructure:"write_timeout"`
}

type RedisServerConfig struct {
	Addrs        []string      `mapstructure:"addrs"`
	PassWord     string        `mapstructure:"password"`
	BaseTimeout  time.Duration `mapstructure:"base_timeout"`
	DialTimeout  time.Duration
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	MaxRetries   int
}

type LogConfig struct {
	RuntimeOption    *microlog.Ops `mapstructure:"runtime_option"`     // 运行时
	RepeatExtOption  *microlog.Ops `mapstructure:"repeat_ext_option"`  // 重复过期
	RequestOption    *microlog.Ops `mapstructure:"request_option"`     // 请求
	RequestCsvOption *microlog.Ops `mapstructure:"request_csv_option"` // 请求CSV
}
