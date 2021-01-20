package stack

import (
	"github.com/gregod-com/grgd/interfaces"
	"github.com/urfave/cli/v2"
)

// GetCLICommands ...
func GetCLICommands(app *cli.App, core interfaces.ICore) []*cli.Command {
	return []*cli.Command{
		{
			Name:        "up",
			Category:    "local stack",
			Usage:       "Start dev stack with current config",
			Aliases:     []string{"u"},
			Action:      AUp,
			Description: "up description",
		},
		{
			Name:        "down",
			Category:    "local stack",
			Usage:       "Stop dev stack",
			Aliases:     []string{"d"},
			Action:      ADown,
			Description: "down description",
		},
		{
			Name:        "restart",
			Category:    "stack",
			Usage:       "Restart dev stack (or single workload)",
			Aliases:     []string{"r"},
			Action:      ARestart,
			Description: "restart description",
		},
		{
			Name:        "logs",
			Category:    "stack",
			Usage:       "Show all logs for running stack (or single workload)",
			Aliases:     []string{"l"},
			ArgsUsage:   "Args usage",
			Action:      ALogs,
			Description: "logs description",
		},
		{
			Name:      "activate",
			Category:  "stack",
			Usage:     "Activate a workload",
			Aliases:   []string{"act", "a"},
			ArgsUsage: "TODO",
			Action:    AActivate,
			After: func(c *cli.Context) error {
				// UIPlugin.PrintWorkloadOverview(c)
				return nil
			},
			Description: "activate description",
		},
		{
			Name:        "enter",
			Category:    "stack",
			Usage:       "Enter a workload",
			Aliases:     []string{"en"},
			Action:      AActivate,
			Description: "enter description",
		},
		{
			Name:        "execute",
			Category:    "stack",
			Usage:       "Execute a command in workload and view output",
			Aliases:     []string{"exec", "ex"},
			Action:      AActivate,
			Description: "execute description",
		},
		{
			Name:     "test",
			Category: "stack",
			Usage:    "Run unittest inside container",
			Aliases:  []string{"t"},
			// Flags:       myFlags,
			ArgsUsage:   "TODO",
			Action:      AActivate,
			Description: "test description",
		},
	}
}
