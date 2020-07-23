// This package implements the cli for the iam-stack. The underlying framework is depnedent upon urfave/cli.
package main

import (
	"io/ioutil"
	"log"
	"os"
	"plugin"
	"strings"

	"github.com/gregod-com/grgdplugincontracts"
	plugContracts "github.com/gregod-com/grgdplugincontracts"

	idx "github.com/gregod-com/grgd/pluginindex"
)

// LoadPlugins ...
func LoadPlugins(pluginFolder string) ([]plugContracts.ICMDPlugin, plugContracts.IUIPlugin) {
	var loadedUIPlugin plugContracts.IUIPlugin
	var loadedCMDPlugins []plugContracts.ICMDPlugin
	var allActiveMetaData []plugContracts.IPluginMetadata
	var allAvailablePlugins []plugContracts.IPluginMetadata

	pluginBinariesFolder := pluginFolder + "binaries/"

	index := idx.CreatePluginIndex(pluginFolder + "index.yaml")

	fileinfo, err := ioutil.ReadDir(pluginBinariesFolder)
	if err != nil {
		log.Fatal(err)
	}

	// iterate over plugin implementations
	for _, f := range fileinfo {
		pluginPath := pluginBinariesFolder + f.Name()

		if !strings.HasSuffix(pluginPath, ".so") || strings.HasPrefix(pluginPath, ".") {
			continue
		}

		// open .so file and error if something goes wrong
		pluginImpl, err := plugin.Open(pluginPath)
		if err != nil {
			log.Println(err)
			continue
		}

		// check if there is a var or func called `Plugin` in .so file
		symPlugin, err := pluginImpl.Lookup("Plugin")
		if err != nil {
			log.Println(err)
			continue
		}

		// check if the var/func is implementing the grgd plugin interface
		grgdplugin, ok := symPlugin.(plugContracts.IGrgdPlugin)
		if !ok {
			log.Println("Unexpected type from module symbol in plugin at " + pluginPath)
			continue
		}

		metadata, ok := grgdplugin.Init(nil).(grgdplugincontracts.IPluginMetadata)
		if !ok {
			log.Printf("Unexpected implementation of interface IPluginMetadata in plugin %T: %T => %v", grgdplugin, grgdplugin.GetMetaData(nil), grgdplugin.GetMetaData(nil))
			continue
		}

		switch x := index.AddPlugin(metadata); x {
		case "commands":
			log.Println("Found command " + metadata.GetIdentifier())
			loadedCMDPlugins = append(loadedCMDPlugins, grgdplugin.(grgdplugincontracts.ICMDPlugin))
			allActiveMetaData = append(allActiveMetaData, metadata)
		case "ui":
			log.Println("Found ui " + metadata.GetIdentifier())
			loadedUIPlugin, ok = grgdplugin.(plugContracts.IUIPlugin)
			if !ok {
				log.Println("Plugin does not implement IUIPlugin")
				continue
			}
			allActiveMetaData = append(allActiveMetaData, metadata)
		case "remove":
			err := os.Remove(pluginPath)
			if err != nil {
				log.Fatal(err)
			}
			log.Printf("Plugin %v successfully deleted", metadata.GetName())
		case "disabled":
			allAvailablePlugins = append(allAvailablePlugins, metadata)
		default:
			log.Printf("Unknown category %v\n", x)
			allAvailablePlugins = append(allAvailablePlugins, metadata)
		}
	}

	index.Finalize(allActiveMetaData, allAvailablePlugins)

	if loadedUIPlugin == nil {
		log.Println("No UI, using fallback UI")
		loadedUIPlugin = &fallbackui{}
	}
	return loadedCMDPlugins, loadedUIPlugin
}
