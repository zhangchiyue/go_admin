package microlog

import (
	"errors"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"path/filepath"
	"time"
)

const (
	FatalLevelStr = "fatal"
	ErrorLevelStr = "error"
	WarnLevelStr  = "warn"
	InfoLevelStr  = "info"
	DebugLevelStr = "debug"
)

var levelStr2IntMap = map[string]logrus.Level{
	FatalLevelStr: logrus.FatalLevel,
	ErrorLevelStr: logrus.ErrorLevel,
	WarnLevelStr:  logrus.WarnLevel,
	InfoLevelStr:  logrus.InfoLevel,
	DebugLevelStr: logrus.DebugLevel,
}

func levelStr2Int(level string) logrus.Level {
	if l, ok := levelStr2IntMap[level]; ok {
		return l
	}
	return logrus.DebugLevel
}

type Logger interface {
	Debugf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Printf(format string, args ...interface{})
	Warnf(format string, args ...interface{})
	Warningf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Fatalf(format string, args ...interface{})
	Panicf(format string, args ...interface{})

	Debug(args ...interface{})
	Info(args ...interface{})
	Print(args ...interface{})
	Warn(args ...interface{})
	Warning(args ...interface{})
	Error(args ...interface{})
	Fatal(args ...interface{})
	Panic(args ...interface{})
}

type Ops struct {
	MaxAge       int    `toml:"max_age"`
	Path         string `toml:"path"`
	Format       string `toml:"format"`
	Level        string `toml:"level"`
	ReportCaller bool   `toml:"report_caller"`
	Stdout       bool   `toml:""`
}

func NewMicroLog(ops ...*Ops) (Logger, error) {
	return newLogrus(ops...)
}

func newLogrus(opsList ...*Ops) (*logrus.Logger, error) {
	log := logrus.New()
	ops := &Ops{}
	if len(opsList) > 0 {
		ops = opsList[0]
	}
	formatter := GetFormatter(ops.Format)
	if formatter == nil {
		return nil, ErrFormatterNotFound
	}
	if ops.MaxAge <= 0 {
		ops.MaxAge = 7 * 24
	}
	log.SetFormatter(GetFormatter(ops.Format))
	log.SetLevel(levelStr2Int(ops.Level))
	log.SetReportCaller(ops.ReportCaller)
	if len(ops.Path) > 0 {
		absPath, _ := filepath.Abs(ops.Path)
		err := mkLogDir(absPath)
		if err != nil {
			return nil, errors.New("mkdir from log path error: " + err.Error())
		}
		writer, err := rotatelogs.New(
			absPath+".%Y-%m-%d-%H",
			rotatelogs.WithLinkName(absPath),                           // 生成软链，指向最新日志文件
			rotatelogs.WithMaxAge(time.Duration(ops.MaxAge)*time.Hour), // 文件最大保存时间
			rotatelogs.WithRotationTime(time.Minute),                   // 日志切割时间间隔
		)
		if err != nil {
			return nil, errors.New(err.Error() + " newLogrus " + absPath)
		}

		log.SetOutput(io.MultiWriter(os.Stdout, writer))
	}
	return log, nil
}

func mkLogDir(logPath string) error {
	dir, _ := filepath.Split(logPath)
	if len(dir) > 0 {
		err := os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			return err
		}
	}
	return nil
}
