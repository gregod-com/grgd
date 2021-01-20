package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path"
	"strings"
	"time"

	"github.com/gregod-com/grgd/controller/helper"
	"gopkg.in/yaml.v2"
)

func main() {
	if len(os.Args) < 4 {
		os.Exit(1)
	}
	bin := os.Args[1]
	system := os.Args[2]
	platform := os.Args[3]

	binaryRaw := path.Join("bin", bin+"-"+system+"-"+platform)

	version := catchOut(binaryRaw, "-v")
	versionParts := strings.Split(version, " ")
	version = "v" + versionParts[2]
	md5 := catchOut("md5", "-q", binaryRaw)

	file, err := os.Open(binaryRaw)
	if err != nil {
		log.Fatal(err)
	}

	filestat, err := file.Stat()
	if err != nil {
		log.Fatal(err)
	}

	metadata := helper.DownloadMetadata{
		URL:          "https://s3.iamstudent.dev/public/grgd/" + bin + "-" + system + "-" + platform + "-" + version,
		Description:  "macos grgd cli blabla",
		Released:     time.Now(),
		ReleaseNotes: "change XYZ",
		Author:       "gregod",
		Size:         int(filestat.Size()),
		Md5:          md5,
	}

	file.Close()
	catchOut("cp", binaryRaw, binaryRaw+"-"+version)

	index, err1 := ioutil.ReadFile(path.Join("bin", "index.yaml"))
	if err1 != nil {
		log.Fatal(err1)
	}

	indexObject := &helper.IndexObject{}
	if err2 := yaml.Unmarshal(index, indexObject); err2 != nil {
		log.Fatal(err2)
	}

	indexObject.Releases["grgd-cli"].Targets[system+"-"+platform].Versions[version] = metadata

	// SCRIPTS
	home, _ := os.UserHomeDir()
	hackFolder := path.Join(home, ".grgd", "hack")

	fileinfo, err := ioutil.ReadDir(hackFolder)
	if err != nil {
		log.Fatal(err)
	}

	// iterate over grgd hack folder
	for _, f := range fileinfo {
		scriptPath := path.Join(hackFolder, f.Name())

		if strings.HasPrefix(scriptPath, ".") {
			continue
		}

		cName := catchOut(scriptPath, "name")
		cVersion := catchOut(scriptPath, "version")
		cDescription := catchOut(scriptPath, "description")
		cMD5 := catchOut("md5", "-q", scriptPath)

		cmetadata := helper.DownloadMetadata{
			URL:          "https://s3.iamstudent.dev/public/grgd/hack/" + cName + "-hack-" + cVersion,
			Description:  cDescription,
			Released:     time.Now(),
			ReleaseNotes: "change XYZ",
			Author:       "gregod",
			Size:         int(f.Size()),
			Md5:          cMD5,
		}

		indexObject.Releases["grgd-hacks"].Targets[cName].Versions[cVersion] = cmetadata
		catchOut("cp", scriptPath, path.Join("bin", "hack", cName+"-hack-"+cVersion))
	}

	outbyte, err2 := yaml.Marshal(indexObject)
	if err2 != nil {
		log.Fatal(err2)
	}

	ioutil.WriteFile(path.Join("bin", "index.yaml"), outbyte, 0744)
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
