package interfaces

// IProject ...
type IProject interface {
	Init(core ICore) error
	GetID(i ...interface{}) uint
	SetID(id uint) error

	IsInitialized() bool
	SetInitialized(init bool) error

	GetName() string
	SetName(n string) error

	GetPath() string
	SetPath(path string, i ...interface{}) error
	GetServices(i ...interface{}) map[string]IService
	GetServiceByName(serviceName string, i ...interface{}) IService
	GetValues(i ...interface{}) []string

	SetSettingsYamlPath(path string, i ...interface{}) error
	GetSettingsYamlPath(i ...interface{}) string

	WriteSettingsObject(h IHelper, i ...interface{}) error
	ReadSettingsObject(h IHelper, i ...interface{}) (*ProjectMetadata, error)
}

type ProjectMetadata struct {
	Name    string
	Version string
	// only service-name and ref to service.yaml
	Services map[string]interface{}
	// optional
	Meta map[string]interface{}
}
