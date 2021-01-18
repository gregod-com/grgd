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
	GetConfig() IConfig
	GetHelper() IHelper
	GetCMDPlugins() []grgdplugincontracts.ICMDPlugin
	Get(i interface{}) error
}
