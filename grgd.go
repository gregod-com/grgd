// This package implements the cli for the iam-stack. The underlying framework is depnedent upon urfave/cli.
package grgd

import (
	"os"
	"time"

	"github.com/gregod-com/grgd/cmd"
	"github.com/gregod-com/grgd/cmd/flags"
	"github.com/gregod-com/grgd/interfaces"
	"github.com/gregod-com/grgd/pkg/helper"

	"github.com/urfave/cli/v2"
)

func NewApp(core interfaces.ICore, name string, version string) {
	logger := core.GetLogger()

	app := cli.NewApp()
	app.Name = name
	app.Usage = "grgd cli"
	app.Version = version
	app.Metadata = make(map[string]interface{})
	app.Metadata["core"] = core
	app.Flags = append(app.Flags, flags.GetFlags()...)
	// app.CustomAppHelpTemplate = view.GetHelpTemplate()
	app.HideHelpCommand = true
	app.HideHelp = true

	// define behavior before every command execution
	app.Before = func(c *cli.Context) error {
		core := helper.GetExtractor().GetCore(c)
		UI := core.GetUI()
		core.GetConfig().SetActiveProfile(c.String("profile"))
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
