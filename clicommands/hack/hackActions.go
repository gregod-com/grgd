package hack

import (
	"bytes"
	"os"
	"os/exec"
	"path"

	"github.com/gregod-com/grgd/controller/helper"
	"github.com/urfave/cli/v2"
)

func catchOutput(script string, args ...string) (string, error) {
	cmd := exec.Command(script, args...)
	var out bytes.Buffer
	var errout bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &errout

	e := cmd.Run()
	if e != nil {
		return out.String() + errout.String(), e
	}
	return out.String(), nil
}

// AExec ...
func AExec(c *cli.Context) error {
	core := helper.GetExtractor().GetCore(c)
	UI := core.GetUI()
	homedir, err := os.UserHomeDir()
	if err != nil {
		return nil
	}

	script := path.Join(homedir, ".grgd", "hack", c.Command.FullName())
	args := c.Args().Slice()
	args = append([]string{"exec"}, args...)
	out, err := catchOutput(script, args...)
	UI.Println(out)
	if err != nil {
		return err
	}
	return nil
}
