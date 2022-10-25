package logger

import (
	"fmt"
	"log"
	"os"
	"runtime"

	"github.com/gregod-com/grgd/interfaces"
)

var log_level = 4 // level info

// ProvideLogger ....
func ProvideLogger() interfaces.ILogger {

	logger := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	if checkFlag("debug") || checkFlag("d") {
		log_level = 5
	}
	switch checkFlagArg("log-level") {
	case "trace":
		log_level = 6
	case "debug":
		log_level = 5
	case "info":
		log_level = 4
	case "warn":
		log_level = 3
	case "error":
		log_level = 2
	case "fatal":
		log_level = 1
	case "panic":
		log_level = 0
	default:
		log_level = 4
	}
	wrapped_logger := &Logger{logger: logger}
	wrapped_logger.Tracef("provide %T", logger)
	return wrapped_logger
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

// Logger ...
type Logger struct {
	logger *log.Logger
	pkg    string
}

// GetLevel ...
func (l *Logger) GetLevel(i ...interface{}) string {
	switch log_level {
	case 0:
		return "panic"
	case 1:
		return "fatal"
	case 2:
		return "error"
	case 3:
		return "warn"
	case 4:
		return "info"
	case 5:
		return "debug"
	case 6:
		return "trace"
	default:
		return "info"
	}

	return l.logger.Prefix()
}

// Trace ...
func (l *Logger) Trace(i ...interface{}) {
	pc := make([]uintptr, 15)
	n := runtime.Callers(2, pc)
	frames := runtime.CallersFrames(pc[:n])
	frame, _ := frames.Next()
	format := fmt.Sprintf("[%s] %s", frame.Function, i)
	if log_level > 5 {
		l.logger.SetPrefix("TRACE: ")
		l.logger.Println(format)
	}
}

// Debug ...
func (l *Logger) Debug(i ...interface{}) {
	if log_level > 4 {
		l.logger.SetPrefix("DEBUG: ")
		l.logger.Println(i...)
	}
}

// Info ...
func (l *Logger) Info(i ...interface{}) {
	if log_level > 3 {
		l.logger.SetPrefix("INFO: ")
		l.logger.Println(i...)
	}
}

// Warn ...
func (l *Logger) Warn(i ...interface{}) {
	if log_level > 2 {
		l.logger.SetPrefix("WARN: ")
		l.logger.Println(i...)
	}
}

// Error ...
func (l *Logger) Error(i ...interface{}) {
	if log_level > 1 {
		l.logger.SetPrefix("ERROR: ")
		l.logger.Println(i...)
	}
}

// Fatal ...
func (l *Logger) Fatal(i ...interface{}) {
	l.logger.SetPrefix("FATAL: ")
	l.logger.Fatal(i...)
}

// Tracef ...
func (l *Logger) Tracef(format string, i ...interface{}) {
	pc := make([]uintptr, 15)
	n := runtime.Callers(2, pc)
	frames := runtime.CallersFrames(pc[:n])
	frame, _ := frames.Next()
	format = fmt.Sprintf("[%s] %s", frame.Function, format)
	if log_level > 5 {
		l.logger.SetPrefix("TRACE: ")
		l.logger.Printf(format)
	}
}

// Debugf ...
func (l *Logger) Debugf(format string, i ...interface{}) {
	if log_level > 4 {
		l.logger.SetPrefix("DEBUG: ")
		l.logger.Printf(format, i...)
	}
}

// Infof ...
func (l *Logger) Infof(format string, i ...interface{}) {
	if log_level > 3 {
		l.logger.SetPrefix("INFO: ")
		l.logger.Printf(format, i...)
	}
}

// Warnf ...
func (l *Logger) Warnf(format string, i ...interface{}) {
	if log_level > 2 {
		l.logger.SetPrefix("WARN: ")
		l.logger.Printf(format, i...)
	}
}

// Errorf ...
func (l *Logger) Errorf(format string, i ...interface{}) {
	if log_level > 1 {
		l.logger.SetPrefix("ERROR: ")
		l.logger.Printf(format, i...)
	}
}

// Fatalf ...
func (l *Logger) Fatalf(format string, i ...interface{}) {
	l.logger.SetPrefix("FATAL: ")
	l.logger.Fatalf(format, i...)
}
