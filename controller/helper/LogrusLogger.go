package helper

import (
	"grgd/interfaces"

	"github.com/sirupsen/logrus"
)

// CreateLogger ....
func CreateLogger(h interfaces.IHelper) interfaces.ILogger {
	logger := logrus.New()
	if h.CheckFlag("debug") || h.CheckFlag("d") {
		logger.SetLevel(logrus.DebugLevel)
		return logger
	}
	switch h.CheckFlagArg("log-level") {
	case "trace":
		logger.SetLevel(logrus.TraceLevel)
	case "debug":
		logger.SetLevel(logrus.DebugLevel)
	case "info":
		logger.SetLevel(logrus.InfoLevel)
	case "warn":
		logger.SetLevel(logrus.WarnLevel)
	case "error":
		logger.SetLevel(logrus.ErrorLevel)
	case "fatal":
		logger.SetLevel(logrus.FatalLevel)
	case "panic":
		logger.SetLevel(logrus.PanicLevel)
	default:
		logger.SetLevel(logrus.InfoLevel)
	}
	l := &Logger{logger}
	return l
}

// Logger ...
type Logger struct {
	logger *logrus.Logger
}

// Trace ...
func (l *Logger) Trace(i ...interface{}) {
	l.logger.Trace(i...)
}

// Debug ...
func (l *Logger) Debug(i ...interface{}) {
	l.logger.Debug(i...)
}

// Info ...
func (l *Logger) Info(i ...interface{}) {
	l.logger.Info(i...)
}

// Warn ...
func (l *Logger) Warn(i ...interface{}) {
	l.logger.Warn(i...)
}

// Error ...
func (l *Logger) Error(i ...interface{}) {
	l.logger.Error(i...)
}

// Fatal ...
func (l *Logger) Fatal(i ...interface{}) {
	l.logger.Fatal(i...)
}
