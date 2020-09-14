package interfaces

type ILogger interface {
	Trace(i ...interface{})
	Debug(i ...interface{})
	Info(i ...interface{})
	Warn(i ...interface{})
	Error(i ...interface{})
	Fatal(i ...interface{})
}
