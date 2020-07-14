package actions

import (
	"fmt"

	cli "github.com/urfave/cli/v2"
)

// APluginList ...
func APluginList(c *cli.Context) error {
	// for k, v := range c.App.Metadata["AllPlugins"].([]I.IGrgdPlugin) {
	// fmt.Printf("index: %v -> %v -> %v\n", k, v.Name(), v.)
	// }
	return nil
}

// APluginActivate ...
func APluginActivate(c *cli.Context) error {
	go fmt.Println("this is the plugin activate command")
	return nil
}
