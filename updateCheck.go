// This package implements the cli for the iam-stack. The underlying framework is depnedent upon urfave/cli.
package main

import (
	"github.com/gregod-com/grgdplugincontracts"
	cli "github.com/urfave/cli/v2"
)

// CheckUpdate ...
func CheckUpdate(c *cli.Context) error {
	// reader := bufio.NewReader(os.Stdin)
	repoIndex := c.App.Metadata["repoIndex"].(string)
	// currentIndex := c.App.Metadata[PLUGINSKEY].(string) + "/index.yaml"
	remoteIndex := c.App.Metadata[PLUGINSKEY].(string) + "/index-remote.yaml"
	UI := c.App.Metadata["UIPlugin"].(grgdplugincontracts.IUIPlugin)

	err := DownloadFile(remoteIndex, repoIndex)
	if err != nil {
		return err
	}
	// pluginsCurrent := PlugIndex.CreatePluginIndex(currentIndex)
	// pluginsRemote := PlugIndex.CreatePluginIndex(remoteIndex)

	UI.Println(c, "Downloaded: "+repoIndex)
	// for _, plremote := range pluginsRemote.GetPluginList().([]PlugIndex.PluginMetadata) {
	// 	for _, pllocal := range pluginsCurrent.GetPluginList().([]PlugIndex.PluginMetadata) {
	// 		if plremote.Name == pllocal.Name {
	// 			vlocal := semver.New(pllocal.Version)
	// 			vremote := semver.New(plremote.Version)

	// 			if vlocal.LessThan(*vremote) {
	// 				question := fmt.Sprintf("Update plugin    %-15s to v%v (current %v)? [y/n]", plremote.Name, vremote, vlocal)
	// 				if UI.YesNoQuestion(c, question) {
	// 					DownloadFile("filepath", plremote.URL)
	// 				}
	// 			}
	// 		}
	// 	}
	// }
	return nil
}
