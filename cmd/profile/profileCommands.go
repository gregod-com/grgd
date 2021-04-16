package profile

import (
	"github.com/gregod-com/grgd/interfaces"
	"github.com/urfave/cli/v2"
)

// GetCLICommands ...
func GetCLICommands(app *cli.App, core interfaces.ICore) []*cli.Command {
	return []*cli.Command{
		{
			Name:            "profile",
			Category:        "settings",
			Usage:           "Configuration for profile",
			HideHelpCommand: true,
			Flags:           app.Flags,
			Subcommands: []*cli.Command{
				{
					Name:    "list",
					Usage:   "Show all profiles",
					Aliases: []string{"ls"},
					Before:  nil,
					Flags:   app.Flags,
					Action:  AListProfiles,
				},
				{
					Name:   "delete",
					Usage:  "Delete a single profile",
					Before: nil,
					Flags:  app.Flags,
					Action: ADeleteProfile,
				},
			},
		},
	}
}
