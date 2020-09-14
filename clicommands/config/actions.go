package config

import (
	"grgd/controller/helper"

	"github.com/urfave/cli/v2"
)

// SubAConfigYAML ...
func SubAConfigYAML(c *cli.Context) error {
	ext := helper.GetExtractor()
	UI := ext.GetCore(c).GetUI()
	configObject := ext.GetCore(c).GetConfig()
	UI.Println(c, configObject.DumpConfig("yaml"))
	return nil
}

// SubAConfigJSON ...
func SubAConfigJSON(c *cli.Context) error {
	ext := helper.GetExtractor()
	UI := ext.GetCore(c).GetUI()
	configObject := ext.GetCore(c).GetConfig()
	UI.Println(c, configObject.DumpConfig("json"))
	return nil
}

// SubAConfigEdit ...
func SubAConfigEdit(c *cli.Context) error {
	// ext := helper.GetExtractor()
	// UI := ext.GetCore(c).GetUI()
	// configObject := ext.GetCore(c).GetConfig()
	// UI.Println(c, configObject)
	return nil
}
