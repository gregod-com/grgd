package interfaces

import (
	"time"

	"github.com/gregod-com/grgdplugincontracts"
)

// ICore ...
type ICore interface {
	GetStartTime() time.Time
	GetLogger() ILogger
	GetUI() grgdplugincontracts.IUIPlugin
	GetConfig() IConfigObject
	GetHelper() IHelper
	GetFileSystemManipulator() IFileSystemManipulator
	GetPluginLoader() IPluginLoader
	GetDownloader() IDownloader
	Get(i interface{}) error
	// var profilename string
	// var databasePath string
	// var pluginsPath string
	// var cnfg I.IConfigObject
	// var starttime time.Time
}
