package config

import (
	"fmt"

	"gopkg.in/yaml.v3"

	I "github.com/gregod-com/grgd/interfaces"
	"github.com/gregod-com/grgd/pkg/profile"
)

// ProvideConfig ...
func ProvideConfig(dal I.IDAL, ui I.IUIPlugin, logger I.ILogger, helper I.IHelper) I.IConfig {
	config := &ConfigDatabase{
		dal:    dal,
		ui:     ui,
		logger: logger,
		helper: helper,
	}
	// var prof I.IProfile
	var profiles []I.IProfile
	allProfiles, err := dal.ReadAll(profiles)
	if err != nil {
		config.logger.Warnf("Error loading profiles")
	}

	// remove element with key ""
	_, ok := allProfiles[""]
	if ok {
		delete(allProfiles, "")
	}

	config.profiles = make(map[string]I.IProfile)
	for name, prof := range allProfiles {
		p, ok := prof.(I.IProfile)
		if !ok {
			continue
		}
		config.profiles[name] = p
		config.logger.Debugf("Adding profile: %s", name)
	}

	config.profileProvider = profile.InitNewProfile
	config.activeProfile = helper.CheckUserProfile()
	config.logger.Tracef("provide %T", config)
	return config
}

// ConfigDatabase ...
type ConfigDatabase struct {
	dal             I.IDAL
	logger          I.ILogger
	ui              I.IUIPlugin
	helper          I.IHelper
	profileProvider func(name string, ui I.IUIPlugin, log I.ILogger, helper I.IHelper, i ...interface{}) I.IProfile
	projectProvider func(i ...interface{}) I.IProject
	profiles        map[string]I.IProfile
	activeProfile   string
}

// Save ...
func (coDB *ConfigDatabase) Save(i ...interface{}) error {
	coDB.logger.Trace("")
	for k := range coDB.profiles {
		err := coDB.dal.Update(coDB.profiles[k])
		if err != nil {
			coDB.logger.Fatal(err)
		}
		// for _, proj := range coDB.profiles[k].GetProjects() {
		// 	err := coDB.dal.Update(proj)
		// 	if err != nil {
		// 		coDB.logger.Fatal(err)
		// 	}
		// }
	}
	coDB.logger.Trace("")
	return nil
}

// GetAllProfiles ...
func (coDB *ConfigDatabase) GetAllProfiles() (map[string]I.IProfile, error) {
	coDB.logger.Trace("")

	return coDB.profiles, nil
}

// InitNewProfile ...
func (coDB *ConfigDatabase) InitNewProfile(name string) error {
	coDB.logger.Trace("")
	coDB.logger.Debug("Profile %s not set! Starting init process...", name)
	p := coDB.profileProvider(name, coDB.ui, coDB.logger, coDB.helper)
	// p := profile.InitNewProfile(name, coDB.logger, coDB.ui, coDB.helper)
	return coDB.AddProfile(p)
}

// SetActiveProfile ...
func (coDB *ConfigDatabase) SetActiveProfile(name string) error {
	coDB.logger.Trace("")

	coDB.activeProfile = name
	return nil
}

// GetActiveProfile ...
func (coDB *ConfigDatabase) GetActiveProfile() I.IProfile {
	coDB.logger.Trace("")

	if _, ok := coDB.profiles[coDB.activeProfile]; !ok {
		if coDB.InitNewProfile(coDB.activeProfile) != nil {
			coDB.logger.Fatal("Could not create new profile")
		}
	}
	return coDB.profiles[coDB.activeProfile]
}

// GetActiveProfileByName ...
func (coDB *ConfigDatabase) GetActiveProfileByName(profilename string) (I.IProfile, error) {
	coDB.logger.Trace("")

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
	coDB.logger.Tracef("")
	coDB.dal.Delete(p)
	return nil
}

// RemoveProject ...
func (coDB *ConfigDatabase) Remove(i interface{}) error {
	coDB.logger.Tracef("")
	coDB.dal.Delete(i)
	return nil
}

// GetConfigPath ...
func (coDB *ConfigDatabase) GetConfigPath() (string, error) {
	coDB.logger.Tracef("")
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
