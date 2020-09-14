package interfaces

import (
	"github.com/gregod-com/grgdplugincontracts"
)

// IPluginLoader ...
type IPluginLoader interface {
	LoadPlugins(pluginFolder string, index grgdplugincontracts.IPluginIndex, fm IFileSystemManipulator) ([]grgdplugincontracts.ICMDPlugin, grgdplugincontracts.IUIPlugin)
}
