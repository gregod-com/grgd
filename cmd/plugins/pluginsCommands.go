package plugins

import (
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"

	"github.com/gregod-com/grgd/interfaces"
	"github.com/urfave/cli/v2"
)

func warnOnError(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}

func failOnError(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}

// GetCLICommands ...
func GetCLICommands(app *cli.App, core interfaces.ICore) []*cli.Command {
	var fsm interfaces.IFileSystemManipulator
	core.Get(&fsm)

	pluginFolder := core.GetConfig().GetActiveProfile().GetPluginsDir()
	fsm.CheckOrCreateFolder(pluginFolder, 0774)

	fileinfo, err := ioutil.ReadDir(pluginFolder)
	failOnError(err)

	var scripts []*cli.Command
	// iterate over plugin implementations
	for _, f := range fileinfo {
		pluginPath := path.Join(pluginFolder, f.Name())

		if strings.HasPrefix(pluginPath, ".") {
			continue
		}

		os.Chmod(pluginPath, 0744)

		name, err := catchOutput(pluginPath, true, "name")
		failOnError(err)
		shortcuts, err := catchOutput(pluginPath, true, "shortcuts")
		failOnError(err)
		description, err := catchOutput(pluginPath, true, "description")
		failOnError(err)

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
