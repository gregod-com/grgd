package profile

import (
	"github.com/urfave/cli/v2"
)

// GetCLICommands ...
func GetCLICommands(app *cli.App) []*cli.Command {
	return []*cli.Command{
		{
			Name:     "profile",
			Category: "settings",
			Usage:    "Configuration for profile",
			Flags:    app.Flags,
			Subcommands: []*cli.Command{
				{
					Name:   "show",
					Usage:  "Show all profile infos",
					Flags:  app.Flags,
					Action: AListProfiles,
				},
			},
		},
	}
}
