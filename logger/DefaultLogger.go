package logger

import (
	"grgd/interfaces"
	"log"
	"os"
)

// ProvideDefaultLogger ....
func ProvideDefaultLogger(h interfaces.IHelper) interfaces.ILogger {
	logger := log.New(os.Stderr, "", 0)
	defaultLogger := &DeafultLogger{logger: logger}
	return defaultLogger
}

// DeafultLogger ...
type DeafultLogger struct {
	logger *log.Logger
}

// Trace ...
func (l *DeafultLogger) Trace(i ...interface{}) {
	l.logger.Println(i...)
}

// Debug ...
func (l *DeafultLogger) Debug(i ...interface{}) {
	l.logger.Println(i...)
}

// Info ...
func (l *DeafultLogger) Info(i ...interface{}) {
	l.logger.Println(i...)
}

// Warn ...
func (l *DeafultLogger) Warn(i ...interface{}) {
	l.logger.Println(i...)
}

// Error ...
func (l *DeafultLogger) Error(i ...interface{}) {
	l.logger.Println(i...)
}

// Fatal ...
func (l *DeafultLogger) Fatal(i ...interface{}) {
	l.logger.Fatal(i...)
}
