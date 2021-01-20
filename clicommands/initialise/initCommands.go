package initialise

import (
	"github.com/gregod-com/grgd/interfaces"
	"github.com/urfave/cli/v2"
)

// GetCLICommands ...
func GetCLICommands(app *cli.App, core interfaces.ICore) []*cli.Command {
	return []*cli.Command{
		{
			Name:            "initialise",
			Category:        "settings",
			Usage:           "init CLI",
			Aliases:         []string{"ini"},
			Flags:           app.Flags,
			HideHelpCommand: true,
			Action:          AInit,
		},
	}
}
