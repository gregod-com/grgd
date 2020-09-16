// This package implements the cli for the iam-stack. The underlying framework is depnedent upon urfave/cli.
package main

import (
	"log"
	"os"
	"sort"
	"time"

	"grgd/clicommands"
	"grgd/clicommands/flags"
	"grgd/controller/config"
	"grgd/controller/helper"
	"grgd/controller/pluginindex"
	"grgd/core"
	"grgd/interfaces"
	"grgd/logger"
	"grgd/persistence"
	"grgd/view"

	"github.com/urfave/cli/v2"
)

func main() {
	core, CMDPlugins := initCore()

	app := cli.NewApp()
	app.Name = "grgd"
	app.Usage = "grgd cli"
	app.Version = "0.8.0"
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
		return nil
	}

	// define native commands available also without plugins
	app.Commands = append(app.Commands, clicommands.GetCommands(app)...)

	// append native commands with commands found in loaded plugins
	for _, plug := range CMDPlugins {
		app.Commands = append(app.Commands, plug.GetCommands(nil)...)
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
