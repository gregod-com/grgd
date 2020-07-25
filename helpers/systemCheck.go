package helpers

import (
	"os"

	"github.com/urfave/cli/v2"
)

// SystemCheck ...
func SystemCheck(ctx *cli.Context) error {
	path := HomeDir() + "/.grgd"
	if _, notexistserr := os.Stat(path); os.IsNotExist(notexistserr) {
		os.Mkdir(path, os.FileMode(uint32(0760)))
	}
	ctx.App.Metadata["grgdhome"] = path

	pluginPath := HomeDir() + "/.grgd/plugins"
	if _, notexistserr := os.Stat(pluginPath); os.IsNotExist(notexistserr) {
		os.Mkdir(pluginPath, os.FileMode(uint32(0760)))
	}
	// ctx.App.Metadata["grgdplugins"] = pluginPath
	// updateCheckInterval := ctx.App.Metadata["updatecheckinterval"].(time.Duration)
	// pl := pluginindex.CreatePluginIndex(pluginPath + "/index.yaml")
	// if time.Now().After(pl.GetLastChecked().Add(updateCheckInterval)) {
	// 	err := CheckUpdate(ctx)
	// 	if err != nil {
	// 		log.Println("Looks like there was an error fetching updates... skipping this update-cycle")
	// 		// pl.Update()
	// 	}
	// }

	return nil
}
