package update

import (
	"github.com/gregod-com/grgd/pkg/helper"
	"github.com/urfave/cli/v2"
)

// AUpdate ...
func AUpdate(c *cli.Context) error {
	core := helper.GetExtractor().GetCore(c)
	logger := core.GetLogger()
	updater := core.GetUpdater()

	err := updater.CheckUpdate(c.App.Version, core)
	if err != nil {
		logger.Fatal(err)
	}
	return nil
}
