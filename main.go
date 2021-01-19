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
	"k8s.io/client-go/tools/clientcmd"

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
		"IPinger":                helper.ProvidePinger,
		"string":                 gormdal.ProvideDefaultDBPath,
	}

	core := core.RegisterDependecies(dependecies)
	logger := core.GetLogger()

	app := cli.NewApp()
	app.Name = "grgd"
	app.Usage = "grgd cli"
	app.Version = "0.12.1"
	clientCfg, err := clientcmd.NewDefaultClientConfigLoadingRules().Load()
	if err != nil {
		logger.Warn("Could not read kube config")
	}
	app.Metadata = make(map[string]interface{})
	app.Metadata["core"] = core
	app.Metadata["kubeContext"] = clientCfg.CurrentContext
	app.Flags = append(app.Flags, flags.GetFlags()...)
	app.CustomAppHelpTemplate = view.GetHelpTemplate()
	app.HideHelpCommand = true

	// define behavior before every command execution
	app.Before = func(c *cli.Context) error {
		core := helper.GetExtractor().GetCore(c)
		UI := core.GetUI()
		var pinger interfaces.IPinger
		helper.GetExtractor().GetCore(c).Get(&pinger)

		connections := map[string]interface{}{
			"first": &helper.Connection{
				Endpoint: "https://www.google.com",
				TimeOut:  100,
				Success:  false,
			},
		}

		pinger.CheckConnections(connections)
		for _, v := range connections {
			logger.Info(v)
		}

		UI.ClearScreen(c)
		UI.PrintBanner(c)
		return nil
	}

	// define native commands available also without plugins
	app.Commands = append(app.Commands, clicommands.GetCommands(app, core)...)

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
