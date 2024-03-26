package formatter

import (
	"bytes"
	"github.com/sirupsen/logrus"
)

type VoidFormatter struct{}

func (f *VoidFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	b := bytes.Buffer{}
	b.Grow(len(entry.Message) + 32)
	b.WriteString(entry.Message)
	b.WriteString("\n")
	return b.Bytes(), nil
}
