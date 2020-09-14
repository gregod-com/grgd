package controller

import (
	I "grgd/interfaces"
	"reflect"
	"time"

	"github.com/gregod-com/grgdplugincontracts"
)

// CreateCore ...
func CreateCore(
	start time.Time,
	logger I.ILogger,
	config I.IConfigObject,
	helper I.IHelper,
	fsmanipulator I.IFileSystemManipulator,
	pluginloader I.IPluginLoader,
	ui grgdplugincontracts.IUIPlugin,
	downloader I.IDownloader) I.ICore {
	core := &Core{}
	core.start = start
	core.logger = logger
	core.config = config
	core.helper = helper
	core.fsmanipulator = fsmanipulator
	core.pluginloader = pluginloader
	core.ui = ui
	core.downloader = downloader
	return core
}

// Core ...
type Core struct {
	start         time.Time
	logger        I.ILogger
	config        I.IConfigObject
	helper        I.IHelper
	fsmanipulator I.IFileSystemManipulator
	pluginloader  I.IPluginLoader
	ui            grgdplugincontracts.IUIPlugin
	downloader    I.IDownloader
}

// GetStartTime ...
func (c *Core) GetStartTime() time.Time {
	return c.start
}

// GetLogger ...
func (c *Core) GetLogger() I.ILogger {
	return c.logger
}

// GetConfig ...
func (c *Core) GetConfig() I.IConfigObject {
	return c.config
}

// GetHelper ...
func (c *Core) GetHelper() I.IHelper {
	return c.helper
}

// GetFileSystemManipulator ...
func (c *Core) GetFileSystemManipulator() I.IFileSystemManipulator {
	return c.fsmanipulator
}

// GetPluginLoader ...
func (c *Core) GetPluginLoader() I.IPluginLoader {
	return c.pluginloader
}

// GetUI ...
func (c *Core) GetUI() grgdplugincontracts.IUIPlugin {
	return c.ui
}

// GetUI ...
func (c *Core) GetDownloader() I.IDownloader {
	return c.downloader
}

// Get ...
func (c *Core) Get(i interface{}) error {
	r := reflect.TypeOf(c)
	r.FieldByName("start")
	return nil
}
