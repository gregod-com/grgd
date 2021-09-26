package interfaces

// IProject ...
type IProject interface {
	Init(core ICore) error
	GetID() uint
	SetID(id uint) error

	IsInitialized() bool
	SetInitialized(init bool) error

	GetName() string
	SetName(n string) error

	GetPath() string
	SetPath(path string, i ...interface{}) error
	GetServices(i ...interface{}) map[string]ServiceMetadata
	GetServiceByName(serviceName string, i ...interface{}) ServiceMetadata
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
	Services      map[string]ServiceLocator
	IgnoreFolders []string
	// optional
	Meta map[string]interface{}
}

type ServiceLocator struct {
	Active bool
	Path   string
}
