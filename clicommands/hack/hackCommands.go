package hack

import (
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"

	"github.com/urfave/cli/v2"
)

func checkErr(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}

// GetCLICommands ...
func GetCLICommands(app *cli.App) []*cli.Command {
	homedir, err := os.UserHomeDir()
	if err != nil {
		return nil
	}

	hackFolder := path.Join(homedir, ".grgd", "hack")

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

		name, err := catchOutput(pluginPath, "name")
		checkErr(err)
		shortcuts, err := catchOutput(pluginPath, "shortcuts")
		checkErr(err)
		description, err := catchOutput(pluginPath, "description")
		checkErr(err)

		current := cli.Command{
			Name:        name,
			Category:    "local hack",
			Usage:       description,
			Aliases:     strings.Split(shortcuts, ","),
			Action:      AExec,
			Description: description,
		}

		scripts = append(scripts, &current)
	}
	return scripts
}
