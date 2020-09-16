package config

import (
	"io/ioutil"
	"time"

	log "github.com/sirupsen/logrus"

	I "grgd/interfaces"

	"gopkg.in/yaml.v2"
)

// CreateConfigObjectYAML ...
func CreateConfigObjectYAML(configpath string) I.IConfigObject {
	var obj = &ConfigObjectYAML{path: configpath}
	obj.initFromFile()
	// return obj
	return nil
}

// ConfigObjectYAML implements the IConfigObject based on yaml file...
type ConfigObjectYAML struct {
	path         string
	PluginConfig []string
	PluginFolder []string
	LastUsed     time.Time
}

//                                                       _                _
//  _   _   _ __     ___  __  __  _ __     ___    _ __  | |_    ___    __| |
// | | | | | '_ \   / _ \ \ \/ / | '_ \   / _ \  | '__| | __|  / _ \  / _` |
// | |_| | | | | | |  __/  >  <  | |_) | | (_) | | |    | |_  |  __/ | (_| |
//  \__,_| |_| |_|  \___| /_/\_\ | .__/   \___/  |_|     \__|  \___|  \__,_|
//                               |_|

// InitFromFile ..
func (yamlObj *ConfigObjectYAML) initFromFile() error {
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

func (yamlObj *ConfigObjectYAML) getSourceAsBytes() []byte {
	iamconf, err := ioutil.ReadFile(yamlObj.path)
	if err != nil {
		yamlObj.Update()
		log.Println("A new config file has been created at " + yamlObj.path)
		log.Fatal("Run the 'init' command next to configure your stack.")
	}
	return iamconf
}

//                                       _                _
//   ___  __  __  _ __     ___    _ __  | |_    ___    __| |
//  / _ \ \ \/ / | '_ \   / _ \  | '__| | __|  / _ \  / _` |
// |  __/  >  <  | |_) | | (_) | | |    | |_  |  __/ | (_| |
//  \___| /_/\_\ | .__/   \___/  |_|     \__|  \___|  \__,_|
//               |_|

// Update ...
func (yamlObj *ConfigObjectYAML) Update() error {
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
func (yamlObj *ConfigObjectYAML) GetSourceAsString() string {
	return string(yamlObj.getSourceAsBytes())
}

// GetConfigPath ...
func (yamlObj *ConfigObjectYAML) GetConfigPath() string {
	return yamlObj.path
}

// GetProjectDirs ...
func (yamlObj *ConfigObjectYAML) GetProjectDirs() []string {
	// return yamlObj.ProjectDirectories
	return []string{"das"}
}

// GetLastUsed ...
func (yamlObj *ConfigObjectYAML) GetLastUsed() time.Time {
	return yamlObj.LastUsed
}

// // GetWorkloadMetadata ...
// func (yamlObj *ConfigObjectYAML) GetWorkloadMetadata() map[string]I.IWorkloadMetadata {
// 	var wlmeta = make(map[string]I.IWorkloadMetadata)
// 	for k := range yamlObj.WorkloadsMetadata {
// 		wlmeta[k] = yamlObj.WorkloadsMetadata[k]
// 	}
// 	return wlmeta
// }

// // GetWorkloads ...
// func (yamlObj *ConfigObjectYAML) GetWorkloads() map[string]I.IWorkload {
// 	workloads := map[string]I.IWorkload{}
// 	// for k := range yamlObj.WorkloadsMetadata {
// 	// 	workloads[k] = dc.CreateWorkload(yamlObj.WorkloadsMetadata[k])
// 	// }
// 	return workloads
// }

// // GetRegistries ...
// func (yamlObj *ConfigObjectYAML) GetRegistries() map[string]string {
// 	return yamlObj.Registries
// }

// // AddWorkloadShortcut ...
// func (yamlObj *ConfigObjectYAML) AddWorkloadShortcut(shortcut string, workload string) error {
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
// func (yamlObj *ConfigObjectYAML) RemoveWorkloadShortcut(shortcut string) error {
// 	// check if shortcut exists
// 	if val := yamlObj.Shortcuts[shortcut]; val == "" {
// 		return errors.New("ShortcutNotFound")
// 	}
// 	delete(yamlObj.Shortcuts, shortcut)
// 	return nil
// }

// // GetWorkloadShortcuts ...
// func (yamlObj *ConfigObjectYAML) GetWorkloadShortcuts() map[string]string {
// 	return yamlObj.Shortcuts
// }

// // GetWorkloadByShortcut ...
// func (yamlObj *ConfigObjectYAML) GetWorkloadByShortcut(shortcut string) string {
// 	return yamlObj.Shortcuts[shortcut]
// }
