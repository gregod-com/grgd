package pluginindex

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"

	"github.com/gregod-com/implementations"

	I "github.com/gregod-com/interfaces"
	"golang.org/x/mod/semver"
	yaml "gopkg.in/yaml.v2"
)

// PluginIndex ...
type PluginIndex struct {
	path                   string                                        `yaml:"-"`
	PluginMetadataList     map[string]I.IPluginMetadata                  `yaml:"-"`
	PluginMetadataListYAML map[string]implementations.PluginMetadataImpl `yaml:"plugins"`
	Lastchecked            time.Time                                     `yaml:"lastchecked"`
}

// CreatePluginIndex ...
func CreatePluginIndex(path string) I.IPluginIndex {
	var obj = &PluginIndex{path: path}
	obj.initFromFile()
	return obj
}

func (yamlObj *PluginIndex) initFromFile() error {
	err := yaml.Unmarshal(yamlObj.getSourceAsBytes(), yamlObj)
	if err != nil {
		log.Fatal(err)
	}
	yamlObj.PluginMetadataList = make(map[string]I.IPluginMetadata, len(yamlObj.PluginMetadataListYAML))
	for k := range yamlObj.PluginMetadataListYAML {
		yamlObj.PluginMetadataList[k] = ConvertInterfaceToImpl(yamlObj.PluginMetadataListYAML[k])
	}
	return nil
}

// Update ...
func (yamlObj *PluginIndex) Update() error {
	yamlObj.Lastchecked = time.Now()

	for k := range yamlObj.PluginMetadataList {
		yamlObj.PluginMetadataListYAML[k] = PluginMetaConverter(yamlObj.PluginMetadataList[k])
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

// PrintConfig ...
func (yamlObj *PluginIndex) PrintConfig() error {
	fmt.Println(yamlObj.GetSourceAsString())
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
func (yamlObj *PluginIndex) GetPluginList() map[string]I.IPluginMetadata {
	return yamlObj.PluginMetadataList
}

// AddPlugin ...
func (yamlObj *PluginIndex) AddPlugin(newplugImpl I.IPluginMetadata) string {

	identifier := newplugImpl.GetCategory() + "-" + newplugImpl.GetName()

	if plug, ok := yamlObj.PluginMetadataListYAML[identifier]; ok {
		switch semver.Compare(plug.GetVersion(), newplugImpl.GetVersion()) {
		case -1:
			fmt.Printf("Update plugin %v from %v to %v? [y/n] ", newplugImpl.GetName(), plug.GetVersion(), newplugImpl.GetVersion())
			if ynQuestion() {
				newplugImpl.SetActive(true)
				yamlObj.PluginMetadataListYAML[identifier] = PluginMetaConverter(newplugImpl)
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
	yamlObj.PluginMetadataListYAML[identifier] = PluginMetaConverter(newplugImpl)
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

// PluginMetaConverter ...
func PluginMetaConverter(pmeta I.IPluginMetadata) implementations.PluginMetadataImpl {
	return implementations.PluginMetadataImpl{
		Name:     pmeta.GetName(),
		Version:  pmeta.GetVersion(),
		Size:     pmeta.GetSize(),
		URL:      pmeta.GetURL(),
		Category: pmeta.GetCategory(),
		Active:   pmeta.GetActive(),
	}
}

// PluginMetaConverter ...
func ConvertInterfaceToImpl(pmeta implementations.PluginMetadataImpl) I.IPluginMetadata {
	return &implementations.PluginMetadataImpl{
		Name:     pmeta.GetName(),
		Version:  pmeta.GetVersion(),
		Size:     pmeta.GetSize(),
		URL:      pmeta.GetURL(),
		Category: pmeta.GetCategory(),
		Active:   pmeta.GetActive(),
	}
}
