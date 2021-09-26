package interfaces

// IConfig interface ...
type IConfig interface {
	Save(i ...interface{}) error

	SetActiveProfile(name string) error
	GetActiveProfile() IProfile

	GetAllProfiles() (map[string]IProfile, error)
	GetActiveProfileByName(profilename string) (IProfile, error)

	AddProfile(p IProfile) error
	RemoveProfile(p IProfile) error

	Remove(i interface{}) error

	GetConfigPath() (string, error)
	SetConfigPath(path string) error

	DumpConfig(i ...interface{}) interface{}
}
