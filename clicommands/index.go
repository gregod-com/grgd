package clicommands

import (
	"github.com/gregod-com/grgd/clicommands/aws"
	"github.com/gregod-com/grgd/clicommands/config"
	"github.com/gregod-com/grgd/clicommands/hack"
	"github.com/gregod-com/grgd/clicommands/k8s"
	"github.com/gregod-com/grgd/clicommands/update"

	"github.com/urfave/cli/v2"
)

// GetCommands ...
func GetCommands(app *cli.App) []*cli.Command {
	var myCommands []*cli.Command

	myCommands = append(myCommands, aws.GetCLICommands(app)...)
	myCommands = append(myCommands, k8s.GetCLICommands(app)...)
	myCommands = append(myCommands, config.GetCLICommands(app)...)
	// myCommands = append(myCommands, initialise.GetCLICommands(app)...)
	// myCommands = append(myCommands, shortcut.GetCLICommands(app)...)
	myCommands = append(myCommands, update.GetCLICommands(app)...)
	myCommands = append(myCommands, hack.GetCLICommands(app)...)

	// myCommands = append(myCommands, project.GetCLICommands(app)...)
	// myCommands = append(myCommands, service.GetCLICommands(app)...)
	// myCommands = append(myCommands, profile.GetCLICommands(app)...)
	// myCommands = append(myCommands, plugin.GetCLICommands(app)...)
	// myCommands = append(myCommands, dns.GetCLICommands(app)...)
	// myCommands = append(myCommands, stack.GetCLICommands(app)...)

	return myCommands
}
