package config

import (
	"github.com/urfave/cli/v2"
)

// GetCLICommands ...
func GetCLICommands(app *cli.App) []*cli.Command {
	return []*cli.Command{
		{
			Name:            "config",
			Usage:           "view and edit current configuration",
			Aliases:         []string{"conf", "c"},
			Flags:           app.Flags,
			HideHelpCommand: true,
			Subcommands: []*cli.Command{
				{
					Name:    "yaml",
					Usage:   "print config file in yaml format",
					Aliases: []string{"y"},
					Action:  SubAConfigYAML,
					Flags:   app.Flags,
				},
				{
					Name:    "json",
					Usage:   "print config file in json format",
					Aliases: []string{"j"},
					Action:  SubAConfigJSON,
					Flags:   app.Flags,
				},
				{
					Name:    "edit",
					Usage:   "edit the config file",
					Aliases: []string{"e"},
					Action:  SubAConfigEdit,
					Flags:   app.Flags,
				},
			},
		},
	}
}
