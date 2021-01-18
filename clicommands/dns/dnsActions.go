package dns

import (
	"github.com/gregod-com/grgd/controller/helper"
	"github.com/urfave/cli/v2"
)

// AEnter ...
func AEnter(c *cli.Context) error {
	core := helper.GetExtractor().GetCore(c)
	UI := core.GetUI()
	UI.Println("Hellos from dns command")

	return nil
}
