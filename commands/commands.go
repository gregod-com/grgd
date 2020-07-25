package commands

import (
	A "github.com/gregod-com/grgd/actions"
	T "github.com/gregod-com/grgd/templates"
	"github.com/urfave/cli/v2"
)

// GetCommands ...
func GetCommands(app *cli.App) []*cli.Command {
	return []*cli.Command{
		{
			Name:        "init",
			Usage:       "Initialze the " + app.Name,
			Flags:       app.Flags,
			Action:      A.AInit,
			Description: T.Description(app, "init"),
		},
		{
			Name:        "plugins",
			Usage:       "Configuration for plugins",
			Aliases:     []string{"p"},
			Flags:       app.Flags,
			Description: T.Description(app, "config"),
			Subcommands: []*cli.Command{
				{
					Name:        "list",
					Usage:       "Show all plugins",
					Aliases:     []string{"ls"},
					Action:      A.APluginList,
					Description: T.Description(app, "config-yaml"),
				},
				{
					Name:        "activate",
					Usage:       "Activate plugins",
					Aliases:     []string{"a"},
					Action:      A.APluginActivate,
					Description: T.Description(app, "config-yaml"),
				},
			},
		},
		{
			Name:            "config",
			Usage:           "view and edit current configuration",
			Aliases:         []string{"conf", "c"},
			Flags:           app.Flags,
			HideHelpCommand: true,
			// Action:      A.AConfig,
			// Description: A.AConfigDescription,
			Subcommands: []*cli.Command{
				{
					Name:    "yaml",
					Usage:   "print config file in yaml format",
					Aliases: []string{"y"},
					Action:  A.SubAConfigYAML,
				},
				{
					Name:    "edit",
					Usage:   "edit the config file",
					Aliases: []string{"e"},
					Action:  A.SubAConfigEdit,
				},
			},
		},
		{
			Name:        "shortcuts",
			Usage:       "Show and edit shortcut names for workloads",
			Aliases:     []string{"sc"},
			Flags:       app.Flags,
			Description: T.Description(app, "shortcuts"),
			Subcommands: []*cli.Command{
				{
					Name:        "list",
					Aliases:     []string{"ls"},
					Usage:       "list all shortcuts",
					Action:      A.SubAShortcutList,
					Description: A.SubAShortcutListDescription,
				},
				{
					Name:        "add",
					Aliases:     []string{"a"},
					Usage:       "add new shortcut `sc add shortcut workload` ",
					Action:      A.SubAShortcutAdd,
					Description: A.SubAShortcutAddDescription,
				},
				{
					Name:        "remove",
					Aliases:     []string{"r"},
					Usage:       "remove a shortcut `sc remove shortcut` ",
					Action:      A.SubAShortcutRemove,
					Description: A.SubAShortcutRemoveDescription,
				},
			},
			After: func(c *cli.Context) error {
				A.SubAShortcutList(c)
				return nil
			},
		},
	}
}
