package logger

import (
	"runtime/debug"
	"strings"

	"github.com/sirupsen/logrus"
)

// These are the different logging levels. You can set the logging level to log
// on your instance of logger, obtained with `logrus.New()`.
const (
	// PanicLevel level, highest level of severity. Logs and then calls panic with the
	// message passed to Debug, Info, ...
	PanicLevel Level = iota
	// FatalLevel level. Logs and then calls `logger.Exit(1)`. It will exit even if the
	// logging level is set to Panic.
	FatalLevel Level = 1
	// ErrorLevel level. Logs. Used for errors that should definitely be noted.
	// Commonly used for hooks to send errors to an error tracking service.
	ErrorLevel Level = 2
	// WarnLevel level. Non-critical entries that deserve eyes.
	WarnLevel Level = 3
	// InfoLevel level. General operational entries about what's going on inside the
	// application.
	InfoLevel Level = 4
	// DebugLevel level. Usually only enabled when debugging. Very verbose logging.
	DebugLevel Level = 5
	// TraceLevel level. Designates finer-grained informational events than the Debug.
	TraceLevel Level = 6
)

//Level -
type Level uint32

// iota is a predeclared identifier representing the untyped integer ordinal
// number of the current const specification in a (usually parenthesized)
// const declaration. It is zero-indexed.
const iota = 0 // Untyped int.

// логер библиотеки logrus
var loger *logrus.Logger

func init() {
	loger = logrus.New()
	loger.SetFormatter(&logrus.TextFormatter{
		ForceColors: true,
		//DisableQuote: true,
		FullTimestamp:    true,
		DisableTimestamp: false,
		TimestampFormat:  "2006.01.02 15:04:05",
	})
	loger.SetLevel(logrus.DebugLevel)
}

//SetLevel - установить уровень логов
func SetLevel(lvl Level) {
	loger.SetLevel(logrus.Level(lvl))
}

//Trace -
func Trace(args ...interface{}) {
	loger.Trace(args...)
}

//Debug -
func Debug(args ...interface{}) {
	loger.Debug(args...)
}

// Info -
func Info(args ...interface{}) {
	loger.Info(args...)
}

//Warn -
func Warn(args ...interface{}) {
	loger.Warn(args...)
}

//Warning -
func Warning(args ...interface{}) {
	loger.Warning(args...)
}

//Error -
func Error(args ...interface{}) {
	s := string(debug.Stack())
	lines := strings.Split(s, "\n")
	args = append(args, ":", lines[len(lines)-2])
	loger.Error(args...)
}

//Fatal -
func Fatal(args ...interface{}) {
	args = append(args, string(debug.Stack()))
	loger.Fatal(args...)
}

//Panic -
func Panic(args ...interface{}) {
	loger.Panic(args...)
}
