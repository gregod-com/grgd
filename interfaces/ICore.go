package interfaces

import (
	"time"
)

// ICore ...
type ICore interface {
	GetStartTime() time.Time
	GetLogger() ILogger
	GetUI() IUIPlugin
	GetConfig() IConfig
	GetHelper() IHelper
	GetNetworker() INetworker
	GetUpdater() IUpdater
	GetCMDPlugins() []ICMDPlugin
	Get(i interface{}) error
}
