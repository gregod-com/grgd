// This package implements the cli for the iam-stack. The underlying framework is depnedent upon urfave/cli.
package main

import (
	"log"
	"os"
	"sort"
	"time"

	"github.com/gregod-com/grgd/cmd"
	"github.com/gregod-com/grgd/cmd/flags"
	"github.com/gregod-com/grgd/core"
	"github.com/gregod-com/grgd/interfaces"
	"github.com/gregod-com/grgd/pkg/config"
	"github.com/gregod-com/grgd/pkg/gormdal"
	"github.com/gregod-com/grgd/pkg/helper"
	"github.com/gregod-com/grgd/pkg/logger"
	"github.com/gregod-com/grgd/pkg/profile"
	"github.com/gregod-com/grgd/view"

	"github.com/urfave/cli/v2"
)

func main() {
	dependecies := map[string]interface{}{
		"IHelper":    helper.ProvideHelper,
		"IUIPlugin":  view.ProvideFallbackUI,
		"ILogger":    logger.ProvideLogrusLogger,
		"INetworker": helper.ProvideNetworker,
		"IDAL":       gormdal.ProvideDAL,
		"IConfig":    config.ProvideConfig,
		"IProfile":   profile.ProvideProfile,
	}

	core := core.RegisterDependecies(dependecies)
	logger := core.GetLogger()

	app := cli.NewApp()
	app.Name = "grgd"
	app.Usage = "grgd cli"
	app.Version = "0.14.1"
	app.Metadata = make(map[string]interface{})
	app.Metadata["core"] = core
	app.Flags = append(app.Flags, flags.GetFlags()...)
	app.CustomAppHelpTemplate = view.GetHelpTemplate()
	app.HideHelpCommand = true

	// define behavior before every command execution
	app.Before = func(c *cli.Context) error {
		core := helper.GetExtractor().GetCore(c)
		UI := core.GetUI()
		var networker interfaces.INetworker
		helper.GetExtractor().GetCore(c).Get(&networker)

		connections := map[string]interface{}{
			"first": &helper.Connection{
				Endpoint: "https://www.google.com",
				TimeOut:  100,
				Success:  false,
			},
		}

		networker.CheckConnections(connections)
		for _, v := range connections {
			logger.Info(v)
		}

		UI.ClearScreen(c)
		UI.PrintBanner(c)
		return nil
	}

	// define native commands available also without plugins
	app.Commands = append(app.Commands, cmd.GetCommands(app, core)...)

	// append native commands with commands found in loaded plugins
	for _, plug := range core.GetCMDPlugins() {
		switch arr := plug.GetCommands(nil).(type) {
		case []*cli.Command:
			app.Commands = append(app.Commands, arr...)
		default:
			logger.Errorf("Commands are not implemented as []*cli.Command but %T", arr)
		}
	}

	// define behavior after every command execution
	app.After = func(c *cli.Context) error {
		core := helper.GetExtractor().GetCore(c)
		cnfg := core.GetConfig()
		err := cnfg.Save()
		if err != nil {
			logger.Fatalf("So this has failes %v", err)
		}
		logger.Trace(time.Since(core.GetStartTime()))
		return nil
	}

	sort.Sort(cli.FlagsByName(app.Flags))

	apperr := app.Run(os.Args)
	if apperr != nil {
		log.Fatal(apperr)
	}
}
