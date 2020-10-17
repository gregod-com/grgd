package config

import (
	"grgd/controller/helper"

	"github.com/urfave/cli/v2"
)

// SubAConfigYAML ...
func SubAConfigYAML(c *cli.Context) error {
	ext := helper.GetExtractor()
	core := ext.GetCore(c)
	UI := core.GetUI()
	configObject := core.GetConfig()
	UI.Println(configObject.DumpConfig("yaml"), c)
	return nil
}

// SubAConfigJSON ...
func SubAConfigJSON(c *cli.Context) error {
	ext := helper.GetExtractor()
	UI := ext.GetCore(c).GetUI()
	configObject := ext.GetCore(c).GetConfig()
	UI.Println(configObject.DumpConfig("json"), c)
	return nil
}

// SubAConfigEdit ...
func SubAConfigEdit(c *cli.Context) error {
	// ext := helper.GetExtractor()
	// UI := ext.GetCore(c).GetUI()
	// configObject := ext.GetCore(c).GetConfig()
	// UI.Println(configObject,c)
	return nil
}
