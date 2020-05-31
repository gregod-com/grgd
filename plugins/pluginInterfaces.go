package grgdplugin

// Greeter ...
type Greeter interface {
	Greet()
}

// GrgdPlugin ...
type GrgdPlugin interface {
	Init(i interface{}) interface{}
	Version(i interface{}) interface{}
	Author(i interface{}) interface{}
	Category(i interface{}) interface{}
	Methods(i interface{}) map[string]func(interface{}) interface{}
}
