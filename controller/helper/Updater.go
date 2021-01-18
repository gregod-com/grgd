package helper

import (
	"github.com/gregod-com/grgd/interfaces"
)

// ProvideUpdater ...
func ProvideUpdater(logger interfaces.ILogger) interfaces.IUpdater {
	up := new(Updater)
	up.logger = logger
	return up
}

// Updater ...
type Updater struct {
	logger interfaces.ILogger
}

// CheckUpdate ...
func (h *Updater) CheckUpdate(core interfaces.ICore) error {
	// UI := core.GetUI()
	var downloader interfaces.IDownloader
	err := core.Get(&downloader)
	if err != nil {
		return err
	}

	err = downloader.Load("/usr/local/bin/grgd", "https://s3.iamstudent.dev/public/grgd/grgd-darwin")
	if err != nil {
		return err
	}

	return nil
}
