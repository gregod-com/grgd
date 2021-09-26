package grgd

import (
	"os"
	"time"

	"github.com/gregod-com/grgd/cmd"
	"github.com/gregod-com/grgd/cmd/flags"
	"github.com/gregod-com/grgd/interfaces"
	"github.com/gregod-com/grgd/pkg/helper"
	"github.com/gregod-com/grgd/view"

	"github.com/urfave/cli/v2"
)

func NewApp(core interfaces.ICore, name string, version string, hooks map[string]func(*cli.Context) error) {
	logger := core.GetLogger()

	app := cli.NewApp()
	app.Name = name
	app.Usage = "custom grgd cli"
	app.Version = version
	app.Metadata = make(map[string]interface{})
	app.Metadata["core"] = core
	for k, v := range hooks {
		app.Metadata[k] = v
	}
	app.Flags = append(app.Flags, flags.GetFlags()...)
	app.CustomAppHelpTemplate = view.GetHelpTemplate()
	app.HideHelpCommand = true
	app.HideHelp = true
	app.UsageText = name + ` [global options] command [command options] [arguments...]`

	// define behavior before every command execution
	app.Before = func(c *cli.Context) error {
		core := helper.GetExtractor().GetCore(c)
		UI := core.GetUI()
		core.GetConfig().SetActiveProfile(c.String("profile"))
		log := core.GetLogger()
		u := core.GetUpdater()
		if u != nil {
			if err := core.GetUpdater().CheckSinceLastUpdate(app.Version, core); err != nil {
				log.Warnf("Error checking for update: %s ", err.Error())
			}
		}
		UI.ClearScreen(c)
		UI.PrintBanner(c)

		return nil
	}

	// define native commands available also without plugins
	app.Commands = append(app.Commands, cmd.GetCommands(app, core)...)

	// append native commands with commands found in loaded plugins
	for _, plug := range core.GetCMDPlugins() {
		switch arr := plug.GetCommands(app).(type) {
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
		return nil
	}

	// sort.Sort(cli.FlagsByName(app.Flags))

	apperr := app.Run(os.Args)
	if apperr != nil {
		logger.Fatal(apperr)
	}
	logger.Tracef("took: %vms\n", time.Since(core.GetStartTime()).Milliseconds())
}
