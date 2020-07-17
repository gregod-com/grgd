// This package implements the cli for the iam-stack. The underlying framework is depnedent upon urfave/cli.
package main

import (
	"log"
	"os"
	"time"

	PlugIndex "github.com/gregod-com/grgd/pluginindex"
	"github.com/gregod-com/grgdplugins/shared"

	cli "github.com/urfave/cli/v2"
)

// SystemCheck ...
func SystemCheck(ctx *cli.Context) error {
	path := shared.HomeDir() + "/.grgd"
	if _, notexistserr := os.Stat(path); os.IsNotExist(notexistserr) {
		os.Mkdir(path, os.FileMode(uint32(0760)))
	}
	ctx.App.Metadata["grgdhome"] = path

	pluginPath := shared.HomeDir() + "/.grgd/plugins"
	if _, notexistserr := os.Stat(pluginPath); os.IsNotExist(notexistserr) {
		os.Mkdir(pluginPath, os.FileMode(uint32(0760)))
	}
	ctx.App.Metadata[PLUGINSKEY] = pluginPath
	updateCheckInterval := ctx.App.Metadata["updatecheckinterval"].(time.Duration)
	pl := PlugIndex.CreatePluginIndex(pluginPath + "/index.yaml")
	if time.Now().After(pl.GetLastChecked().Add(updateCheckInterval)) {
		err := CheckUpdate(ctx)
		if err != nil {
			log.Println("Looks like there was an error fetching updates... skipping this update-cycle")
			pl.Update()
		}
	}

	return nil
}
