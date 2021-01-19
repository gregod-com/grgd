package update

import (
	"github.com/gregod-com/grgd/controller/helper"
	"github.com/gregod-com/grgd/interfaces"
	"github.com/urfave/cli/v2"
)

// AUpdate ...
func AUpdate(c *cli.Context) error {
	core := helper.GetExtractor().GetCore(c)
	logger := core.GetLogger()
	var updater interfaces.IUpdater
	core.Get(&updater)

	err := updater.CheckUpdate(c.App.Version, core)
	if err != nil {
		logger.Fatal(err)
	}
	return nil
}
