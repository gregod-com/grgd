package pluginindex

import (
	"bufio"
	"fmt"
	"grgd/controller/helper"
	"grgd/interfaces"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"

	"github.com/gregod-com/grgdplugincontracts"
	"github.com/urfave/cli/v2"
	"golang.org/x/mod/semver"
	"gopkg.in/yaml.v2"
)

// ProvidePluginIndex ...
func ProvidePluginIndex(config interfaces.IConfigObject, logger interfaces.ILogger) grgdplugincontracts.IPluginIndex {
	plugIdx := new(PluginIndex)
	plugIdx.config = config
	plugIdx.logger = logger
	return plugIdx
}

// PluginIndex ...
type PluginIndex struct {
	path           string
	config         interfaces.IConfigObject
	logger         interfaces.ILogger
	PluginSettings map[string]grgdplugincontracts.IPluginMetadata `yaml:"plugins"`
}

// CreatePluginIndexFromCLIContext ...
func CreatePluginIndexFromCLIContext(c *cli.Context) grgdplugincontracts.IPluginIndex {
	var pluginIndexPath string
	ext := helper.GetExtractor()
	ext.GetMetadataFatal(c.App.Metadata, "pluginIndex", &pluginIndexPath)
	return CreatePluginIndex(pluginIndexPath)
}

// CreatePluginIndex ...
func CreatePluginIndex(path string) grgdplugincontracts.IPluginIndex {
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

// UnmarshalYAML ...
func (yamlObj *PluginIndex) UnmarshalYAML(unmarshal func(interface{}) error) error {
	// var base map[string]map[string]PluginMetadataImpl

	// err := unmarshal(&base)
	// if err != nil {
	// 	return err
	// }

	// yamlObj.PluginSettings = make(map[string]grgdplugincontracts.IPluginMetadata, len(base["plugins"]))

	// // move all concrete implementations to interfacemap
	// for k := range base["plugins"] {
	// 	v := base["plugins"][k]
	// 	yamlObj.PluginSettings[k] = &v
	// }
	return nil
}

// Finalize ...
func (yamlObj *PluginIndex) Finalize(activePlugins []grgdplugincontracts.IPluginMetadata, availablePlugins []grgdplugincontracts.IPluginMetadata) error {
	// disable all plugins
	for k := range yamlObj.PluginSettings {
		yamlObj.PluginSettings[k].SetActive(false)
		yamlObj.PluginSettings[k].SetLoaded(false)
	}

	// enable all loaded plugins
	for _, v := range activePlugins {
		yamlObj.PluginSettings[v.GetIdentifier()].SetActive(true)
		yamlObj.PluginSettings[v.GetIdentifier()].SetLoaded(true)
	}

	// enable all loaded plugins
	for _, v := range availablePlugins {
		yamlObj.PluginSettings[v.GetIdentifier()].SetLoaded(true)
	}

	yamlObj.Update()
	return nil
}

// Update ...
func (yamlObj *PluginIndex) Update() error {
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
		yamlObj.Update()
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
	return time.Now()
}

// GetPluginList ...
func (yamlObj *PluginIndex) GetPluginList() map[string]grgdplugincontracts.IPluginMetadata {
	return yamlObj.PluginSettings
}

// GetPluginListActive ...
func (yamlObj *PluginIndex) GetPluginListActive() []grgdplugincontracts.IPluginMetadata {
	var keys []grgdplugincontracts.IPluginMetadata
	for _, v := range yamlObj.PluginSettings {
		if v.GetActive() && v.GetLoaded() {
			keys = append(keys, v)
		}
	}
	return keys
}

// GetPluginListInactive ...
func (yamlObj *PluginIndex) GetPluginListInactive() []grgdplugincontracts.IPluginMetadata {
	var keys []grgdplugincontracts.IPluginMetadata
	for _, v := range yamlObj.PluginSettings {
		if !v.GetActive() && v.GetLoaded() {
			keys = append(keys, v)
		}
	}
	return keys
}

// GetPluginListOffline ...
func (yamlObj *PluginIndex) GetPluginListOffline() []grgdplugincontracts.IPluginMetadata {
	var keys []grgdplugincontracts.IPluginMetadata
	for _, v := range yamlObj.PluginSettings {
		if !v.GetLoaded() {
			keys = append(keys, v)
		}
	}
	return keys
}

// ToggleActive ...
func (yamlObj *PluginIndex) ToggleActive(plugID string) bool {
	if p, ok := yamlObj.PluginSettings[plugID]; ok {
		p.ToggleActive()
		yamlObj.PluginSettings[plugID] = p
		yamlObj.Update()
		return p.GetActive()
	}
	return false
}

// AddPlugin ...
func (yamlObj *PluginIndex) AddPlugin(newplug grgdplugincontracts.IPluginMetadata) string {
	// log.Println("Checking plugin " + newplug.GetIdentifier())
	identifier := newplug.GetIdentifier()

	if plug, ok := yamlObj.PluginSettings[identifier]; ok {

		newplug.SetActive(plug.GetActive())

		switch semver.Compare(plug.GetVersion(), newplug.GetVersion()) {
		case -1:
			fmt.Printf("Update plugin %v from %v to %v? [y/n] ", newplug.GetName(), plug.GetVersion(), newplug.GetVersion())
			if ynQuestion() {
				yamlObj.PluginSettings[identifier] = newplug
				return newplug.GetCategory()
			}
			return "disabled"
		case 1:
			fmt.Printf("An older version for plugin %v (%v) was found. You should remove the plugin file? [y/n]", newplug.GetName(), newplug.GetVersion())
			if ynQuestion() {
				return "remove"
			}
			return "disabled"
		default:
			yamlObj.PluginSettings[identifier] = plug
			if plug.GetActive() {
				return plug.GetCategory()
			}
			return "disabled"
		}
	}
	return newplug.GetCategory()
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
