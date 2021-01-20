package update

import (
	"github.com/gregod-com/grgd/interfaces"
	"github.com/urfave/cli/v2"
)

// GetCLICommands ...
func GetCLICommands(app *cli.App, core interfaces.ICore) []*cli.Command {
	return []*cli.Command{
		{
			Name:     "update",
			Category: "settings",
			Usage:    "Check and load updates",
			Action:   AUpdate,
		},
	}
}
