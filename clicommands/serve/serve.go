package serve

import (
	"github.com/urfave/cli/v2"
)

// GetCLICommands ...
func GetCLICommands(app *cli.App) []*cli.Command {
	return []*cli.Command{
		{
			Name:   "serve",
			Usage:  "serve via http",
			Flags:  app.Flags,
			Action: AServe,
		},
	}
}
