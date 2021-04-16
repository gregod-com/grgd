package interfaces

// IProfile ...
type IProfile interface {
	SetID(id uint) error
	GetID() uint

	IsInitialized() bool
	SetInitialized(init bool) error

	GetName() string
	SetName(n string) error

	GetCurrentProject() IProject

	GetCurrentProjectID() uint
	SetCurrentProjectID(id uint) error

	GetMetaData(key string) string
	SetMetaData(key string, value string)

	GetBasePath() string
	GetUpdateURL() string
	GetProjects() map[string]IProject

	AddProject() error
	AddProjectByName(name string) error
	AddProjectDirect(p IProject) error

	RemoveProject(p IProject) error
	RemoveProjectByName(name string) error

	GetValues(i ...interface{}) []string
	GetValuesAsMap(i ...interface{}) map[string]string
	GetProjectsTable() [][]string
}
