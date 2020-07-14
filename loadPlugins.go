// This package implements the cli for the iam-stack. The underlying framework is depnedent upon urfave/cli.
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"plugin"

	idx "github.com/gregod-com/grgd/pluginindex"
	I "github.com/gregod-com/interfaces"
)

// LoadPlugins ...
func LoadPlugins(pluginFolder string) map[string]I.IGrgdPlugin {
	loadedPlugins := map[string]I.IGrgdPlugin{}
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
			log.Fatal(err)
		}

		// check if there is a var or func called `Plugin` in .so file
		symPlugin, err := pluginImpl.Lookup("Plugin")
		if err != nil {
			log.Fatal(err)
		}

		// check if the var/func is implementing the grgd plugin interface
		grgdplugin, ok := symPlugin.(I.IGrgdPlugin)
		if !ok {
			log.Fatal("Unexpected type from module symbol in Plugin at " + pluginPath)
		}

		// check if the init method returns a valid plugin metatda interface
		metadata, ok := grgdplugin.Init(pluginPath).(I.IPluginMetadata)
		if !ok {
			log.Fatal("Unexpected implementation of interface IPluginMetadata")
		}

		index.AddPlugin(metadata)
		fmt.Println(index.GetPluginList())
		if err := index.Update(); err != nil {
			log.Fatal(err.Error())
		}
		os.Exit(0)
		loadedPlugins[metadata.GetName()] = grgdplugin

	}
	return loadedPlugins
}
