package log

import (
	"adx-admin/pkg/configer"
	"adx-admin/pkg/microlog"
)

func InitLog(conf *configer.Config) (microlog.Logger, error) {
	logConfig := conf.LogConfig
	runtimeLog, err := microlog.NewMicroLog(logConfig.RuntimeOption)
	if err != nil {
		return nil, err
	}
	return runtimeLog, nil
}
