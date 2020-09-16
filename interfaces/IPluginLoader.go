package interfaces

import (
	"github.com/gregod-com/grgdplugincontracts"
)

// IPluginLoader ...
type IPluginLoader interface {
	LoadPlugins(pluginFolder string) ([]grgdplugincontracts.ICMDPlugin, grgdplugincontracts.IUIPlugin)
}
