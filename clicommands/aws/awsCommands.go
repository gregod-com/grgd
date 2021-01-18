package aws

import (
	"github.com/urfave/cli/v2"
)

// GetCLICommands ...
func GetCLICommands(app *cli.App) []*cli.Command {
	return []*cli.Command{
		{
			Name:        "dbconnector",
			Category:    "aws",
			Usage:       "Connect to remote database via EC2 instances",
			Aliases:     []string{"dbcon"},
			Action:      ADBConnector,
			Description: "Connect to remote database via EC2 instances",
		},
		{
			Name:        "nodeconnector",
			Category:    "aws",
			Usage:       "Connect to remote EC2 instances",
			Aliases:     []string{"ncon"},
			Action:      ANodeConnector,
			Description: "Connect to remote EC2 instances",
		},
	}
}
