package config

import (
	"grgd/controller/helper"

	"github.com/urfave/cli/v2"
)

// SubAConfigYAML ...
func SubAConfigYAML(c *cli.Context) error {
	core := helper.GetExtractor().GetCore(c)
	UI := core.GetUI()
	configObject := core.GetConfig()
	UI.Println(configObject.DumpConfig("yaml"))
	return nil
}

// SubAConfigJSON ...
func SubAConfigJSON(c *cli.Context) error {
	core := helper.GetExtractor().GetCore(c)
	UI := core.GetUI()
	configObject := core.GetConfig()
	UI.Println(configObject.DumpConfig("json"))
	return nil
}

// SubAConfigEdit ...
func SubAConfigEdit(c *cli.Context) error {
	return nil
}
