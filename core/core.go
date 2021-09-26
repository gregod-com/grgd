package core

import (
	"errors"
	"fmt"
	"reflect"
	"time"

	"github.com/gregod-com/grgd/interfaces"
	"github.com/urfave/cli/v2"
)

var tempLogger interfaces.ILogger

// RegisterDependecies ...
func RegisterDependecies(implsTemp map[string]interface{}) (interfaces.ICore, error) {
	impls := make(map[string]interface{})
	impls["start"] = time.Now()
	solvedlast, solvedCurrent := 0, 0
	foundProvider, foundLogger := false, false

	// logger has to be provided
	tempLoggerProvider, foundProvider := implsTemp["ILogger"].(func() interfaces.ILogger)
	tempLogger, foundLogger = implsTemp["ILogger"].(interfaces.ILogger)
	if !foundProvider && !foundLogger {
		return nil, fmt.Errorf("the implementation for ILogger was not provided. Can not start application without logger (found %T )", implsTemp["ILogger"])
	}
	if foundProvider {
		tempLogger = tempLoggerProvider()
	}

	for {
		for target, elem := range implsTemp {
			tempLogger.Tracef("checking %s\n", target)

			if elem == nil {
				tempLogger.Tracef("elem in target %s is nil, ignoring\n", target)
				continue
			}

			switch reflect.ValueOf(elem).Kind() {
			case reflect.Ptr, reflect.Struct:
				// the implementation is provided directly
				tempLogger.Tracef("target %s is provided directly\n", target)
				impls[target] = elem
				implsTemp[target] = nil
			case reflect.Func:
				//
				tempLogger.Tracef("calling provider function for target %s \n", target)
				if addDependecyFromProviderFunction(elem, impls) {
					implsTemp[target] = nil
				}
			default:
				tempLogger.Warnf("Type %v is not supported for injection. Ignoring.", reflect.TypeOf(elem))
				implsTemp[target] = nil
			}
		}

		solvedCurrent = 0
		for _, v := range implsTemp {
			if v == nil {
				solvedCurrent++
			}
		}
		tempLogger.Tracef("resolved %v / %v dependecies", solvedCurrent, len(implsTemp))
		if solvedCurrent >= len(implsTemp) {
			break
		}
		if solvedlast == solvedCurrent {
			return nil, fmt.Errorf("there seems to be a circular dependency")
		}
		solvedlast = solvedCurrent
	}

	core := &Core{implementations: impls}

	for k, v := range impls {
		core.GetLogger().Tracef("%-25v ->\t%T", k, v)
	}
	return core, nil
}

func addDependecyFromProviderFunction(elem interface{}, impls map[string]interface{}) bool {
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
					tempLogger.Warnf("Impossible to inject unknown variadric interfaces")
					return true
				}
				postpone = true
			default:
				tempLogger.Tracef("...waiting for dependency %s", depKey.Name())
				// if argument was not found yet in impls => try again later
				postpone = true
			}
		}
		injection = append(injection, reflect.ValueOf(dep))
	}
	if !postpone {
		retArg := typ.Out(0)
		switch retArg.Kind() {
		case reflect.Slice:
			key := retArg.Elem().Name()
			key = "[]" + key + "s"
			if _, ok := impls[key]; !ok {
				tempLogger.Tracef("assigned %s to key %s ", reflect.ValueOf(elem).Call(injection)[0].Interface(), key)
				impls[key] = reflect.ValueOf(elem).Call(injection)[0].Interface()
			} else {
				tempLogger.Tracef("assigned %s to key %s ", reflect.ValueOf(elem).Call(injection)[0].Interface(), key)
				if slc, ok := impls[key].([]interfaces.ICMDPlugin); ok {
					if slc2, ok := reflect.ValueOf(elem).Call(injection)[0].Interface().([]interfaces.ICMDPlugin); ok {
						impls[key] = append(slc, slc2...)
					}
				}
			}
			return true
		default:
			key := retArg.Name()
			if _, ok := impls[key]; !ok {
				impls[key] = reflect.ValueOf(elem).Call(injection)[0].Interface()
			}
			return true
		}
	}
	return false
}

// Core ...
type Core struct {
	implementations map[string]interface{}
}

// GetStartTime ...
func (c *Core) GetStartTime() time.Time {
	impl, ok := c.implementations["start"].(time.Time)
	if !ok {
		c.GetLogger().Fatal("Implementation not set or wrong type!")
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
		return nil
	}
	return a
}

// GetUI ...
func (c *Core) GetUI() interfaces.IUIPlugin {
	a, ok := c.implementations["IUIPlugin"].(interfaces.IUIPlugin)
	if !ok {
		return nil
	}
	return a
}

// GetConfig ...
func (c *Core) GetConfig() interfaces.IConfig {
	a, ok := c.implementations["IConfig"].(interfaces.IConfig)
	if !ok {
		c.GetLogger().Fatalf("Config is nil")
	}
	return a
}

// GetHelper ...
func (c *Core) GetHelper() interfaces.IHelper {
	a, ok := c.implementations["IHelper"].(interfaces.IHelper)
	if !ok {
		return nil
		// a = helper.ProvideHelper(logger.ProvideLogrusLogger())
		// c.implementations["IHelper"] = a
	}
	return a
}

// GetCMDPlugins ...
func (c *Core) GetCMDPlugins() []interfaces.ICMDPlugin {
	cmds, ok := c.implementations["[]ICMDPlugins"].([]interfaces.ICMDPlugin)
	if ok {
		return cmds
	}
	return nil
}

// GetHelper ...
func (c *Core) GetNetworker() interfaces.INetworker {
	a, ok := c.implementations["INetworker"].(interfaces.INetworker)
	if !ok {
		return nil
		// a = helper.ProvideNetworker(logger.ProvideLogrusLogger())
		// c.implementations["INetworker"] = a
	}
	return a
}

// GetUpdater ...
func (c *Core) GetUpdater() interfaces.IUpdater {
	a, ok := c.implementations["IUpdater"].(interfaces.IUpdater)
	if !ok {
		return nil
	}
	return a
}

// GetUpdater ...
func (c *Core) CallPreHook(i interface{}) error {
	log := c.GetLogger()
	log.Trace()
	ctx, ok := i.(*cli.Context)
	if !ok {
		return fmt.Errorf("wrong type passed %T", i)
	}

	hook, ok := ctx.App.Metadata["hooks_pre_"+ctx.Command.Name].(func(*cli.Context) error)
	if !ok {
		log.Debugf("no pre hook for command %s found", ctx.Command.Name)
		return nil
	}
	log.Infof("executing pre hook for command %s", ctx.Command.Name)
	err := hook(ctx)
	if err != nil {
		return err
	}
	return nil
}

// GetUpdater ...
func (c *Core) CallPostHook(i interface{}) error {
	log := c.GetLogger()
	log.Trace()
	ctx, ok := i.(*cli.Context)
	if !ok {
		return fmt.Errorf("wrong type passed")
	}

	hook, ok := ctx.App.Metadata["hooks_post_"+ctx.Command.Name].(func(*cli.Context) error)
	if !ok {
		log.Debugf("no post hook for command %s found", ctx.Command.Name)
		return nil
	}
	log.Infof("executing post hook for command %s", ctx.Command.Name)
	err := hook(ctx)
	if err != nil {
		return err
	}
	return nil
}
