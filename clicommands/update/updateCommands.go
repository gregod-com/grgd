package update

import (
	"github.com/urfave/cli/v2"
)

// GetCLICommands ...
func GetCLICommands(app *cli.App) []*cli.Command {
	return []*cli.Command{
		{
			Name:     "update",
			Category: "settings",
			Usage:    "Check and load updates",
			Action:   AUpdate,
		},
	}
}
