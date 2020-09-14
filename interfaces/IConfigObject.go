package interfaces

// IConfigObject interface ...
type IConfigObject interface {
	Save(i ...interface{}) error
	GetAllProfiles() (map[string]IProfile, error)
	GetProfile() IProfile
	GetProfileByName(profilename string) error
	AddProfile(p IProfile) error
	RemoveProfile(p IProfile) error

	GetAllProjects() (map[string][]IProject, error)
	GetProjects() (map[string]IProject, error)
	GetProjectByName(projectName string) (IProject, error)
	AddProject(p string, i ...interface{}) error
	RemoveProject(p IProject, i ...interface{}) error

	SwitchCurrentProject(i ...interface{}) (IProject, error)

	GetConfigPath() (string, error)
	SetConfigPath(path string) error

	DumpConfig(i ...interface{}) interface{}
}

// GetLastUsed() time.Time

// WasCommandUsed(string) bool
// LearnedCommands() int
// MarkCommandLerned(string) error

// those should maybe be moved to a interface that focuses on workloads
// GetWorkloadMetadata() map[string]IWorkloadMetadata
// GetWorkloads() map[string]IWorkload
// GetRegistries() map[string]string
// AddWorkloadShortcut(string, string) error
// RemoveWorkloadShortcut(string) error
// GetWorkloadShortcuts() map[string]string
// GetWorkloadByShortcut(string) string