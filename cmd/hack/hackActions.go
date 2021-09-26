package hack

import (
	"path"

	"github.com/gregod-com/grgd/pkg/helper"
	"github.com/urfave/cli/v2"
)

// AExec ...
func AExec(c *cli.Context) error {
	core := helper.GetExtractor().GetCore(c)
	// UI := core.GetUI()
	h := core.GetHelper()
	log := core.GetLogger()
	script := path.Join(
		core.GetConfig().GetActiveProfile().GetMetaData("hackDir"),
		c.Command.FullName())

	args := c.Args().Slice()
	prefix := "exec"
	if len(args) > 0 && (args[0] == "help" || args[0] == "version" || args[0] == "description") {
		prefix = args[0]
	}

	args = append([]string{prefix}, args...)
	log.Debug("Calling plugin with arguments: %v\n", args)

	if _, err2 := h.CatchOutput(script, false, args...); err2 != nil {
		return err2
	}
	return nil
}
