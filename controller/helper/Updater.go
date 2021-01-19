package helper

import (
	"io/ioutil"
	"os"
	"path"
	"runtime"
	"sort"
	"time"

	"github.com/gregod-com/grgd/interfaces"
	"golang.org/x/mod/semver"
	"gopkg.in/yaml.v2"
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
func (h *Updater) CheckUpdate(version string, core interfaces.ICore) error {
	// UI := core.GetUI()
	var downloader interfaces.IDownloader
	err := core.Get(&downloader)
	if err != nil {
		return err
	}
	cnfg := core.GetConfig()
	UI := core.GetUI()

	indexpath := path.Join(cnfg.GetProfile().GetBasePath(), "index.yaml")

	err = downloader.Load(indexpath, cnfg.GetProfile().GetUpdateURL())
	if err != nil {
		return err
	}

	index, err := ioutil.ReadFile(indexpath)
	if err != nil {
		return err
	}

	indexObject := &IndexObject{}
	if err2 := yaml.Unmarshal(index, indexObject); err2 != nil {
		return err2
	}

	UI.Println("----SCRIPTS/HACK----")
	for scriptname, script := range indexObject.Releases["grgd-hacks"].Targets {
		for semversion, dlMeta := range script.Versions {
			UI.Printf("%-20v %v -> %v\n", scriptname, semversion, dlMeta.URL)
			scriptname = ""
		}
	}
	UI.Println()

	UI.Println("----GRGD-CLI-------")
	newerVersions := []string{}

	sortedversions := []string{}
	for k := range indexObject.Releases["grgd-cli"].Targets[runtime.GOOS+"-"+runtime.GOARCH].Versions {
		sortedversions = append(sortedversions, k)
	}
	sort.Strings(sortedversions)

	for _, k := range sortedversions {
		v := indexObject.Releases["grgd-cli"].Targets[runtime.GOOS+"-"+runtime.GOARCH].Versions[k]
		if len(os.Args) > 2 {
			if os.Args[2] == k {
				newerVersions = append(newerVersions, k)
			}
		}
		switch semver.Compare("v"+version, k) {
		case -1:
			newerVersions = append(newerVersions, k)
			UI.Printf("%v (new)     -> %v \n", k, v.URL)
		case 0:
			UI.Printf("%v (current) -> %v \n", k, v.URL)
		case 1:
			UI.Printf("%v (old)     -> %v \n", k, v.URL)
		}
	}

	UI.Println("-------------------")

	sort.Strings(newerVersions)
	if len(newerVersions) > 0 && UI.YesNoQuestionf("Do you want to update to version %v now?", newerVersions[0]) {
		url := indexObject.Releases["grgd-cli"].Targets[runtime.GOOS+"-"+runtime.GOARCH].Versions[newerVersions[0]].URL
		err = downloader.Load("/usr/local/bin/grgd-partial", url)
		if err != nil {
			return err
		}
		os.Rename("/usr/local/bin/grgd-partial", "/usr/local/bin/grgd")
		os.Chmod("/usr/local/bin/grgd", 0744)
	}

	return nil
}

// IndexObject ...
type IndexObject struct {
	Releases map[string]Category `yaml:"category"`
}

// Category ...
type Category struct {
	Targets map[string]Target `yaml:"target"`
}

type Target struct {
	Versions map[string]DownloadMetadata `yaml:"version"`
}

// Version ...
// type Version struct {
// 	DownloadMetadata map[string]DownloadMetadata
// }

// DownloadMetadata ...
type DownloadMetadata struct {
	Author       string
	Description  string
	Md5          string
	Released     time.Time
	Size         int
	URL          string
	ReleaseNotes string
}
