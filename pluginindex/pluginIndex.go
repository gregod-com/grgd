package pluginindex

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"

	plugContracts "github.com/gregod-com/grgdplugincontracts"
	"golang.org/x/mod/semver"
	yaml "gopkg.in/yaml.v2"
)

// PluginIndex ...
type PluginIndex struct {
	path                   string                                   `yaml:"-"`
	PluginMetadataList     map[string]plugContracts.IPluginMetadata `yaml:"-"`
	PluginMetadataListYAML map[string]PluginMetadataImpl            `yaml:"plugins"`
	Lastchecked            time.Time                                `yaml:"lastchecked"`
}

// CreatePluginIndex ...
func CreatePluginIndex(path string) plugContracts.IPluginIndex {
	var obj = &PluginIndex{path: path}
	obj.initFromFile()
	return obj
}

func (yamlObj *PluginIndex) initFromFile() error {
	err := yaml.Unmarshal(yamlObj.getSourceAsBytes(), yamlObj)
	if err != nil {
		log.Fatal(err)
	}
	yamlObj.PluginMetadataList = make(map[string]plugContracts.IPluginMetadata, len(yamlObj.PluginMetadataListYAML))
	for k := range yamlObj.PluginMetadataListYAML {
		yamlObj.PluginMetadataList[k] = convertInterfaceToImpl(yamlObj.PluginMetadataListYAML[k])
	}
	return nil
}

// Update ...
func (yamlObj *PluginIndex) Update() error {
	yamlObj.Lastchecked = time.Now()

	for k := range yamlObj.PluginMetadataList {
		yamlObj.PluginMetadataListYAML[k] = convertImplToInterface(yamlObj.PluginMetadataList[k])
	}

	newyaml, err := yaml.Marshal(yamlObj)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(yamlObj.path, newyaml, 0644)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

// GetSourceAsString ...
func (yamlObj *PluginIndex) GetSourceAsString() string {
	return string(yamlObj.getSourceAsBytes())
}

func (yamlObj *PluginIndex) getSourceAsBytes() []byte {
	config, err := ioutil.ReadFile(yamlObj.path)
	if err != nil {
		// yamlObj.Update()
		log.Println("A new config file has been created at " + yamlObj.path)
		log.Fatal("Run the 'init' command next to configure your stack.")
	}
	return config
}

// GetConfigPath ...
func (yamlObj *PluginIndex) GetConfigPath() string {
	return yamlObj.path
}

// GetLastChecked ...
func (yamlObj *PluginIndex) GetLastChecked() time.Time {
	return yamlObj.Lastchecked
}

// GetPluginList ...
func (yamlObj *PluginIndex) GetPluginList() map[string]plugContracts.IPluginMetadata {
	return yamlObj.PluginMetadataList
}

// GetPluginListActive ...
func (yamlObj *PluginIndex) GetPluginListActive() map[string]plugContracts.IPluginMetadata {
	keys := map[string]plugContracts.IPluginMetadata{}
	for k, v := range yamlObj.PluginMetadataList {
		if v.GetActive() {
			keys[k] = yamlObj.PluginMetadataList[k]
		}
	}
	return keys
}

// GetPluginListInactive ...
func (yamlObj *PluginIndex) GetPluginListInactive() map[string]plugContracts.IPluginMetadata {
	keys := map[string]plugContracts.IPluginMetadata{}
	for k, v := range yamlObj.PluginMetadataList {
		if !v.GetActive() {
			keys[k] = yamlObj.PluginMetadataList[k]
		}
	}
	return keys
}

// AddPlugin ...
func (yamlObj *PluginIndex) AddPlugin(newplugImpl plugContracts.IPluginMetadata) string {

	identifier := newplugImpl.GetCategory() + "-" + newplugImpl.GetName()

	if plug, ok := yamlObj.PluginMetadataListYAML[identifier]; ok {
		switch semver.Compare(plug.GetVersion(), newplugImpl.GetVersion()) {
		case -1:
			fmt.Printf("Update plugin %v from %v to %v? [y/n] ", newplugImpl.GetName(), plug.GetVersion(), newplugImpl.GetVersion())
			if ynQuestion() {
				newplugImpl.SetActive(true)
				yamlObj.PluginMetadataListYAML[identifier] = convertImplToInterface(newplugImpl)
				yamlObj.PluginMetadataList[identifier] = newplugImpl
				yamlObj.Update()
				fmt.Println(yamlObj.PluginMetadataListYAML)
			}
			return newplugImpl.GetCategory()
		case 1:
			fmt.Printf("An older version for plugin %v (%v) was found. You should remove the plugin file? [y/n]", newplugImpl.GetName(), newplugImpl.GetVersion())
			if ynQuestion() {
				return "remove"
			}
		default:
			if plug.GetActive() {
				// log.Println("Loading existing active plugin" + plug.GetName())
				return plug.GetCategory()
			}
			return "disabled"
		}
	}
	yamlObj.PluginMetadataListYAML[identifier] = convertImplToInterface(newplugImpl)
	fmt.Println(yamlObj.PluginMetadataListYAML[identifier])
	return newplugImpl.GetCategory()
}

func ynQuestion() bool {
	reader := bufio.NewReader(os.Stdin)
	answer, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	answer = strings.Replace(answer, "\n", "", -1)
	if strings.Contains(answer, "y") {
		return true
	}
	return false
}

func convertImplToInterface(pmeta plugContracts.IPluginMetadata) PluginMetadataImpl {
	return PluginMetadataImpl{
		Name:     pmeta.GetName(),
		Version:  pmeta.GetVersion(),
		URL:      pmeta.GetURL(),
		Category: pmeta.GetCategory(),
		Active:   pmeta.GetActive(),
		Path:     pmeta.GetPath(),
	}
}

func convertInterfaceToImpl(pmeta PluginMetadataImpl) plugContracts.IPluginMetadata {
	return &PluginMetadataImpl{
		Name:     pmeta.GetName(),
		Version:  pmeta.GetVersion(),
		URL:      pmeta.GetURL(),
		Category: pmeta.GetCategory(),
		Active:   pmeta.GetActive(),
		Path:     pmeta.GetPath(),
	}
}
