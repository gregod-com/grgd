package helper

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path"
	"runtime"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/gregod-com/grgd/interfaces"
	"golang.org/x/mod/semver"
	"gopkg.in/yaml.v2"
)

// ProvideNetworker ...
func ProvideNetworker(logger interfaces.ILogger) interfaces.INetworker {
	networker := new(Networker)
	networker.logger = logger
	logger.Tracef("provide %T", networker)
	return networker
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

// Connection ...
type Connection struct {
	Endpoint string
	TimeOut  int
	Success  bool
}

// Networker ...
type Networker struct {
	logger interfaces.ILogger
}

// CheckUpdate ...
func (n *Networker) CheckUpdate(version string, core interfaces.ICore) error {
	cnfg := core.GetConfig()
	UI := core.GetUI()

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
		cName := catchOut(scriptPath, "name")
		cVersion := catchOut(scriptPath, "version")
		versionMap[cName] = cVersion
	}

	err = n.Load(indexpath, cnfg.GetActiveProfile().GetUpdateURL())
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
				err = n.Load(path.Join(hackFolder, scriptname+"-partial"), script.Versions[semversion].URL)
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
		err = n.Load("/usr/local/bin/grgd-partial", url)
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

var downloadSize uint64

// Load ...
func (n *Networker) Load(filepath string, url string) error {

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	size, _ := strconv.Atoi(resp.Header.Get("Content-Length"))
	downloadSize = uint64(size)

	counter := &WriteCounter{}
	// Write the body to file
	_, err = io.Copy(out, io.TeeReader(resp.Body, counter))
	fmt.Println()
	return err
}

// WriteCounter ...
type WriteCounter struct {
	Total uint64
}

func (wc *WriteCounter) Write(p []byte) (int, error) {
	n := len(p)
	wc.Total += uint64(n)
	wc.PrintProgress()
	return n, nil
}

// PrintProgress ...
func (wc WriteCounter) PrintProgress() {
	// Clear the line by using a character return to go back to the start and remove
	// the remaining characters by filling it with spaces
	fmt.Printf("\r%s", strings.Repeat(" ", 35))

	// Return again and print current status of download
	// We use the humanize package to print the bytes in a meaningful way (e.g. 10 MB)
	fmt.Printf("\rDownloading... %d (%d MB) (%d%%) ", wc.Total/1024/1024, downloadSize/1024/1024, (wc.Total * 100 / downloadSize))
}

// CheckConnections ...
func (n *Networker) CheckConnections(conns map[string]interface{}) {
	for k := range conns {
		if conn, ok := conns[k].(Connection); ok {
			_, err := http.Get(conn.Endpoint)
			if err != nil {
				log.Fatal(err)
			}
			conn.Success = true
		}
	}
	return
}
