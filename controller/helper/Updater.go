package helper

import (
	"grgd/interfaces"
)

// ProvideUpdater ...
func ProvideUpdater(logger interfaces.ILogger) interfaces.IUpdater {
	up := &Updater{logger: logger}
	return up
}

// Updater ...
type Updater struct {
	logger interfaces.ILogger
}

// CheckUpdate ...
func (h *Updater) CheckUpdate(core interfaces.ICore) error {
	UI := core.GetUI()
	downloader := core.GetDownloader()
	// core.GetConfig().GetPlu
	// ext.GetMetadataFatal(c.App.Metadata, "pluginIndex", &repoIndex)
	// ext.GetMetadataFatal(c.App.Metadata, "remoteIndex", &remoteIndex)
	// ext.GetMetadataFatal(c.App.Metadata, "UIPlugin", &UI)

	err := downloader.Load("file_location", "repo_url")
	if err != nil {
		return err
	}

	// pluginsCurrent := PlugIndex.CreatePluginIndex(currentIndex)
	// pluginsRemote := PlugIndex.CreatePluginIndex(remoteIndex)

	UI.Println(nil, "Downloaded: ")
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
