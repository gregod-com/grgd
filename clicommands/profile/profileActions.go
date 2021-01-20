package profile

import (
	"github.com/gregod-com/grgd/controller/helper"
	"github.com/urfave/cli/v2"
)

// AListProfiles ...
func AListProfiles(c *cli.Context) error {
	core := helper.GetExtractor().GetCore(c)
	UI := core.GetUI()
	UI.Println("This is the profile command...")
	return nil
}
