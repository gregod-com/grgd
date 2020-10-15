package core

import (
	"errors"
	"grgd/interfaces"
	"log"
	"reflect"
	"time"

	"github.com/gregod-com/grgd/controller/helper"
	"github.com/gregod-com/grgdplugincontracts"
)

// RegisterDependecies ...
func RegisterDependecies(implsTemp []interface{}) interfaces.ICore {
	impls := make(map[string]interface{})
	impls["start"] = time.Now()

	iter := 0
	solved := 0
	solvedCurrent := -1
	for {
		iter++
		if solvedCurrent == 0 {
			log.Fatal("There seems to be a circular dependency...")
		}
		solvedCurrent = 0

		for _, elem := range implsTemp {
			if elem == nil {
				solvedCurrent++
				break
			}

			val := reflect.ValueOf(elem)
			typ := reflect.TypeOf(elem)

			switch val.Kind() {
			// 		// 		case reflect.Ptr:
			// 		// 			log.Printf("Found Ptr with val: %v, kind: %v, typeof %v", val, typ.Kind(), typ)
			// 		// 			intrfce := typ.Elem()
			// 		// 			log.Printf("Saving at %v", intrfce.String())

			// 		// 			if _, ok := impls[intrfce.String()]; !ok {
			// 		// 				impls[intrfce.String()] = reflect.ValueOf(elem).Elem().Interface()
			// 		// 				continue
			// 		// 			}

			// 		// 			implsTemp = append(implsTemp[:k], implsTemp[k:])
			// 		// 			break
			case reflect.Func:
				// 			// 		case reflect.Interface:
				// 			// 			log.Fatalf("Passed unknow type to register: %T", elem)
				// 			// 			implsTemp = append(implsTemp[:k], implsTemp[k:])
				// 			// 			break
			default:
				// 			// 			log.Printf("Found %T with val: %v, kind: %v, typeof %v", elem, val, typ.Kind(), typ)
				// 			// 			log.Printf("Saving at %v", typ.String())
				// 			// 			if _, ok := impls[typ.String()]; !ok {
				// 			// 				implsTemp = append(implsTemp[:k], implsTemp[k:])
				// 			// 				break
				// 			// 			}
				// 			// 			offset++
				// 			implsTemp = append(implsTemp[:k], implsTemp[k:])
				solvedCurrent++
				continue
			}

			// 		if _, ok := impls[typ.Out(0).String()]; ok {
			// 			continue
			// 		}
			injection := make([]reflect.Value, 0)
			postpone := false
			for dependecyNr := 0; dependecyNr < typ.NumIn(); dependecyNr++ {
				depKey := typ.In(dependecyNr).String()
				dep, ok := impls[depKey]
				if !ok {
					postpone = true
				}
				injection = append(injection, reflect.ValueOf(dep))
			}
			if !postpone {
				key := typ.Out(0).String()
				if _, ok := impls[key]; !ok {
					impls[key] = reflect.ValueOf(elem).Call(injection)[0].Interface()
				}
				solvedCurrent++
				continue
			}
		}

		solved += solvedCurrent
		if iter++; solved >= len(implsTemp) || iter > 100 {
			break
		} else {
			log.Printf("solvedcurrent %v solved %v of %v elem", solvedCurrent, solved, len(implsTemp))
		}

	}

	core := &Core{implementations: impls}

	// TODO: find elegant solution to update cli automatically

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
	// for k := range impls {
	// 	core.GetLogger().Trace(k)
	// }

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
		log.Fatal("no logger found")
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
		a = helper.ProvideHelper()
		c.implementations["interfaces.IHelper"] = a
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
