package config

import (
	"github.com/gregod-com/grgdplugincontracts"

	"grgd/interfaces"
)

// ProvideConfigObject ...
func ProvideConfigObject(dal interfaces.IDAL, logger interfaces.ILogger) interfaces.IConfigObject {
	config := new(ConfigObjectDatabase)
	config.dal = dal
	config.logger = logger
	return config
}

// ConfigObjectDatabase ...
type ConfigObjectDatabase struct {
	dal            interfaces.IDAL
	logger         interfaces.ILogger
	ui             grgdplugincontracts.IUIPlugin
	profiles       map[string]interfaces.IProfile
	currentProfile string
}

// Save ...
func (coDB *ConfigObjectDatabase) Save(i ...interface{}) error {
	coDB.dal.Update(i)
	coDB.logger.Trace("Saving...")
	return nil
}

// GetAllProfiles ...
func (coDB *ConfigObjectDatabase) GetAllProfiles() (map[string]interfaces.IProfile, error) {
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

// GetProfile ...
func (coDB *ConfigObjectDatabase) GetProfile() interfaces.IProfile {
	current, ok := coDB.profiles[coDB.currentProfile]
	if !ok {
		for k := range coDB.profiles {
			coDB.currentProfile = coDB.profiles[k].GetName()
			return coDB.profiles[k]
		}
		// coDB.AddProfile(wrappers.InitNewProfile(&persistence.Profile{Name: coDB.currentProfile}, coDB.logger, &view.FallbackUI{}, &helper.FSManipulator{}))
		coDB.logger.Fatal("Current Profile not set! Init Here")
	}
	return current
}

// AddProfile ...
func (coDB *ConfigObjectDatabase) AddProfile(p interfaces.IProfile) error {
	coDB.logger.Fatal("here in add profile")
	return nil

}

// RemoveProfile ...
func (coDB *ConfigObjectDatabase) RemoveProfile(p interfaces.IProfile) error {
	return nil
}

// GetProfileByName ...
func (coDB *ConfigObjectDatabase) GetProfileByName(profilename string) error {
	return nil

}

// GetAllProjects ...
func (coDB *ConfigObjectDatabase) GetAllProjects() (map[string][]interfaces.IProject, error) {
	return nil, nil

}

// GetProjects ...
func (coDB *ConfigObjectDatabase) GetProjects() (map[string]interfaces.IProject, error) {
	return coDB.GetProfile().GetProjects(), nil
}

// GetProjectByName ...
func (coDB *ConfigObjectDatabase) GetProjectByName(projectName string) (interfaces.IProject, error) {
	p := coDB.GetProfile().GetProjects()[projectName]
	return p, nil
}

// AddProject ...
func (coDB *ConfigObjectDatabase) AddProject(p string, i ...interface{}) error {
	return coDB.GetProfile().AddProject(p)
}

// RemoveProject ...
func (coDB *ConfigObjectDatabase) RemoveProject(p interfaces.IProject, i ...interface{}) error {
	// db := coDB.dal.Delete(&persistence.GRGDProject{}, &persistence.GRGDProject{Name: p.GetName()})
	// delete(coDB.GetProfile().GetProjects(), p.GetName())
	return nil
}

// SwitchCurrentProject ...
func (coDB *ConfigObjectDatabase) SwitchCurrentProject(i ...interface{}) (interfaces.IProject, error) {

	return nil, nil
}

// GetConfigPath ...
func (coDB *ConfigObjectDatabase) GetConfigPath() (string, error) {
	return "", nil
}

// SetConfigPath ...
func (coDB *ConfigObjectDatabase) SetConfigPath(path string) error {
	return nil
}

// DumpConfig ...
func (coDB *ConfigObjectDatabase) DumpConfig(i ...interface{}) interface{} {
	return nil
}
