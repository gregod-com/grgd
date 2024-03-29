package project

import (
	"github.com/gregod-com/grgd/interfaces"
	"github.com/urfave/cli/v2"
)

// GetCLICommands ...
func GetCLICommands(app *cli.App, core interfaces.ICore) []*cli.Command {
	return []*cli.Command{
		{
			Name:     "project",
			Category: "settings",
			Usage:    "Configuration for projects",
			Flags:    app.Flags,
			Subcommands: []*cli.Command{
				{
					Name:    "list",
					Usage:   "Show all projects",
					Aliases: []string{"ls"},
					Flags:   app.Flags,
					Action:  AListProject,
				},
				{
					Name:      "switch",
					Usage:     "Switch current project",
					Aliases:   []string{"current"},
					Flags:     app.Flags,
					Action:    ASwitchProject,
					ArgsUsage: "[project name]",
				},
				{
					Name:   "add",
					Usage:  "Add project",
					Flags:  app.Flags,
					Action: AAddProject,
				},
				{
					Name:    "delete",
					Usage:   "Delete project",
					Aliases: []string{"del"},
					Flags:   app.Flags,
					Action:  ADeleteProject,
				},
				{
					Name:   "edit",
					Usage:  "Edit project",
					Flags:  app.Flags,
					Action: AEditProject,
				},
			},
		},
	}
}
