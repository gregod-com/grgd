package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"

	I "github.com/gregod-com/interfaces"

	yaml "gopkg.in/yaml.v2"
)

// ConfigObjectYaml implements the IConfigObject based on yaml file...
type ConfigObjectYaml struct {
	path              string                       `yaml:"a,omitempty"`
	Debug             bool                         `yaml:"debug"`
	ProjectDirectory  string                       `yaml:"projectDir"`
	Registries        map[string]string            `yaml:"registries"`
	WorkloadsMetadata map[string]*WorkloadMetadata `yaml:"services"`
	ServicesToIgnore  []string                     `yaml:"servicestoignore"`
	Shortcuts         map[string]string            `yaml:"shortcuts"`
	CommandsUsed      map[string]bool              `yaml:"commands"`
	LastUsed          time.Time                    `yaml:"lastused"`
}

func (yamlObj *ConfigObjectYaml) getSourceAsBytes() []byte {
	iamconf, err := ioutil.ReadFile(yamlObj.path)
	if err != nil {
		yamlObj.Update()
		log.Println("A new config file has been created at " + yamlObj.path)
		log.Fatal("Run the 'init' command next to configure your stack.")
	}
	return iamconf
}

// InitFromFile ..
func (yamlObj *ConfigObjectYaml) initFromFile() error {
	userpath, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	if !strings.Contains(yamlObj.path, userpath) {
		yamlObj.path = userpath + yamlObj.path
	}

	err = yaml.Unmarshal(yamlObj.getSourceAsBytes(), yamlObj)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

// CreateConfigObjectYaml ...
func CreateConfigObjectYaml(configpath string) I.IConfigObject {
	var obj = &ConfigObjectYaml{path: configpath}
	obj.initFromFile()
	return obj
}

// Update ...
func (yamlObj *ConfigObjectYaml) Update() error {
	yamlObj.LastUsed = time.Now()
	newyaml, err := yaml.Marshal(yamlObj)
	if err != nil {
		log.Fatal(err)
	}

	err = ioutil.WriteFile(yamlObj.path, newyaml, 0644)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

// PrintConfig ...
func (yamlObj *ConfigObjectYaml) PrintConfig() error {
	fmt.Println(yamlObj.GetSourceAsString())
	return nil
}

// GetSourceAsString ...
func (yamlObj *ConfigObjectYaml) GetSourceAsString() string {
	return string(yamlObj.getSourceAsBytes())
}

// GetConfigPath ...
func (yamlObj *ConfigObjectYaml) GetConfigPath() string {
	return yamlObj.path
}

// IsDebug ...
func (yamlObj *ConfigObjectYaml) IsDebug() bool {
	return yamlObj.Debug

}

// GetProjectDir ...
func (yamlObj *ConfigObjectYaml) GetProjectDir() string {
	return yamlObj.ProjectDirectory
}

// GetWorkloads ...
func (yamlObj *ConfigObjectYaml) GetWorkloads() map[string]I.IWorkload {
	workloads := map[string]I.IWorkload{}
	// for k := range yamlObj.WorkloadsMetadata {
	// 	workloads[k] = dc.CreateWorkload(yamlObj.WorkloadsMetadata[k])
	// }
	return workloads
}

// GetWorkloadMetadata ...
func (yamlObj *ConfigObjectYaml) GetWorkloadMetadata() map[string]I.IWorkloadMetadata {
	var wlmeta = make(map[string]I.IWorkloadMetadata)
	for k := range yamlObj.WorkloadsMetadata {
		wlmeta[k] = yamlObj.WorkloadsMetadata[k]
	}
	return wlmeta
}

// GetRegistries ...
func (yamlObj *ConfigObjectYaml) GetRegistries() map[string]string {
	return yamlObj.Registries
}

// AddWorkloadShortcut ...
func (yamlObj *ConfigObjectYaml) AddWorkloadShortcut(shortcut string, workload string) error {
	// check if shortcut exists
	if val := yamlObj.Shortcuts[shortcut]; val == "" {
		// check if the workload name is valid
		for _, s := range yamlObj.WorkloadsMetadata {
			if s.GetName() == workload {
				yamlObj.Shortcuts[shortcut] = workload
				return nil
			}
		}
		return errors.New("WorkloadNotFound")
	}
	return errors.New("ShortcutExists")
}

// RemoveWorkloadShortcut ...
func (yamlObj *ConfigObjectYaml) RemoveWorkloadShortcut(shortcut string) error {
	// check if shortcut exists
	if val := yamlObj.Shortcuts[shortcut]; val == "" {
		return errors.New("ShortcutNotFound")
	}
	delete(yamlObj.Shortcuts, shortcut)
	return nil
}

// GetWorkloadShortcuts ...
func (yamlObj *ConfigObjectYaml) GetWorkloadShortcuts() map[string]string {
	return yamlObj.Shortcuts
}

// GetWorkloadByShortcut ...
func (yamlObj *ConfigObjectYaml) GetWorkloadByShortcut(shortcut string) string {
	return yamlObj.Shortcuts[shortcut]
}

// WasCommandUsed ...
func (yamlObj *ConfigObjectYaml) WasCommandUsed(command string) bool {
	if yamlObj.CommandsUsed[command] {
		return true
	}
	return false
}

// LearnedCommands ...
func (yamlObj *ConfigObjectYaml) LearnedCommands() int {
	return len(yamlObj.CommandsUsed)
}

// MarkCommandLerned ...
func (yamlObj *ConfigObjectYaml) MarkCommandLerned(command string) error {
	yamlObj.CommandsUsed[command] = true
	yamlObj.Update()
	return nil
}

// GetLastUsed ...
func (yamlObj *ConfigObjectYaml) GetLastUsed() time.Time {
	return yamlObj.LastUsed
}
