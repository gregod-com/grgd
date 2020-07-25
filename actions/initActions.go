package actions

import (
	"github.com/gregod-com/grgd/helpers"
	"github.com/gregod-com/grgdplugincontracts"
	"github.com/urfave/cli/v2"
)

// AInit ...
func AInit(c *cli.Context) error {
	var UI grgdplugincontracts.IUIPlugin
	helpers.ExtractMetadataFatal(c.App.Metadata, "UIPlugin", &UI)
	UI.Println(c, "this is the init command")
	return nil
}
