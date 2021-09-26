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

type ServiceMetadata struct {
	Name    string
	Version string `yaml:"serviceVersion"`
	// define how service is deployed for dev/staging/live...
	RunTimes map[string]RunTime
	//optional
	Meta map[string]interface{}

	VersionHandles map[string]*VersionHandle
}

type RunTime struct {
	// define technology type (i.e. kubernetes, vm, serverless....)
	Technology string
	// define technology for dev deployment (i.e. skaffold&helm, skaffold & manifests, kubectl, helm-only,...)
	Helper map[string]interface{}
}

type VersionHandle struct {
	Path    string `yaml:"path"`
	Regexpr string `yaml:"regex"`
	Content []byte `yaml:"-"`
}
