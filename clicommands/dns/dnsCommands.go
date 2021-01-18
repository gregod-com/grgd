package dns

import (
	"github.com/urfave/cli/v2"
)

// GetCLICommands ...
func GetCLICommands(app *cli.App) []*cli.Command {
	return []*cli.Command{
		{
			Name:        "dns",
			Category:    "state / monitoring",
			Usage:       "View and edit DNS routing",
			Action:      AEnter,
			Description: "Edit DNS Settings for local development",
		},
	}
}
