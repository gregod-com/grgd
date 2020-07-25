// This package implements the cli for the iam-stack. The underlying framework is depnedent upon urfave/cli.
package helpers

import (
	"github.com/gregod-com/grgdplugincontracts"
	"github.com/urfave/cli/v2"
)

// CheckUpdate ...
func CheckUpdate(c *cli.Context) error {
	// reader := bufio.NewReader(os.Stdin)
	var UI grgdplugincontracts.IUIPlugin
	var repoIndex, remoteIndex string
	ExtractMetadataFatal(c.App.Metadata, "pluginIndex", &repoIndex)
	ExtractMetadataFatal(c.App.Metadata, "remoteIndex", &remoteIndex)
	ExtractMetadataFatal(c.App.Metadata, "UIPlugin", &UI)

	// repoIndex, ok := c.App.Metadata["repoIndex"].(string)
	// if !ok {
	// 	return errors.New("Metadata map is missing key `repoIndex`")
	// }
	// currentIndex := c.App.Metadata[PLUGINSKEY].(string) + "/index.yaml"
	// remoteIndex := c.App.Metadata["grgdplugins"].(string) + "/index-remote.yaml"

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
