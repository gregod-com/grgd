package logger

import (
	"grgd/interfaces"

	"github.com/sirupsen/logrus"
)

// ProvideLogrusLogger ....
func ProvideLogrusLogger(h interfaces.IHelper) interfaces.ILogger {
	// h, _ := i[0].(interfaces.IHelper)

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
	logrusLogger := &LogrusLogger{logger: logger}
	return logrusLogger
}

// LogrusLogger ...
type LogrusLogger struct {
	logger *logrus.Logger
}

// Trace ...
func (l *LogrusLogger) Trace(i ...interface{}) {
	l.logger.Trace(i...)
}

// Debug ...
func (l *LogrusLogger) Debug(i ...interface{}) {
	l.logger.Debug(i...)
}

// Info ...
func (l *LogrusLogger) Info(i ...interface{}) {
	l.logger.Info(i...)
}

// Warn ...
func (l *LogrusLogger) Warn(i ...interface{}) {
	l.logger.Warn(i...)
}

// Error ...
func (l *LogrusLogger) Error(i ...interface{}) {
	l.logger.Error(i...)
}

// Fatal ...
func (l *LogrusLogger) Fatal(i ...interface{}) {
	l.logger.Fatal(i...)
}
