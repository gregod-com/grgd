package config

import (
	"github.com/gregod-com/grgd/interfaces"
	"github.com/urfave/cli/v2"
)

// GetCLICommands ...
func GetCLICommands(app *cli.App, core interfaces.ICore) []*cli.Command {
	return []*cli.Command{
		{
			Name:            "config",
			Category:        "settings",
			Usage:           "view and edit current configuration",
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
					Name:   "edit",
					Usage:  "edit the config file",
					Action: SubAConfigEdit,
					Flags:  app.Flags,
				},
			},
		},
	}
}
