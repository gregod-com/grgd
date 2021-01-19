package hack

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path"

	"github.com/gregod-com/grgd/controller/helper"
	"github.com/urfave/cli/v2"
)

func catchOutput(script string, silent bool, args ...string) (string, error) {
	cmd := exec.Command(script, args...)
	var out, errout bytes.Buffer
	cmd.Stdout = io.MultiWriter(os.Stdout, &out)
	cmd.Stderr = io.MultiWriter(os.Stderr, &errout)
	if silent {
		cmd.Stdout = &out
		cmd.Stderr = &errout
	}
	err := cmd.Run()
	return out.String() + errout.String(), err
}

// AExec ...
func AExec(c *cli.Context) error {
	core := helper.GetExtractor().GetCore(c)

	script := path.Join(core.GetConfig().GetProfile().GetBasePath(), "hack", c.Command.FullName())

	args := c.Args().Slice()
	prefix := "exec"
	if len(args) > 0 && (args[0] == "help" || args[0] == "version" || args[0] == "description") {
		prefix = args[0]
	}
	args = append([]string{prefix}, args...)

	if _, err2 := catchOutput(script, false, args...); err2 != nil {
		return err2
	}
	return nil
}

func checkLsExists() {
	path, err := exec.LookPath("ls")
	if err != nil {
		fmt.Printf("didn't find 'ls' executable\n")
	} else {
		fmt.Printf("'ls' executable is in '%s'\n", path)
	}
}
