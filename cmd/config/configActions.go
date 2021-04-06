package config

import (
	"github.com/gregod-com/grgd/pkg/helper"
	"github.com/urfave/cli/v2"
)

// SubAConfigYAML ...
func SubAConfigYAML(c *cli.Context) error {
	core := helper.GetExtractor().GetCore(c)
	UI := core.GetUI()
	config := core.GetConfig()
	UI.Println(config)
	return nil
}

// SubAConfigJSON ...
func SubAConfigJSON(c *cli.Context) error {
	core := helper.GetExtractor().GetCore(c)
	UI := core.GetUI()
	config := core.GetConfig()
	UI.Println(config.DumpConfig("json"))
	return nil
}

// SubAConfigEdit ...
func SubAConfigEdit(c *cli.Context) error {
	return nil
}
