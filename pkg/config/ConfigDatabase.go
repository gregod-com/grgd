package config

import (
	"fmt"

	"gopkg.in/yaml.v2"

	I "github.com/gregod-com/grgd/interfaces"
	"github.com/gregod-com/grgd/pkg/profile"
)

// ProvideConfig ...
func ProvideConfig(dal I.IDAL, ui I.IUIPlugin, logger I.ILogger, fsm I.IFileSystemManipulator, prof I.IProfile) I.IConfig {
	config := &ConfigDatabase{
		dal:    dal,
		ui:     ui,
		logger: logger,
		fsm:    fsm,
	}
	mp, err := dal.ReadAll(prof)
	if err != nil {
		config.logger.Warnf("Error loading profiles")
	}
	config.profiles = make(map[string]I.IProfile)
	for k, v := range mp {
		config.profiles[k] = v.(I.IProfile)
	}
	config.logger.Tracef("provide %T", config)
	return config
}

// ConfigDatabase ...
type ConfigDatabase struct {
	dal           I.IDAL
	logger        I.ILogger
	ui            I.IUIPlugin
	fsm           I.IFileSystemManipulator
	profiles      map[string]I.IProfile
	activeProfile string `yaml:"activeProfile"`
}

// Save ...
func (coDB *ConfigDatabase) Save(i ...interface{}) error {
	for k := range coDB.profiles {
		err := coDB.dal.Update(coDB.profiles[k])
		if err != nil {
			coDB.logger.Fatal(err)
		}
	}

	coDB.logger.Trace("")
	return nil
}

// GetAllProfiles ...
func (coDB *ConfigDatabase) GetAllProfiles() (map[string]I.IProfile, error) {
	return coDB.profiles, nil
}

// InitNewProfile ...
func (coDB *ConfigDatabase) InitNewProfile(name string) error {
	coDB.logger.Tracef("Profile %s not set! Starting init process...", name)
	p := profile.InitNewProfile(name, coDB.logger, coDB.ui, coDB.fsm)
	return coDB.AddProfile(p)
}

// SetActiveProfile ...
func (coDB *ConfigDatabase) SetActiveProfile(name string) error {
	coDB.activeProfile = name
	return nil
}

// GetActiveProfile ...
func (coDB *ConfigDatabase) GetActiveProfile() I.IProfile {
	if _, ok := coDB.profiles[coDB.activeProfile]; !ok {
		if coDB.InitNewProfile(coDB.activeProfile) != nil {
			coDB.logger.Fatal("Could not create new profile")
		}
	}
	return coDB.profiles[coDB.activeProfile]
}

// GetActiveProfileByName ...
func (coDB *ConfigDatabase) GetActiveProfileByName(profilename string) (I.IProfile, error) {
	current, ok := coDB.profiles[profilename]
	if !ok {
		coDB.logger.Tracef("Profile %s not set! Starting init process...", profilename)
		// p := profile.InitNewProfile(profilename, coDB.logger, coDB.ui, coDB.fsm)
		// coDB.AddProfile(p)
		return nil, nil
	}
	return current, nil
}

// AddProfile ...
func (coDB *ConfigDatabase) AddProfile(p I.IProfile) error {
	if _, ok := coDB.profiles[p.GetName()]; ok {
		return fmt.Errorf("Profile %v already exists", p.GetName())
	}
	coDB.profiles[p.GetName()] = p
	return nil
}

// RemoveProfile ...
func (coDB *ConfigDatabase) RemoveProfile(p I.IProfile) error {
	return nil
}

// GetAllProjects ...
func (coDB *ConfigDatabase) GetAllProjects() (map[string][]I.IProject, error) {
	return nil, nil

}

// GetProjects ...
func (coDB *ConfigDatabase) GetProjects() (map[string]I.IProject, error) {
	return coDB.GetActiveProfile().GetProjects(), nil
}

// GetProjectByName ...
func (coDB *ConfigDatabase) GetProjectByName(projectName string) (I.IProject, error) {
	p := coDB.GetActiveProfile().GetProjects()[projectName]
	return p, nil
}

// AddProject ...
func (coDB *ConfigDatabase) AddProject(p string, i ...interface{}) error {
	return coDB.GetActiveProfile().AddProject(p)
}

// RemoveProject ...
func (coDB *ConfigDatabase) RemoveProject(p I.IProject, i ...interface{}) error {
	// db := coDB.dal.Delete(&persistence.GRGDProject{}, &persistence.GRGDProject{Name: p.GetName()})
	// delete(coDB.GetActiveProfile().GetProjects(), p.GetName())
	return nil
}

// SwitchCurrentProject ...
func (coDB *ConfigDatabase) SwitchCurrentProject(i ...interface{}) (I.IProject, error) {
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
	profile := coDB.GetActiveProfile()
	b, err := yaml.Marshal(coDB.profiles)
	if err != nil {
		coDB.logger.Fatal(err)
	}
	coDB.ui.Printf("---- YAML for profile: [%v] ----- \n%v \n", profile.GetName(), string(b))
	return coDB
}

func (coDB *ConfigDatabase) String() string {
	// profile := coDB.GetActiveProfile()
	b, err := yaml.Marshal(&coDB)
	if err != nil {
		coDB.logger.Fatal(err)
	}
	return string(b)
}
