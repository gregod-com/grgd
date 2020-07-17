// This package implements the cli for the iam-stack. The underlying framework is depnedent upon urfave/cli.
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"plugin"

	plugContracts "github.com/gregod-com/grgdplugincontracts"

	idx "github.com/gregod-com/grgd/pluginindex"
)

// LoadPlugins ...
func LoadPlugins(pluginFolder string) (map[string]plugContracts.IGrgdPlugin, plugContracts.IUIPlugin) {
	var loadedUIPlugin plugContracts.IUIPlugin
	loadedCMDPlugins := map[string]plugContracts.IGrgdPlugin{}
	pluginBinariesFolder := pluginFolder + "binaries/"

	index := idx.CreatePluginIndex(pluginFolder + "index.yaml")

	fileinfo, err := ioutil.ReadDir(pluginBinariesFolder)
	if err != nil {
		log.Fatal(err)
	}

	// iterate over plugin implementations
	for _, f := range fileinfo {
		pluginPath := pluginBinariesFolder + f.Name()

		// open .so file and error if something goes wrong
		pluginImpl, err := plugin.Open(pluginPath)
		if err != nil {
			log.Println(err)
			continue
		}

		// check if there is a var or func called `Plugin` in .so file
		symPlugin, err := pluginImpl.Lookup("Plugin")
		if err != nil {
			log.Fatal(err)
		}

		// check if the var/func is implementing the grgd plugin interface
		grgdplugin, ok := symPlugin.(plugContracts.IGrgdPlugin)
		if !ok {
			log.Println("Unexpected type from module symbol in Plugin at " + pluginPath)
			continue
		}

		metadata, ok := grgdplugin.Init(nil).(plugContracts.IPluginMetadata)
		if !ok {
			log.Printf("Unexpected implementation of interface IPluginMetadata in plugin %T: %T => %v", grgdplugin, grgdplugin.GetMetaData(nil), grgdplugin.Init(nil))
			continue
		}

		switch x := index.AddPlugin(metadata); x {
		case "commands":
			loadedCMDPlugins[metadata.GetName()] = grgdplugin
			index.Update()
		case "ui":
			loadedUIPlugin, ok = grgdplugin.(plugContracts.IUIPlugin)
			if !ok {
				fmt.Println("Plugin does not implement IUIPlugin")
			}
		case "remove":
			err := os.Remove(pluginPath)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("Plugin %v successfully deleted", metadata.GetName())
		case "disabled":
			// fmt.Println("Ignoring Plugin")
		default:
			fmt.Printf("Unknown category %v\n", x)
		}

		if err := index.Update(); err != nil {
			log.Fatal(err.Error())
		}
	}
	return loadedCMDPlugins, loadedUIPlugin
}
