package cmd

import (
	"github.com/gregod-com/grgd/cmd/update"
	"github.com/gregod-com/grgd/interfaces"

	"github.com/urfave/cli/v2"
)

// GetCommands ...
func GetCommands(app *cli.App, core interfaces.ICore) []*cli.Command {
	var myCommands []*cli.Command
	// myCommands = append(myCommands, config.GetCLICommands(app, core)...)
	// myCommands = append(myCommands, hack.GetCLICommands(app, core)...)
	myCommands = append(myCommands, update.GetCLICommands(app, core)...)
	// myCommands = append(myCommands, profile.GetCLICommands(app, core)...)
	// myCommands = append(myCommands, project.GetCLICommands(app, core)...)
	// myCommands = append(myCommands, serve.GetCLICommands(app, core)...)
	// myCommands = append(myCommands, service.GetCLICommands(app, core)...)
	// myCommands = append(myCommands, stack.GetCLICommands(app, core)...)
	// myCommands = append(myCommands, initialise.GetCLICommands(app)...)

	return myCommands
}
