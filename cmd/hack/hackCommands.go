package hack

import (
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"

	"github.com/gregod-com/grgd/interfaces"
	"github.com/urfave/cli/v2"
)

func checkErr(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}

// GetCLICommands ...
func GetCLICommands(app *cli.App, core interfaces.ICore) []*cli.Command {
	h := core.GetHelper()
	hackFolder := core.GetConfig().GetActiveProfile().GetMetaData("hackDir")
	h.CheckOrCreateFolder(hackFolder, 0774)

	fileinfo, err := ioutil.ReadDir(hackFolder)
	checkErr(err)

	var scripts []*cli.Command
	// iterate over plugin implementations
	for _, f := range fileinfo {
		pluginPath := path.Join(hackFolder, f.Name())

		if strings.HasPrefix(f.Name(), ".") {
			continue
		}

		os.Chmod(pluginPath, 0744)

		name, err := h.CatchOutput(pluginPath, true, "name")
		if err != nil {
			continue
		}
		shortcuts, err := h.CatchOutput(pluginPath, true, "shortcuts")
		if err != nil {
			continue
		}
		description, err := h.CatchOutput(pluginPath, true, "description")
		if err != nil {
			continue
		}

		current := cli.Command{
			Name:        name,
			Category:    "plugins",
			Usage:       description,
			Aliases:     strings.Split(shortcuts, ","),
			Action:      AExec,
			Description: description,
		}

		scripts = append(scripts, &current)
	}
	return scripts
}
