package formatter

import (
	"bytes"
	"github.com/sirupsen/logrus"
)

type TimeFormatter struct{}

func (f *TimeFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	b := bytes.Buffer{}
	b.Grow(len(entry.Message) + 32)
	b.WriteString(entry.Time.Format("[2006-01-02 15:04:05]\t"))
	b.WriteString(entry.Message)
	b.WriteString("\n")
	return b.Bytes(), nil
}
