package shortcut

import (
	"github.com/urfave/cli/v2"
)

// GetCLICommands ...
func GetCLICommands(app *cli.App) []*cli.Command {
	return []*cli.Command{
		{
			Name:    "shortcuts",
			Usage:   "Show and edit shortcut names for workloads",
			Aliases: []string{"sc"},
			Flags:   app.Flags,
			Subcommands: []*cli.Command{
				{
					Name:        "list",
					Aliases:     []string{"ls"},
					Usage:       "list all shortcuts",
					Action:      SubAShortcutList,
					Description: SubAShortcutListDescription,
				},
				{
					Name:        "add",
					Aliases:     []string{"a"},
					Usage:       "add new shortcut `sc add shortcut workload` ",
					Action:      SubAShortcutAdd,
					Description: SubAShortcutAddDescription,
				},
				{
					Name:        "remove",
					Aliases:     []string{"r"},
					Usage:       "remove a shortcut `sc remove shortcut` ",
					Action:      SubAShortcutRemove,
					Description: SubAShortcutRemoveDescription,
				},
			},
			After: SubAShortcutList,
		},
	}
}
