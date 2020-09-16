package core

import (
	"errors"
	"grgd/interfaces"
	"log"
	"reflect"
	"time"

	"github.com/gregod-com/grgdplugincontracts"
)

// RegisterDependecies ...
func RegisterDependecies(implsTemp []interface{}) interfaces.ICore {
	impls := make(map[string]interface{})
	impls["start"] = time.Now()

	iter := 0
	for {
		for _, v := range implsTemp {
			providerFunction := reflect.TypeOf(v)
			if _, ok := impls[providerFunction.Out(0).String()]; ok {
				continue
			}
			injection := make([]reflect.Value, 0)
			postpone := false
			for dependecyNr := 0; dependecyNr < providerFunction.NumIn(); dependecyNr++ {
				depKey := providerFunction.In(dependecyNr).String()
				dep, ok := impls[depKey]
				if !ok {
					postpone = true
					// log.Fatal("Wrong DEP Order!!!!, " + depKey + " not found or provided after " + providerFunction.Out(0).String())
				}
				injection = append(injection, reflect.ValueOf(dep))
			}
			if !postpone {
				impls[providerFunction.Out(0).String()] = reflect.ValueOf(v).Call(injection)[0].Interface()
			}
		}

		if iter++; len(implsTemp) < iter {
			log.Fatal("There seems to be a circular dependency...")
		}
		if len(impls)-1 == len(implsTemp) {
			break
		}
	}

	core := &Core{implementations: impls}

	// TODO: find elegant solution to update cli automatically

	// TODO: store plugin information in database?
	// configImpl.GetPluginsDir()
	// fsmanipulatorImpl.CheckOrCreateFolder(pluginsPath, os.FileMode(uint32(0760)))
	// pluginsIndex := pluginindex.CreatePluginIndex(path.Join(pluginsPath, "index.yaml"))

	var pl interfaces.IPluginLoader
	core.Get(&pl)
	fsmanipulatorImpl := core.GetFileSystemManipulator()
	pluginsPath := fsmanipulatorImpl.HomeDir(".grgd", "plugins")
	CMDPlugins, _ := pl.LoadPlugins(pluginsPath)
	impls["commands"] = CMDPlugins
	for k := range impls {
		core.GetLogger().Trace(k)
	}

	// upImpl.CheckUpdate(core)
	return core
}

// Core ...
type Core struct {
	implementations map[string]interface{}
}

// GetStartTime ...
func (c *Core) GetStartTime() time.Time {
	impl, ok := c.implementations["start"].(time.Time)
	if !ok {
		log.Fatal("Implementation not set or wrong type!")
	}
	return impl
}

// Get ...
func (c *Core) Get(i interface{}) error {
	val := reflect.ValueOf(i)
	typ := reflect.TypeOf(i)
	key := typ.Elem().String()
	logimpl, ok := c.implementations[key]
	if !ok {
		return errors.New("Could not find implementation for interface " + key)
	}

	if !reflect.TypeOf(logimpl).Implements(typ.Elem()) {
		return errors.New("Passed corrupt implementation for " + key + " to DI container.")
	}
	val.Elem().Set(reflect.ValueOf(logimpl))

	return nil
}

// GetLogger ...
func (c *Core) GetLogger() interfaces.ILogger {
	a, ok := c.implementations["interfaces.ILogger"].(interfaces.ILogger)
	if !ok {
		return nil
	}
	return a
}

// GetUI ...
func (c *Core) GetUI() grgdplugincontracts.IUIPlugin {
	a, ok := c.implementations["grgdplugincontracts.IUIPlugin"].(grgdplugincontracts.IUIPlugin)
	if !ok {
		return nil
	}
	return a
}

// GetConfig ...
func (c *Core) GetConfig() interfaces.IConfigObject {
	a, ok := c.implementations["interfaces.IConfigObject"].(interfaces.IConfigObject)
	if !ok {
		return nil
	}
	return a
}

// GetHelper ...
func (c *Core) GetHelper() interfaces.IHelper {
	a, ok := c.implementations["interfaces.IHelper"].(interfaces.IHelper)
	if !ok {
		return nil
	}
	return a
}

// GetFileSystemManipulator ...
func (c *Core) GetFileSystemManipulator() interfaces.IFileSystemManipulator {
	a, ok := c.implementations["interfaces.IFileSystemManipulator"].(interfaces.IFileSystemManipulator)
	if !ok {
		return nil
	}
	return a
}

// GetCMDPlugins ...
func (c *Core) GetCMDPlugins() []grgdplugincontracts.ICMDPlugin {
	cmds, ok := c.implementations["commands"].([]grgdplugincontracts.ICMDPlugin)
	if ok {
		return cmds
	}
	return nil
}
