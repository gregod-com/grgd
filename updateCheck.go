// This package implements the cli for the iam-stack. The underlying framework is depnedent upon urfave/cli.
package main

import (
	cli "github.com/urfave/cli/v2"
)

// CheckUpdate ...
func CheckUpdate(ctx *cli.Context) error {
	// reader := bufio.NewReader(os.Stdin)
	repoIndex := ctx.App.Metadata["repoIndex"].(string)
	// currentIndex := ctx.App.Metadata[PLUGINSKEY].(string) + "/index.yaml"
	remoteIndex := ctx.App.Metadata[PLUGINSKEY].(string) + "/index-remote.yaml"
	err := DownloadFile(remoteIndex, repoIndex)
	if err != nil {
		return err
	}
	// pluginsCurrent := PlugIndex.CreatePluginIndex(currentIndex)
	// pluginsRemote := PlugIndex.CreatePluginIndex(remoteIndex)

	// fmt.Println("Downloaded: " + repoIndex)
	// for _, plremote := range pluginsRemote.GetPluginList().([]PlugIndex.PluginMetadata) {
	// 	for _, pllocal := range pluginsCurrent.GetPluginList().([]PlugIndex.PluginMetadata) {
	// 		if plremote.Name == pllocal.Name {
	// 			vlocal := semver.New(pllocal.Version)
	// 			vremote := semver.New(plremote.Version)
	// 			if vlocal.LessThan(*vremote) {
	// 				fmt.Printf("Update plugin    %-15s to v%v (current %v)? [y/n]", plremote.Name, vremote, vlocal)
	// 				yes, _ := reader.ReadString('\n')
	// 				if yes == "y\n" {
	// 					fmt.Println(plremote.Sha)
	// 					DownloadFile("filepath", plremote.URL)
	// 				}
	// 			}
	// 		}
	// 	}

	// 	// fmt.Println(p.Name)
	// 	// fmt.Println(p.Version)
	// 	// fmt.Println(p.Size)
	// 	// fmt.Println(p.Sha)
	// 	// fmt.Println(p.URL)
	// }
	// pl.Update()
	return nil
}
