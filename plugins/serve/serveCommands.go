package serve

import (
	"github.com/gregod-com/grgd/interfaces"
	"github.com/urfave/cli/v2"
)

// GetCLICommands ...
func GetCLICommands(app *cli.App, core interfaces.ICore) []*cli.Command {
	return []*cli.Command{
		{
			Name:   "serve",
			Usage:  "serve via http",
			Flags:  app.Flags,
			Action: AServe,
		},
	}
}
