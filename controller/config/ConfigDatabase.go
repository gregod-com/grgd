package config

import (
	"os"

	"gopkg.in/yaml.v2"

	"github.com/gregod-com/grgd/controller/profile"
	"github.com/gregod-com/grgd/interfaces"
)

// ProvideConfig ...
func ProvideConfig(dal interfaces.IDAL,
	ui interfaces.IUIPlugin,
	logger interfaces.ILogger,
	fsm interfaces.IFileSystemManipulator) interfaces.IConfig {
	config := new(ConfigDatabase)
	config.dal = dal
	config.ui = ui
	config.logger = logger
	config.fsm = fsm
	p, err := dal.GetProfile()
	if err == nil && p != nil {
		config.currentProfile = p.GetName()
		config.profiles = make(map[string]interfaces.IProfile)
		config.profiles[p.GetName()] = profile.CreateProfile(p)
	}
	return config
}

// ConfigDatabase ...
type ConfigDatabase struct {
	dal            interfaces.IDAL
	logger         interfaces.ILogger
	ui             interfaces.IUIPlugin
	fsm            interfaces.IFileSystemManipulator
	profiles       map[string]interfaces.IProfile
	currentProfile string
}

// Save ...
func (coDB *ConfigDatabase) Save(i ...interface{}) error {
	for k := range coDB.profiles {
		err := coDB.dal.Update(coDB.profiles[k].Model())
		if err != nil {
			coDB.logger.Fatal(err)
		}
	}

	coDB.logger.Trace("Saving...")
	return nil
}

// GetAllProfiles ...
func (coDB *ConfigDatabase) GetAllProfiles() (map[string]interfaces.IProfile, error) {
	// cB.dal.
	// var profs []*persistence.Profile
	// persistence.GetAll(&profs)
	// for _, v := range profs {
	// 	persistence.GetAll(&v)
	// 	coDB.profiles[v.Name] = wrappers.CreateProfile(v)
	// }
	// return coDB.profiles, nil
	return nil, nil
}

// InitNewProfile ...
func (coDB *ConfigDatabase) InitNewProfile(name string) error {

	coDB.logger.Tracef("Profile %s not set! Starting init process...", name)
	p := profile.InitNewProfile(name, coDB.logger, coDB.ui, coDB.fsm)
	return coDB.AddProfile(p)
}

// GetProfile ...
func (coDB *ConfigDatabase) GetProfile() interfaces.IProfile {
	if _, ok := coDB.profiles[coDB.currentProfile]; !ok {
		coDB.currentProfile = os.Getenv("USER")

		if _, ok := coDB.profiles[coDB.currentProfile]; !ok {
			if coDB.InitNewProfile(coDB.currentProfile) != nil {
				coDB.logger.Fatal("Could not create new profile")
			}
		}
	}
	return coDB.profiles[coDB.currentProfile]
}

// GetProfileByName ...
func (coDB *ConfigDatabase) GetProfileByName(profilename string) (interfaces.IProfile, error) {
	current, ok := coDB.profiles[profilename]
	if !ok {
		coDB.logger.Tracef("Profile %s not set! Starting init process...", profilename)
		p := profile.InitNewProfile(profilename, coDB.logger, coDB.ui, coDB.fsm)
		coDB.AddProfile(p)
		return p, nil
	}
	return current, nil
}

// AddProfile ...
func (coDB *ConfigDatabase) AddProfile(p interfaces.IProfile) error {

	if coDB.profiles == nil {
		coDB.profiles = make(map[string]interfaces.IProfile)
		coDB.profiles[p.GetName()] = p
		coDB.currentProfile = p.GetName()
	}
	return nil
}

// RemoveProfile ...
func (coDB *ConfigDatabase) RemoveProfile(p interfaces.IProfile) error {
	return nil
}

// GetAllProjects ...
func (coDB *ConfigDatabase) GetAllProjects() (map[string][]interfaces.IProject, error) {
	return nil, nil

}

// GetProjects ...
func (coDB *ConfigDatabase) GetProjects() (map[string]interfaces.IProject, error) {
	return coDB.GetProfile().GetProjects(), nil
}

// GetProjectByName ...
func (coDB *ConfigDatabase) GetProjectByName(projectName string) (interfaces.IProject, error) {
	p := coDB.GetProfile().GetProjects()[projectName]
	return p, nil
}

// AddProject ...
func (coDB *ConfigDatabase) AddProject(p string, i ...interface{}) error {
	return coDB.GetProfile().AddProject(p)
}

// RemoveProject ...
func (coDB *ConfigDatabase) RemoveProject(p interfaces.IProject, i ...interface{}) error {
	// db := coDB.dal.Delete(&persistence.GRGDProject{}, &persistence.GRGDProject{Name: p.GetName()})
	// delete(coDB.GetProfile().GetProjects(), p.GetName())
	return nil
}

// SwitchCurrentProject ...
func (coDB *ConfigDatabase) SwitchCurrentProject(i ...interface{}) (interfaces.IProject, error) {
	return nil, nil
}

// GetConfigPath ...
func (coDB *ConfigDatabase) GetConfigPath() (string, error) {
	return "", nil
}

// SetConfigPath ...
func (coDB *ConfigDatabase) SetConfigPath(path string) error {
	return nil
}

// DumpConfig ...
func (coDB *ConfigDatabase) DumpConfig(i ...interface{}) interface{} {
	profile := coDB.GetProfile()
	b, err := yaml.Marshal(profile.Model())
	if err != nil {
		coDB.logger.Fatal(err)
	}
	coDB.ui.Printf("---- YAML for profile: [%v] ----- \n%v \n", profile.GetName(), string(b))
	return profile
}
