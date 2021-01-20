package interfaces

// IPluginLoader ...
type IPluginLoader interface {
	LoadPlugins(pluginFolder string) ([]ICMDPlugin, IUIPlugin)
	// LoadHack(scriptsFolder string) []grgdplugincontracts.ICMDPlugin
}
