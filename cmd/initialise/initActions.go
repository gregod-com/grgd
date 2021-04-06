package initialise

import (
	"github.com/gregod-com/grgd/pkg/helper"

	"github.com/urfave/cli/v2"
)

// AInit ...
func AInit(c *cli.Context) error {
	UI := helper.GetExtractor().GetCore(c).GetUI()
	UI.Println("this is the init command")
	return nil
}
