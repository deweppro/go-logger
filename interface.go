package goLogger

import (
	"io"
)

type Logger interface {
	SetOutput(out io.Writer)
	Infof(format string, args ...interface{})
	Warnf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Debugf(format string, args ...interface{})
}
