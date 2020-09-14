package interfaces

// IProfile ...
type IProfile interface {
	IsInitialized() bool
	GetName() string
	GetBasePath() string
	GetProjects() map[string]IProject
	AddProject(p string) error
	RemoveProject(p IProject) error
	RemoveProjectByName(p string) error
	GetCurrentProject() IProject
	SetCurrentProject(p IProject) error
	GetValues(i ...interface{}) []string
	GetProjectsTable() [][]string
	Model() interface{}
}
