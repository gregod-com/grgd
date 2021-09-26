package helper

import (
	"io/ioutil"
	"os"
	"path"
	"runtime"
	"sort"
	"strings"

	"github.com/gregod-com/grgd/interfaces"
	"golang.org/x/mod/semver"
	"gopkg.in/yaml.v2"
)

// ProvideNetworker ...
func ProvideUpdater(logger interfaces.ILogger, networker interfaces.INetworker) interfaces.IUpdater {
	updater := new(Updater)
	updater.logger = logger
	updater.networker = networker
	logger.Tracef("provide %T", updater)
	return updater
}

type Updater struct {
	logger    interfaces.ILogger
	networker interfaces.INetworker
}

// CheckUpdate ...
func (u *Updater) CheckUpdate(version string, core interfaces.ICore) error {
	cnfg := core.GetConfig()
	UI := core.GetUI()
	h := core.GetHelper()

	indexpath := path.Join(cnfg.GetActiveProfile().GetBasePath(), "index.yaml")
	versionMap := map[string]string{}

	// SCRIPTS
	hackFolder := path.Join(cnfg.GetActiveProfile().GetBasePath(), "hack")

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
		cName, err := h.CatchOutput(scriptPath, false, "name")
		if err != nil {
			return err
		}
		cVersion, err := h.CatchOutput(scriptPath, false, "version")
		if err != nil {
			return err
		}
		versionMap[cName] = cVersion
	}

	// err = u.networker.Load(indexpath, cnfg.GetActiveProfile().GetUpdateURL())
	// if err != nil {
	// 	return err
	// }

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
				err = u.networker.Load(path.Join(hackFolder, scriptname+"-partial"), script.Versions[semversion].URL)
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
		err = u.networker.Load("/usr/local/bin/grgd-partial", url)
		if err != nil {
			return err
		}
		os.Rename("/usr/local/bin/grgd-partial", "/usr/local/bin/grgd")
		os.Chmod("/usr/local/bin/grgd", 0744)
	}

	return nil
}

func (u *Updater) CheckSinceLastUpdate(version string, core interfaces.ICore) error {
	return nil
}
