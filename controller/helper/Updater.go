package helper

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path"
	"runtime"
	"sort"
	"strings"
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
	versionMap := map[string]string{}

	// SCRIPTS
	hackFolder := path.Join(cnfg.GetProfile().GetBasePath(), "hack")

	fileinfo, err := ioutil.ReadDir(hackFolder)
	if err != nil {
		return err
	}

	// iterate over grgd hack folder
	for _, f := range fileinfo {
		scriptPath := path.Join(hackFolder, f.Name())
		if strings.HasPrefix(scriptPath, ".") {
			continue
		}
		cName := catchOut(scriptPath, "name")
		cVersion := catchOut(scriptPath, "version")
		versionMap[cName] = cVersion
	}

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
		sn := scriptname
		didUpdate := false
		sortedScriptVersions := []string{}
		for semversion := range script.Versions {
			sortedScriptVersions = append(sortedScriptVersions, semversion)
		}
		sortSemverSlice(sortedScriptVersions)

		for _, semversion := range sortedScriptVersions {
			remark := ""
			switch semver.Compare(versionMap[scriptname], semversion) {
			case -1:
				remark = "(newer)"
			case 0:
				remark = "(current)"
			case 1:
				remark = "(older)"
			}
			UI.Printf("%-20v %-10v %v -> %v\n", sn, remark, semversion, script.Versions[semversion].URL)
			if didUpdate == false && remark == "(newer)" && UI.YesNoQuestionf("Would you like to update script `%v` from %v to %v now?",
				scriptname,
				versionMap[scriptname],
				semversion) {
				UI.Println("DOWNLOADING!!!!")
				err = downloader.Load(path.Join(hackFolder, scriptname+"-partial"), script.Versions[semversion].URL)
				if err != nil {
					return err
				}
				os.Rename(path.Join(hackFolder, scriptname+"-partial"), path.Join(hackFolder, scriptname))
				os.Chmod(path.Join(hackFolder, scriptname), 0744)
				didUpdate = true
			}
			sn = ""
		}
		UI.Println()
	}

	UI.Println("----GRGD-CLI-------")
	newerVersions := []string{}

	sortedversions := []string{}
	for k := range indexObject.Releases["grgd-cli"].Targets[runtime.GOOS+"-"+runtime.GOARCH].Versions {
		sortedversions = append(sortedversions, k)
	}
	sortSemverSlice(sortedversions)
	grgd := "grgd"
	for _, k := range sortedversions {
		v := indexObject.Releases["grgd-cli"].Targets[runtime.GOOS+"-"+runtime.GOARCH].Versions[k]
		if len(os.Args) > 2 {
			if os.Args[2] == k {
				newerVersions = append(newerVersions, k)
			}
		}
		remark := ""
		switch semver.Compare("v"+version, k) {
		case -1:
			newerVersions = append(newerVersions, k)
			remark = "(newer)"
			// UI.Printf("%v (new)     -> %v \n", k, v.URL)
		case 0:
			remark = "(current)"
			// UI.Printf("%v (current) -> %v \n", k, v.URL)
		case 1:
			remark = "(older)"
			// UI.Printf("%v (old)     -> %v \n", k, v.URL)
		}
		UI.Printf("%-20v %-10v %v -> %v\n", grgd, remark, k, v.URL)
		grgd = ""
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

func catchOut(binPath string, args ...string) string {
	cmd := exec.Command(binPath, args...)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal("Error executing: " + err.Error())
	}
	return strings.TrimSuffix(out.String(), "\n")
}

func sortSemverSlice(semverSlice []string) error {
	// sort.Strings(sortedScriptVersions)
	sort.Slice(semverSlice, func(i, j int) bool {
		switch semver.Compare(semverSlice[i], semverSlice[j]) {
		case 1:
			return true
		default:
			return false
		}
	})
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
