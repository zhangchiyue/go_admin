package microlog

import (
	"adx-admin/pkg/microlog/formatter"
	"errors"
	"github.com/sirupsen/logrus"
	"reflect"
)

var formatterMap map[string]logrus.Formatter
var ErrNameAlreadyExist = errors.New("formatter name already exist")
var ErrFormatterNotFound = errors.New("formatter not found")

func init() {
	formatterMap = make(map[string]logrus.Formatter, 16)
	_ = Register(&formatter.AccessFormatter{})
	_ = Register(&formatter.TimeFormatter{})
	_ = Register(&formatter.VoidFormatter{})
	_ = Register(&formatter.RuntimeFormatter{})
}

func Register(formatter logrus.Formatter) error {
	name := getClassName(formatter)
	if _, ok := formatterMap[name]; ok {
		return ErrNameAlreadyExist
	}
	formatterMap[name] = formatter
	return nil
}

func getClassName(i interface{}) string {
	return reflect.TypeOf(i).Elem().Name()
}

func GetFormatter(name string) logrus.Formatter {
	if len(name) == 0 {
		return &formatter.AccessFormatter{}
	}
	return formatterMap[name]
}
