package profile

import (
	"github.com/gregod-com/grgd/interfaces"
	"github.com/urfave/cli/v2"
)

// GetCLICommands ...
func GetCLICommands(app *cli.App, core interfaces.ICore) []*cli.Command {
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
