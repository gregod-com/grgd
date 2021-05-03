package service

import (
	"github.com/gregod-com/grgd/interfaces"
	"github.com/urfave/cli/v2"
)

// GetCLICommands ...
func GetCLICommands(app *cli.App, core interfaces.ICore) []*cli.Command {
	return []*cli.Command{
		{
			Name:     "service",
			Category: "settings",
			Aliases:  []string{"s"},
			Usage:    "Configuration for services",
			Flags:    app.Flags,
			Subcommands: []*cli.Command{
				{
					Name:    "list",
					Usage:   "Show service",
					Aliases: []string{"ls"},
					Flags:   app.Flags,
					Action:  AListService,
				},
				{
					Name:   "add",
					Usage:  "Add service",
					Flags:  app.Flags,
					Action: AAddService,
				},
				{
					Name:    "delete",
					Usage:   "Delete service",
					Aliases: []string{"del"},
					Flags:   app.Flags,
					Action:  ADeleteService,
				},
				{
					Name:   "edit",
					Usage:  "Edit service",
					Flags:  app.Flags,
					Action: AEditService,
				},
				{
					Name:   "disable",
					Usage:  "Disable service",
					Flags:  app.Flags,
					Action: AToggleService,
				},
			},
		},
	}
}
