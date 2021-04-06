package logger

import (
	"fmt"
	"os"
	"runtime"

	"github.com/gregod-com/grgd/interfaces"

	"github.com/sirupsen/logrus"
)

// ProvideLogrusLogger ....
func ProvideLogrusLogger() interfaces.ILogger {
	logger := logrus.New()
	if checkFlag("debug") || checkFlag("d") {
		logger.SetLevel(logrus.DebugLevel)
	}
	switch checkFlagArg("log-level") {
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
	logrusLogger.Tracef("provide %T", logrusLogger)
	return logrusLogger
}

// CheckFlagArg ...
func checkFlagArg(flag string) string {
	for k, v := range os.Args {
		if v == "--"+flag && len(os.Args) > k+1 {
			return os.Args[k+1]
		}
	}
	return ""
}

// CheckFlag ...
func checkFlag(flag string) bool {
	for _, v := range os.Args {
		if v == "-"+flag {
			return true
		}
		if v == "--"+flag {
			return true
		}
	}
	return false
}

// LogrusLogger ...
type LogrusLogger struct {
	logger *logrus.Logger
	pkg    string
}

// GetLevel ...
func (l *LogrusLogger) GetLevel(i ...interface{}) string {
	return l.logger.GetLevel().String()
}

// Trace ...
func (l *LogrusLogger) Trace(i ...interface{}) {
	pc := make([]uintptr, 15)
	n := runtime.Callers(2, pc)
	frames := runtime.CallersFrames(pc[:n])
	frame, _ := frames.Next()
	format := fmt.Sprintf("[%s] %s", frame.Function, i)
	l.logger.Trace(format)
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

// Tracef ...
func (l *LogrusLogger) Tracef(format string, i ...interface{}) {
	pc := make([]uintptr, 15)
	n := runtime.Callers(2, pc)
	frames := runtime.CallersFrames(pc[:n])
	frame, _ := frames.Next()
	format = fmt.Sprintf("[%s] %s", frame.Function, format)
	l.logger.Tracef(format, i...)
}

// Debugf ...
func (l *LogrusLogger) Debugf(format string, i ...interface{}) {
	l.logger.Debugf(format, i...)
}

// Infof ...
func (l *LogrusLogger) Infof(format string, i ...interface{}) {
	l.logger.Infof(format, i...)
}

// Warnf ...
func (l *LogrusLogger) Warnf(format string, i ...interface{}) {
	l.logger.Warnf(format, i...)
}

// Errorf ...
func (l *LogrusLogger) Errorf(format string, i ...interface{}) {
	l.logger.Errorf(format, i...)
}

// Fatalf ...
func (l *LogrusLogger) Fatalf(format string, i ...interface{}) {
	l.logger.Fatalf(format, i...)
}
