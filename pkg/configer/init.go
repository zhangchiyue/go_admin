package configer

import (
	"fmt"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

//var Conf Config

var _ins = viper.New()

func Init(flags *pflag.FlagSet) (*Config, error) {
	if err := _ins.BindPFlags(flags); err != nil {
		return nil, err
	}
	_ins.SetConfigName(_ins.GetString("env"))
	_ins.SetConfigType("yaml")
	_ins.AddConfigPath("config/")
	if err := _ins.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("error on parsing configuration file: %v", err)
	}
	var Conf *Config
	if err := _ins.Unmarshal(&Conf); err != nil {
		return nil, fmt.Errorf("error on unmarshal configuration file: %v", err)
	}
	return Conf, nil
}
