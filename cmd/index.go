package cmd

import (
	"github.com/gregod-com/grgd/cmd/hack"
	"github.com/gregod-com/grgd/cmd/profile"
	"github.com/gregod-com/grgd/cmd/project"
	"github.com/gregod-com/grgd/cmd/service"
	"github.com/gregod-com/grgd/cmd/update"
	"github.com/gregod-com/grgd/interfaces"

	cli "github.com/urfave/cli/v2"
)

// GetCommands ...
func GetCommands(app *cli.App, core interfaces.ICore) []*cli.Command {
	var myCommands []*cli.Command
	// myCommands = append(myCommands, config.GetCLICommands(app, core)...)
	myCommands = append(myCommands, hack.GetCLICommands(app, core)...)
	myCommands = append(myCommands, update.GetCLICommands(app, core)...)
	myCommands = append(myCommands, profile.GetCLICommands(app, core)...)
	myCommands = append(myCommands, project.GetCLICommands(app, core)...)
	myCommands = append(myCommands, service.GetCLICommands(app, core)...)
	// myCommands = append(myCommands, stack.GetCLICommands(app, core)...)

	return myCommands
}
