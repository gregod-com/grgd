package flags

import (
	"github.com/urfave/cli/v2"
)

// GetFlags ...
func GetFlags(app *cli.App) []cli.Flag {
	return []cli.Flag{
		&cli.BoolFlag{
			Name:    "debug",
			Aliases: []string{"d"},
			Usage:   "run the command in debug mode",
		},
		&cli.BoolFlag{
			Name:    "verbose",
			Aliases: []string{"vvv"},
			Usage:   "run the command in verbose mode",
		},
		&cli.BoolFlag{
			Name:    "silent",
			Aliases: []string{"s"},
			Usage:   "mute all outputs",
		},
	}
}
