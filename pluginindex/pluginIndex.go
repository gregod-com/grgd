package pluginindex

import (
	"fmt"
	"io/ioutil"
	"log"
	"time"

	I "github.com/gregod-com/interfaces"
	yaml "gopkg.in/yaml.v2"
)

// PluginIndex ...
type PluginIndex struct {
	path               string               `yaml:"a,omitempty"`
	PluginMetadataList []PluginMetadataImpl `yaml:"plugins5"`
	Lastchecked        time.Time            `yaml:"lastchecked"`
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
	return nil
}

// Update ...
func (yamlObj *PluginIndex) Update() error {
	yamlObj.Lastchecked = time.Now()
	newyaml, err := yaml.Marshal(yamlObj)
	if err != nil {
		return err
	}

	log.Println(yamlObj.PluginMetadataList)

	log.Println(string(newyaml))

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
	iamconf, err := ioutil.ReadFile(yamlObj.path)
	if err != nil {
		// yamlObj.Update()
		log.Println("A new config file has been created at " + yamlObj.path)
		log.Fatal("Run the 'init' command next to configure your stack.")
	}
	return iamconf
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
func (yamlObj *PluginIndex) GetPluginList() []I.IPluginMetadata {
	returnArray := []I.IPluginMetadata{}
	for _, v := range yamlObj.PluginMetadataList {
		returnArray = append(returnArray, v)
	}
	return returnArray
}

// AddPlugin ...
func (yamlObj *PluginIndex) AddPlugin(newplug I.IPluginMetadata) error {
	yamlObj.PluginMetadataList = append(yamlObj.PluginMetadataList, newplug)
	return nil
}
