package core

import (
	"grgd/interfaces"
	"log"
	"os"
)

// ProvideDefaultLogger ....
func ProvideDefaultLogger() interfaces.ILogger {
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

// Tracef ...
func (l *DeafultLogger) Tracef(format string, i ...interface{}) {
	l.logger.Printf(format, i...)
}

// Debugf ...
func (l *DeafultLogger) Debugf(format string, i ...interface{}) {
	l.logger.Printf(format, i...)
}

// Infof ...
func (l *DeafultLogger) Infof(format string, i ...interface{}) {
	l.logger.Printf(format, i...)
}

// Warnf ...
func (l *DeafultLogger) Warnf(format string, i ...interface{}) {
	l.logger.Printf(format, i...)
}

// Errorf ...
func (l *DeafultLogger) Errorf(format string, i ...interface{}) {
	l.logger.Printf(format, i...)
}

// Fatalf ...
func (l *DeafultLogger) Fatalf(format string, i ...interface{}) {
	l.logger.Fatalf(format, i...)
}
