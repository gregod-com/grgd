package core

import (
	"errors"
	"reflect"
	"time"

	"github.com/gregod-com/grgd/interfaces"
	"github.com/gregod-com/grgd/view"

	log "github.com/sirupsen/logrus"

	"github.com/gregod-com/grgd/controller/helper"

	"github.com/gregod-com/grgdplugincontracts"
)

// RegisterDependecies ...
func RegisterDependecies(implsTemp map[string]interface{}) interfaces.ICore {
	impls := make(map[string]interface{})
	impls["start"] = time.Now()
	iter, solved, solvedCurrent := 0, 0, -1

	for {
		iter++
		if solvedCurrent == 0 {
			log.Fatal("There seems to be a circular dependency...")
		}
		solvedCurrent = 0

		for target, elem := range implsTemp {
			if _, ok := impls[target]; ok {
				continue
			}

			if elem == nil {
				solvedCurrent++
				continue
			}

			switch reflect.ValueOf(elem).Kind() {
			case reflect.Ptr, reflect.Struct:
				// the implementation is provided directly
				impls[target] = elem
				solvedCurrent++
			case reflect.Func:
				// the implementation is provided via the provider function
				solvedCurrent += addDependecyFromProviderFunction(elem, impls)
			default:
				log.Warnf("Type %v is not supported for injection. Ignoring.",
					reflect.TypeOf(elem),
				)
				solvedCurrent++
			}
		}

		solved += solvedCurrent
		if solved >= len(implsTemp) {
			break
		}
	}

	core := &Core{implementations: impls}

	// TODO: store plugin information in database?
	// configImpl.GetPluginsDir()
	// fsmanipulatorImpl.CheckOrCreateFolder(pluginsPath, os.FileMode(uint32(0760)))
	// pluginsIndex := pluginindex.CreatePluginIndex(path.Join(pluginsPath, "index.yaml"))

	// var pl interfaces.IPluginLoader
	// core.Get(&pl)
	// fsmanipulatorImpl := core.GetFileSystemManipulator()
	// pluginsPath := fsmanipulatorImpl.HomeDir(".grgd", "plugins")
	// CMDPlugins, _ := pl.LoadPlugins(pluginsPath)
	// impls["commands"] = CMDPlugins

	for k, v := range impls {
		core.GetLogger().Tracef("%-25v ->\t%T", k, v)
	}
	return core
}

func addDependecyFromProviderFunction(elem interface{}, impls map[string]interface{}) int {
	typ := reflect.TypeOf(elem)

	injection := make([]reflect.Value, 0)
	postpone := false
	for dependecyNr := 0; dependecyNr < typ.NumIn(); dependecyNr++ {
		depKey := typ.In(dependecyNr)
		dep, ok := impls[depKey.Name()]
		if !ok {
			switch depKey.Kind() {
			case reflect.Slice:
				if depKey.Elem().Name() == "" {
					log.Printf("Impossible to inject unknown variadric interfaces")
					return 1
				}
				postpone = true
			default:
				// if argument was not found yet in impls => try again later
				postpone = true
			}
		}
		injection = append(injection, reflect.ValueOf(dep))
	}
	if !postpone {
		key := typ.Out(0).Name()
		if _, ok := impls[key]; !ok {
			impls[key] = reflect.ValueOf(elem).Call(injection)[0].Interface()
		}
		return 1
	}
	return 0
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
	key := typ.Elem().Name()
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
	a, ok := c.implementations["ILogger"].(interfaces.ILogger)
	if !ok {
		a = ProvideDefaultLogger()
		c.implementations["ILogger"] = a
	}
	return a
}

// GetUI ...
func (c *Core) GetUI() grgdplugincontracts.IUIPlugin {
	a, ok := c.implementations["IUIPlugin"].(grgdplugincontracts.IUIPlugin)
	if !ok {
		a = view.ProvideFallbackUI()
		c.implementations["IUIPlugin"] = a
	}
	return a
}

// GetConfig ...
func (c *Core) GetConfig() interfaces.IConfigObject {
	a, ok := c.implementations["IConfigObject"].(interfaces.IConfigObject)
	if !ok {
		log.Fatal("ConfigObject is nil")
	}
	return a
}

// GetHelper ...
func (c *Core) GetHelper() interfaces.IHelper {
	a, ok := c.implementations["IHelper"].(interfaces.IHelper)
	if !ok {
		a = helper.ProvideHelper()
		c.implementations["IHelper"] = a
	}
	return a
}

// GetFileSystemManipulator ...
func (c *Core) GetFileSystemManipulator() interfaces.IFileSystemManipulator {
	a, ok := c.implementations["IFileSystemManipulator"].(interfaces.IFileSystemManipulator)
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
