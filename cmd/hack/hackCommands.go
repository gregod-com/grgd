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
	helper := core.GetHelper()

	hackFolder := helper.HomeDir(".grgd", "hack")
	helper.CheckOrCreateFolder(hackFolder, 0774)

	fileinfo, err := ioutil.ReadDir(hackFolder)
	checkErr(err)

	var scripts []*cli.Command
	// iterate over plugin implementations
	for _, f := range fileinfo {
		pluginPath := path.Join(hackFolder, f.Name())

		if strings.HasPrefix(pluginPath, ".") {
			continue
		}

		os.Chmod(pluginPath, 0744)

		name, err := catchOutput(pluginPath, true, "name")
		checkErr(err)
		shortcuts, err := catchOutput(pluginPath, true, "shortcuts")
		checkErr(err)
		description, err := catchOutput(pluginPath, true, "description")
		checkErr(err)

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
