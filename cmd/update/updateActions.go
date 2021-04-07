package update

import (
	"github.com/gregod-com/grgd/pkg/helper"
	"github.com/urfave/cli/v2"
)

// AUpdate ...
func AUpdate(c *cli.Context) error {
	core := helper.GetExtractor().GetCore(c)
	logger := core.GetLogger()
	networker := core.GetNetworker()

	err := networker.CheckUpdate(c.App.Version, core)
	if err != nil {
		logger.Fatal(err)
	}
	return nil
}
