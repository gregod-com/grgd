package actions

import (
	"github.com/gregod-com/grgdplugincontracts"
	"github.com/urfave/cli/v2"
)

// AInit ...
func AInit(c *cli.Context) error {
	UI := c.App.Metadata["UIPlugin"].(grgdplugincontracts.IUIPlugin)
	UI.Println(c, "this is the init command")

	return nil
}
