// This package implements the cli for the iam-stack. The underlying framework is depnedent upon urfave/cli.
package main

import (
	"log"
	"os"
	"sort"
	"time"

	"github.com/gregod-com/grgd/clicommands"
	"github.com/gregod-com/grgd/clicommands/flags"
	"github.com/gregod-com/grgd/controller/config"
	"github.com/gregod-com/grgd/controller/helper"
	"github.com/gregod-com/grgd/controller/pluginindex"
	"github.com/gregod-com/grgd/core"
	"github.com/gregod-com/grgd/gormdal"
	"github.com/gregod-com/grgd/interfaces"
	"github.com/gregod-com/grgd/logger"
	"github.com/gregod-com/grgd/view"

	"github.com/urfave/cli/v2"
)

func main() {
	dependecies := map[string]interface{}{
		"IHelper":                helper.ProvideHelper,
		"IUIPlugin":              view.ProvideFallbackUI,
		"ILogger":                logger.ProvideLogrusLogger,
		"IFileSystemManipulator": helper.ProvideFSManipulator,
		"IUpdater":               helper.ProvideUpdater,
		"IDAL":                   gormdal.ProvideDAL,
		"IDownloader":            helper.ProvideDownloader,
		"IConfig":                config.ProvideConfig,
		"IPluginIndex":           pluginindex.ProvidePluginIndex,
		"IPluginLoader":          helper.ProvidePluginLoader,
	}

	core := core.RegisterDependecies(dependecies)

	app := cli.NewApp()
	app.Name = "grgd"
	app.Usage = "grgd cli"
	app.Version = "0.9.0"
	app.Metadata = make(map[string]interface{})
	app.Metadata["core"] = core
	app.Flags = append(app.Flags, flags.GetFlags()...)
	app.CustomAppHelpTemplate = view.GetHelpTemplate()
	app.HideHelpCommand = true

	// define behavior before every command execution
	app.Before = func(c *cli.Context) error {
		UIPlugin := helper.GetExtractor().GetCore(c).GetUI()
		UIPlugin.ClearScreen(c)
		UIPlugin.PrintBanner(c)
		UIPlugin.Println("\u001b[33m", c.App.Version, "\u001b[0m")
		return nil
	}

	// define native commands available also without plugins
	app.Commands = append(app.Commands, clicommands.GetCommands(app)...)

	// append native commands with commands found in loaded plugins
	for _, plug := range core.GetCMDPlugins() {
		switch arr := plug.GetCommands(nil).(type) {
		case []*cli.Command:
			app.Commands = append(app.Commands, arr...)
		default:
			core.GetLogger().Error("Commands are not implemented as []*cli.Command but %T", arr)
		}
	}

	// define behavior after every command execution
	app.After = func(c *cli.Context) error {
		core := helper.GetExtractor().GetCore(c)
		logger := core.GetLogger()
		// cnfg := core.GetConfig()
		// cnfg.Save()
		logger.Trace(time.Since(core.GetStartTime()))
		return nil
	}

	sort.Sort(cli.FlagsByName(app.Flags))

	apperr := app.Run(os.Args)
	if apperr != nil {
		log.Fatal(apperr)
	}
}
