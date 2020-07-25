package actions

import (
	"github.com/gregod-com/grgd/helpers"
	"github.com/gregod-com/grgdplugincontracts"
	"github.com/gregod-com/interfaces"
	"github.com/urfave/cli/v2"
)

// AConfigDescription ...
const AConfigDescription = `This is the description as a var in single quotes multi line `

// AConfig ...
func AConfig(c *cli.Context) error {
	var UI grgdplugincontracts.IUIPlugin
	helpers.ExtractMetadataFatal(c.App.Metadata, "UIPlugin", &UI)

	UI.Println(c, "this is the config command")
	return nil
}

// SubAConfigYAMLDescription ...
const SubAConfigYAMLDescription = `This is the description as a var in single quotes`

// SubAConfigYAML ...
func SubAConfigYAML(c *cli.Context) error {
	var UI grgdplugincontracts.IUIPlugin
	var configObject interfaces.IConfigObject
	helpers.ExtractMetadataFatal(c.App.Metadata, "UIPlugin", &UI)
	helpers.ExtractMetadataFatal(c.App.Metadata, "config", &configObject)

	UI.Println(c, configObject.GetSourceAsString())
	return nil
}

// SubAConfigEdit ...
func SubAConfigEdit(c *cli.Context) error {
	var UI grgdplugincontracts.IUIPlugin
	helpers.ExtractMetadataFatal(c.App.Metadata, "UIPlugin", &UI)

	UI.Println(c, "This is the config edit subcommand")
	return nil
}
