package helper

import (
	"grgd/interfaces"
	"io/ioutil"
	"log"
	"os"
	"path"
	"plugin"
	"strings"

	"github.com/gregod-com/grgdplugincontracts"
)

// ProvidePluginLoader ...
func ProvidePluginLoader(
	fsm interfaces.IFileSystemManipulator,
	config interfaces.IConfigObject,
	index grgdplugincontracts.IPluginIndex,
	logger interfaces.ILogger,
) interfaces.IPluginLoader {
	return &PluginLoader{fsm: fsm, index: index, config: config, logger: logger}
}

// PluginLoader ...
type PluginLoader struct {
	fsm    interfaces.IFileSystemManipulator
	index  grgdplugincontracts.IPluginIndex
	config interfaces.IConfigObject
	logger interfaces.ILogger
}

// LoadPlugins ...
func (pl *PluginLoader) LoadPlugins(pluginFolder string) ([]grgdplugincontracts.ICMDPlugin, grgdplugincontracts.IUIPlugin) {

	var loadedUIPlugin grgdplugincontracts.IUIPlugin
	var loadedCMDPlugins []grgdplugincontracts.ICMDPlugin
	var allActiveMetaData []grgdplugincontracts.IPluginMetadata
	var allAvailablePlugins []grgdplugincontracts.IPluginMetadata

	pluginBinariesFolder := path.Join(pluginFolder, "binaries")
	pluginBinariesFolderDisabled := path.Join(pluginBinariesFolder, "disabled")

	pl.fsm.CheckOrCreateFolder(pluginBinariesFolder, 0755)
	pl.fsm.CheckOrCreateFolder(pluginBinariesFolderDisabled, 0755)

	fileinfo, err := ioutil.ReadDir(pluginBinariesFolder)
	if err != nil {
		pl.logger.Fatal(err)
	}

	// iterate over plugin implementations
	for _, f := range fileinfo {
		pluginPath := path.Join(pluginBinariesFolder, f.Name())

		if !strings.HasSuffix(pluginPath, ".so") || strings.HasPrefix(pluginPath, ".") {
			continue
		}

		// open .so file and error if something goes wrong
		pluginImpl, err := plugin.Open(pluginPath)
		if err != nil {
			pl.logger.Error(err)
			os.Rename(pluginPath, path.Join(pluginBinariesFolderDisabled, f.Name()))
			pl.logger.Info("moved plugin at %v to `disabled folder` at %v since the build is not compatible with current version of the cli\n", pluginPath, pluginBinariesFolderDisabled)
			pl.logger.Fatal("invoke cli again to start with the remaining plugins")
			continue
		}

		// check if there is a var or func called `Plugin` in .so file
		symPlugin, err := pluginImpl.Lookup("Plugin")
		if err != nil {
			pl.logger.Trace(err)
			continue
		}

		// check if the var/func is implementing the grgd plugin interface
		grgdplugin, ok := symPlugin.(grgdplugincontracts.IGrgdPlugin)
		if !ok {
			pl.logger.Error("Unexpected type from module symbol in plugin at " + pluginPath)
			continue
		}

		metadata, ok := grgdplugin.Init(nil).(grgdplugincontracts.IPluginMetadata)
		if !ok {
			pl.logger.Error("Unexpected implementation of interface IPluginMetadata in plugin %T: %T => %v", grgdplugin, grgdplugin.GetMetaData(nil), grgdplugin.GetMetaData(nil))
			continue
		}

		switch x := pl.index.AddPlugin(metadata); x {
		case "commands":
			pl.logger.Trace("Found command " + metadata.GetIdentifier())
			loadedCMDPlugins = append(loadedCMDPlugins, grgdplugin.(grgdplugincontracts.ICMDPlugin))
			allActiveMetaData = append(allActiveMetaData, metadata)
		case "ui":
			pl.logger.Trace("Found ui " + metadata.GetIdentifier())
			loadedUIPlugin, ok = grgdplugin.(grgdplugincontracts.IUIPlugin)
			if !ok {
				pl.logger.Error("Plugin does not implement IUIPlugin")
				continue
			}
			allActiveMetaData = append(allActiveMetaData, metadata)
		case "remove":
			err := os.Remove(pluginPath)
			if err != nil {
				log.Fatal(err)
			}
			pl.logger.Error("Plugin %v successfully deleted", metadata.GetName())
		case "disabled":
			allAvailablePlugins = append(allAvailablePlugins, metadata)
		default:
			pl.logger.Error("Unknown category %v\n", x)
			allAvailablePlugins = append(allAvailablePlugins, metadata)
		}
	}

	// pl.index.Finalize(allActiveMetaData, allAvailablePlugins)
	return loadedCMDPlugins, loadedUIPlugin
}
