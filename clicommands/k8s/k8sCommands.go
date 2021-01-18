package k8s

import (
	"github.com/urfave/cli/v2"
)

// GetCLICommands ...
func GetCLICommands(app *cli.App) []*cli.Command {
	return []*cli.Command{
		{
			Name:        "certificates",
			Category:    "k8s",
			Usage:       "View and fetch certificates from cluster",
			Aliases:     []string{"cert"},
			Action:      ACertificate,
			Description: "Fetch certificates from remote location",
		},
		{
			Name:        "context",
			Category:    "k8s",
			Usage:       "View and switch current k8s context",
			Aliases:     []string{"ctx"},
			Action:      AContext,
			Description: "View and switch current k8s context",
		},
	}
}
