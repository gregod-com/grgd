package actions

import (
	"fmt"

	I "github.com/gregod-com/interfaces"
	"github.com/urfave/cli"
)

func AEnter(c *cli.Context) error {
	go fmt.Println("this is the enter command")
	return nil
}

func AExecute(c *cli.Context) error {
	go fmt.Println("this is the execute command")
	return nil
}

func ATest(c *cli.Context) error {
	go fmt.Println("this is the test command")
	return nil
}
func AActivate(c *cli.Context) error {
	wls := TranslateShortcuts(c)

	configObject := c.App.Metadata["iamconfig"].(I.IConfigObject)

	for _, workloadToActivate := range wls {
		for _, wl := range configObject.GetWorkloadMetadata() {
			if workloadToActivate == wl.GetName() {
				wl.ToggleActive()
			}
		}
	}
	return nil
}
