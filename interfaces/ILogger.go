package interfaces

// ILogger ...
type ILogger interface {
	GetLevel(i ...interface{}) string
	Trace(i ...interface{})
	Debug(i ...interface{})
	Info(i ...interface{})
	Warn(i ...interface{})
	Error(i ...interface{})
	Fatal(i ...interface{})
	Tracef(format string, i ...interface{})
	Debugf(format string, i ...interface{})
	Infof(format string, i ...interface{})
	Warnf(format string, i ...interface{})
	Errorf(format string, i ...interface{})
	Fatalf(format string, i ...interface{})
}
