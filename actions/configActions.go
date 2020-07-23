package actions

import (
	"github.com/gregod-com/grgdplugincontracts"
	I "github.com/gregod-com/interfaces"
	"github.com/urfave/cli/v2"
)

// AConfigDescription ...
const AConfigDescription = `This is the description as a var in single quotes multi line `

// AConfig ...
func AConfig(c *cli.Context) error {
	UI := c.App.Metadata["UIPlugin"].(grgdplugincontracts.IUIPlugin)
	UI.Println(c, "this is the config command")
	return nil
}

// SubAConfigYAMLDescription ...
const SubAConfigYAMLDescription = `This is the description as a var in single quotes`

// SubAConfigYAML ...
func SubAConfigYAML(c *cli.Context) error {
	UI := c.App.Metadata["UIPlugin"].(grgdplugincontracts.IUIPlugin)
	configObject := c.App.Metadata["config"].(I.IConfigObject)
	UI.Println(c, configObject.GetSourceAsString())
	return nil
}

// SubAConfigEdit ...
func SubAConfigEdit(c *cli.Context) error {
	UI := c.App.Metadata["UIPlugin"].(grgdplugincontracts.IUIPlugin)
	UI.Println(c, "This is the config edit subcommand")
	return nil
}
