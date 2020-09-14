package interfaces

// IProject ...
type IProject interface {
	GetName() string
	GetID(i ...interface{}) uint
	GetPath(i ...interface{}) string
	SetPath(path string, i ...interface{}) error
	GetServices(i ...interface{}) map[string]IService
	GetServiceByName(serviceName string, i ...interface{}) IService
	GetValues(i ...interface{}) []string
}
