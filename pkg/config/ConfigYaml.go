package config

import (
	"io/ioutil"
	"log"
	"time"

	I "github.com/gregod-com/grgd/interfaces"

	"gopkg.in/yaml.v2"
)

// CreateConfigYAML ...
func CreateConfigYAML(configpath string) I.IConfig {
	var obj = &ConfigYAML{path: configpath}
	obj.initFromFile()
	// return obj
	return nil
}

// ConfigYAML implements the IConfig based on yaml file...
type ConfigYAML struct {
	path         string
	PluginConfig []string
	PluginFolder []string
	LastUsed     time.Time
}

// InitFromFile ..
func (yamlObj *ConfigYAML) initFromFile() error {
	// userpath, err := os.UserHomeDir()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// if !strings.Contains(yamlObj.path, userpath) {
	// 	yamlObj.path = userpath + yamlObj.path
	// }

	err := yaml.Unmarshal(yamlObj.getSourceAsBytes(), yamlObj)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

func (yamlObj *ConfigYAML) getSourceAsBytes() []byte {
	conf, err := ioutil.ReadFile(yamlObj.path)
	if err != nil {
		yamlObj.Update()
		log.Println("A new config file has been created at " + yamlObj.path)
		log.Fatal("Run the 'init' command next to configure your stack.")
	}
	return conf
}

// Update ...
func (yamlObj *ConfigYAML) Update() error {
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

// GetSourceAsString ...
func (yamlObj *ConfigYAML) GetSourceAsString() string {
	return string(yamlObj.getSourceAsBytes())
}

// GetConfigPath ...
func (yamlObj *ConfigYAML) GetConfigPath() string {
	return yamlObj.path
}

// GetProjectDirs ...
func (yamlObj *ConfigYAML) GetProjectDirs() []string {
	// return yamlObj.ProjectDirectories
	return []string{"das"}
}

// GetLastUsed ...
func (yamlObj *ConfigYAML) GetLastUsed() time.Time {
	return yamlObj.LastUsed
}

// // GetWorkloadMetadata ...
// func (yamlObj *ConfigYAML) GetWorkloadMetadata() map[string]I.IWorkloadMetadata {
// 	var wlmeta = make(map[string]I.IWorkloadMetadata)
// 	for k := range yamlObj.WorkloadsMetadata {
// 		wlmeta[k] = yamlObj.WorkloadsMetadata[k]
// 	}
// 	return wlmeta
// }

// // GetWorkloads ...
// func (yamlObj *ConfigYAML) GetWorkloads() map[string]I.IWorkload {
// 	workloads := map[string]I.IWorkload{}
// 	// for k := range yamlObj.WorkloadsMetadata {
// 	// 	workloads[k] = dc.CreateWorkload(yamlObj.WorkloadsMetadata[k])
// 	// }
// 	return workloads
// }

// // GetRegistries ...
// func (yamlObj *ConfigYAML) GetRegistries() map[string]string {
// 	return yamlObj.Registries
// }

// // AddWorkloadShortcut ...
// func (yamlObj *ConfigYAML) AddWorkloadShortcut(shortcut string, workload string) error {
// 	// check if shortcut exists
// 	if val := yamlObj.Shortcuts[shortcut]; val == "" {
// 		// check if the workload name is valid
// 		for _, s := range yamlObj.WorkloadsMetadata {
// 			if s.GetName() == workload {
// 				yamlObj.Shortcuts[shortcut] = workload
// 				return nil
// 			}
// 		}
// 		return errors.New("WorkloadNotFound")
// 	}
// 	return errors.New("ShortcutExists")
// }

// // RemoveWorkloadShortcut ...
// func (yamlObj *ConfigYAML) RemoveWorkloadShortcut(shortcut string) error {
// 	// check if shortcut exists
// 	if val := yamlObj.Shortcuts[shortcut]; val == "" {
// 		return errors.New("ShortcutNotFound")
// 	}
// 	delete(yamlObj.Shortcuts, shortcut)
// 	return nil
// }

// // GetWorkloadShortcuts ...
// func (yamlObj *ConfigYAML) GetWorkloadShortcuts() map[string]string {
// 	return yamlObj.Shortcuts
// }

// // GetWorkloadByShortcut ...
// func (yamlObj *ConfigYAML) GetWorkloadByShortcut(shortcut string) string {
// 	return yamlObj.Shortcuts[shortcut]
// }
