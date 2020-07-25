// This package implements the cli for the iam-stack. The underlying framework is depnedent upon urfave/cli.
package main

import (
	"log"
	"os"
	"sort"
	"time"

	"github.com/gregod-com/grgd/commands"
	"github.com/gregod-com/grgd/commands/flags"
	"github.com/gregod-com/grgd/helpers"
	"github.com/gregod-com/grgd/implementations/config"
	"github.com/gregod-com/grgd/implementations/pluginindex"
	"github.com/gregod-com/grgd/templates"
	"github.com/gregod-com/interfaces"

	"github.com/urfave/cli/v2"
)

func main() {
	starttime := time.Now()
	app := cli.NewApp()
	homedir := helpers.HomeDir()
	pluginFolder := homedir + "/.grgd/plugins/"

	if _, notexistserr := os.Stat(pluginFolder); os.IsNotExist(notexistserr) {
		os.MkdirAll(pluginFolder, os.FileMode(uint32(0760)))
	}

	CMDPlugins, UIPlugin := helpers.LoadPlugins(pluginFolder, pluginindex.CreatePluginIndex(pluginFolder+"index.yaml"))
	app.Flags = append(app.Flags, flags.GetFlags(app)...)
	app.Name = "grgd"
	app.Usage = "written in go. Can be used as a sidekick to gregod-menu and gregod-doctor"
	app.Version = "0.7.0"
	app.Metadata = make(map[string]interface{})
	app.Metadata["startTime"] = starttime
	app.CustomAppHelpTemplate = templates.GetHelpTemplate()
	app.HideHelpCommand = true

	// define behavior before every command execution
	app.Before = func(c *cli.Context) error {
		configpath := homedir + "/.grgd/config.yml"
		c.App.Metadata["configpath"] = configpath
		c.App.Metadata["config"] = config.CreateConfigObjectYaml(configpath)

		// TODO: move all those configs into the config file
		c.App.Metadata["pluginIndex"] = homedir + "/.grgd/plugins/index.yaml"
		c.App.Metadata["remoteIndex"] = "https://s3.gregod.com/public/plugins/index.yaml"
		c.App.Metadata["AWS-REGION"] = "eu-central-1"
		c.App.Metadata["updatecheckinterval"] = time.Millisecond * 50

		c.App.Metadata["UIPlugin"] = UIPlugin

		helpers.SystemCheck(c)
		helpers.CheckUpdate(c)

		UIPlugin.ClearScreen(c)
		UIPlugin.PrintBanner(c)
		return nil
	}

	// define native commands available also without plugins
	app.Commands = append(app.Commands, commands.GetCommands(app)...)

	// append native commands with commands found in loaded plugins
	for _, plug := range CMDPlugins {
		app.Commands = append(app.Commands, plug.GetCommands(nil)...)
	}

	// define behavior after every command execution
	app.After = func(c *cli.Context) error {
		var cnfg interfaces.IConfigObject
		helpers.ExtractMetadataFatal(c.App.Metadata, "config", &cnfg)
		cnfg.Update()

		var start time.Time
		if err := helpers.ExtractMetadata(c.App.Metadata, "startTime", &start); err == nil {
			UIPlugin.Println(c, time.Since(start))
		}
		return nil
	}

	sort.Sort(cli.FlagsByName(app.Flags))

	apperr := app.Run(os.Args)
	if apperr != nil {
		log.Fatal(apperr)
	}
}
