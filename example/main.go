// This package implements the cli for the iam-stack. The underlying framework is depnedent upon urfave/cli.
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
	dependecies := map[string]interface{}{
		"ILogger":     logger.ProvideLogrusLogger,
		"IConfig":     config.ProvideConfig,
		"IHelper":     helper.ProvideHelper,
		"INetworker":  helper.ProvideNetworker,
		"IDAL":        gormdal.ProvideDAL,
		"IProfile":    profile.ProvideProfile,
		"IUIPlugin":   view.ProvideFallbackUI,
		"my-commands": ProvideCommands,
	}
	core := core.RegisterDependecies(dependecies)
	grgd.NewApp(core, "example", "0.0.1")
}

func ProvideCommands() []interfaces.ICMDPlugin {
	return []interfaces.ICMDPlugin{
		&group1.CMD{},
	}
}
