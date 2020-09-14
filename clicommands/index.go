package clicommands

import (
	"grgd/clicommands/config"
	"grgd/clicommands/plugin"
	"grgd/clicommands/profile"
	"grgd/clicommands/project"
	"grgd/clicommands/service"
	"grgd/clicommands/shortcut"

	"github.com/urfave/cli/v2"
)

// GetCommands ...
func GetCommands(app *cli.App) []*cli.Command {
	var myCommands []*cli.Command

	myCommands = append(myCommands, project.GetCLICommands(app)...)
	myCommands = append(myCommands, service.GetCLICommands(app)...)
	myCommands = append(myCommands, profile.GetCLICommands(app)...)
	myCommands = append(myCommands, plugin.GetCLICommands(app)...)
	myCommands = append(myCommands, config.GetCLICommands(app)...)
	myCommands = append(myCommands, shortcut.GetCLICommands(app)...)

	return myCommands
}
