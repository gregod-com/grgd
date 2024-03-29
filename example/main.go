package main

import (
	"github.com/gregod-com/grgd"
	"github.com/gregod-com/grgd/core"
	"github.com/gregod-com/grgd/example/group1"
	"github.com/gregod-com/grgd/interfaces"
	"github.com/gregod-com/grgd/pkg/config"
	"github.com/gregod-com/grgd/pkg/gormdal"
	"github.com/gregod-com/grgd/pkg/helper"
	"github.com/gregod-com/grgd/pkg/logger"
	"github.com/gregod-com/grgd/pkg/profile"
	"github.com/gregod-com/grgd/view"
)

func main() {
	log := logger.ProvideLogger()
	dependecies := map[string]interface{}{
		"ILogger":     logger.ProvideLogger,
		"IConfig":     config.ProvideConfig,
		"IHelper":     helper.ProvideHelper,
		"INetworker":  helper.ProvideNetworker,
		"IDAL":        gormdal.ProvideDAL,
		"IProfile":    profile.ProvideProfile,
		"IUIPlugin":   view.ProvideFallbackUI,
		"my-commands": ProvideCommands,
	}
	core, err := core.RegisterDependecies(dependecies)
	if err != nil {
		log.Fatalf("Error with register dependencies: %s", err.Error())
	}

	grgd.NewApp(core, "example", "0.0.1", nil)
}

func ProvideCommands() []interfaces.ICMDPlugin {
	return []interfaces.ICMDPlugin{
		&group1.CMD{},
	}
}
