package interfaces

// IService ...
type IService interface {
	GetName() string
	GetPath(i ...interface{}) string
	SetPath(path string, i ...interface{}) error
	GetActive(i ...interface{}) bool
	SetActive(active bool, i ...interface{}) error
	ToggleActive(i ...interface{}) error
	GetValues(i ...interface{}) []string
}
