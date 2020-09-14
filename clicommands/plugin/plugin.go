package plugin

import (
	"github.com/urfave/cli/v2"
)

// GetCLICommands ...
func GetCLICommands(app *cli.App) []*cli.Command {
	return []*cli.Command{
		{
			Name:  "plugins",
			Usage: "Configuration for plugins",
			Flags: app.Flags,
			Subcommands: []*cli.Command{
				{
					Name:    "list",
					Usage:   "Show all plugins",
					Aliases: []string{"ls"},
					Action:  APluginList,
				},
				{
					Name:    "activate",
					Usage:   "Activate plugins",
					Aliases: []string{"a"},
					Action:  APluginActivate,
				},
			},
		},
	}
}
