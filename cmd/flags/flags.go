package flags

import (
	"github.com/urfave/cli/v2"
)

// GetFlags ...
func GetFlags() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:    "profile",
			Aliases: []string{"p"},
			EnvVars: []string{"USER"},
		},
		&cli.StringFlag{
			Name:  "log-level",
			Value: "info",
		},
	}
}
