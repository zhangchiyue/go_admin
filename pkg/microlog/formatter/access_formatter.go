package formatter

import (
	"bytes"
	"github.com/sirupsen/logrus"
	"strconv"
)

type AccessFormatter struct{}

func (f *AccessFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	b := bytes.Buffer{}
	b.Grow(len(entry.Message) + 128)
	b.WriteString(entry.Time.Format("[2006-01-02 15:04:05] ["))
	b.WriteString(entry.Level.String())
	if entry.Caller != nil {
		b.WriteString("] [")
		l := len(entry.Caller.File)
		s := 0
		if l > 50 {
			s = l - 50
		}
		b.WriteString(entry.Caller.File[s:l])
		b.WriteString(":")
		b.WriteString(strconv.Itoa(entry.Caller.Line))
	}
	b.WriteString("] ")
	b.WriteString(entry.Message)
	b.WriteString("\n")
	return b.Bytes(), nil
}
