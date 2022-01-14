package logger

import (
	"runtime/debug"
	"strings"
)

type Logger interface {
	Trace(args ...interface{})
	Debug(args ...interface{})
	Info(args ...interface{})
	Warn(args ...interface{})
	Warning(args ...interface{})
	Error(args ...interface{})
	Fatal(args ...interface{})
	Panic(args ...interface{})
	GetPrefix() string
}
type Struct struct {
	prefix string
}

func (s Struct) GetPrefix() string {
	return s.prefix
}

func (s Struct) Trace(args ...interface{}) {
	args = append([]interface{}{s.prefix}, args...)
	loger.Trace(args...)
}

func (s Struct) Debug(args ...interface{}) {
	args = append([]interface{}{s.prefix}, args...)
	loger.Debug(args...)
}

func (s Struct) Info(args ...interface{}) {
	args = append([]interface{}{s.prefix}, args...)
	loger.Info(args...)
}

func (s Struct) Warn(args ...interface{}) {
	args = append([]interface{}{s.prefix}, args...)
	loger.Warn(args...)
}

func (s Struct) Warning(args ...interface{}) {
	args = append([]interface{}{s.prefix}, args...)
	loger.Warning(args...)
}

func (s Struct) Error(args ...interface{}) {
	st := string(debug.Stack())
	lines := strings.Split(st, "\n")
	args = append(args, ":", lines[len(lines)-2])
	args = append([]interface{}{s.prefix}, args...)
	loger.Error(args...)
}

func (s Struct) Fatal(args ...interface{}) {
	args = append([]interface{}{s.prefix}, args...)
	args = append(args, string(debug.Stack()))
	loger.Fatal(args...)
}

func (s Struct) Panic(args ...interface{}) {
	args = append([]interface{}{s.prefix}, args...)
	loger.Panic(args...)
}

func NewLogger(prefix string) Logger {
	return Struct{
		prefix: prefix,
	}
}
