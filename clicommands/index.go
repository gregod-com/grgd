package clicommands

import (
	"github.com/gregod-com/grgd/clicommands/aws"
	"github.com/gregod-com/grgd/clicommands/config"
	"github.com/gregod-com/grgd/clicommands/hack"
	"github.com/gregod-com/grgd/clicommands/k8s"
	"github.com/gregod-com/grgd/clicommands/update"
	"github.com/gregod-com/grgd/interfaces"

	"github.com/urfave/cli/v2"
)

// GetCommands ...
func GetCommands(app *cli.App, core interfaces.ICore) []*cli.Command {
	var myCommands []*cli.Command

	myCommands = append(myCommands, aws.GetCLICommands(app)...)
	myCommands = append(myCommands, config.GetCLICommands(app)...)
	myCommands = append(myCommands, hack.GetCLICommands(app, core)...)
	myCommands = append(myCommands, k8s.GetCLICommands(app)...)
	// myCommands = append(myCommands, profile.GetCLICommands(app)...)
	// myCommands = append(myCommands, project.GetCLICommands(app)...)
	// myCommands = append(myCommands, serve.GetCLICommands(app)...)
	// myCommands = append(myCommands, service.GetCLICommands(app)...)
	// myCommands = append(myCommands, stack.GetCLICommands(app)...)
	myCommands = append(myCommands, update.GetCLICommands(app)...)

	// myCommands = append(myCommands, initialise.GetCLICommands(app)...)

	return myCommands
}
