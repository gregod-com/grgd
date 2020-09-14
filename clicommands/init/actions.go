package init

import (
	"grgd/controller/helper"

	"github.com/urfave/cli/v2"
)

// AInit ...
func AInit(c *cli.Context) error {
	UI := helper.GetExtractor().GetCore(c).GetUI()
	UI.Println(c, "this is the init command")
	return nil
}
