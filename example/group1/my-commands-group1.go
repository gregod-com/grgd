package group1

import (
	"fmt"

	"github.com/gregod-com/grgd/pkg/helper"

	cli "github.com/urfave/cli/v2"
)

type CMD struct {
}

func (cmd *CMD) GetCommands(i interface{}) interface{} {
	app, ok := i.(*cli.App)
	if !ok {
		return fmt.Errorf("did not pass *cli.App")
	}
	return []*cli.Command{
		{
			Name:            "my-commands-group-1",
			Category:        "group-1",
			Usage:           "do some stuff",
			HideHelpCommand: true,
			Before:          nil,
			Flags:           app.Flags,
			Subcommands: []*cli.Command{
				{
					Name:    "subcommand-1",
					Usage:   "do some specific stuff",
					Aliases: []string{"sc-1"},
					Flags:   app.Flags,
					Action:  MyCommandAction1,
				},
				{
					Name:    "subcommand-2",
					Usage:   "do some other specific stuff",
					Aliases: []string{"sc-2"},
					Flags:   app.Flags,
					Action:  MyCommandAction2,
				},
			},
		},
	}
}

// MyCommandAction1 ...
func MyCommandAction1(c *cli.Context) error {
	core := helper.GetExtractor().GetCore(c)
	log := core.GetLogger()
	ui := core.GetUI()
	log.Warnf("This is good: %T", c)
	ui.PrintPercentOfScreen(10, 90, "/")
	return nil
}

// MyCommandAction2 ...
func MyCommandAction2(c *cli.Context) error {
	return nil
}
