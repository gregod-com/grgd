package project

import (
	"errors"

	"github.com/gregod-com/grgd/controller/helper"
	"github.com/urfave/cli/v2"
)

// AListProject ...
func AListProject(c *cli.Context) error {
	core := helper.GetExtractor().GetCore(c)
	UI := core.GetUI()
	profile := core.GetConfig().GetProfile()

	rows := profile.GetProjectsTable()
	UI.PrintTable(rows[0], rows[1:], c)

	return nil
}

// ASwitchProject ...
func ASwitchProject(c *cli.Context) error {
	core := helper.GetExtractor().GetCore(c)
	profile := core.GetConfig().GetProfile()

	if p, ok := profile.GetProjects()[c.Args().First()]; ok {
		profile.SetCurrentProject(p)
	}

	AListProject(c)
	return nil
}

// ADeleteProject ...
func ADeleteProject(c *cli.Context) error {
	core := helper.GetExtractor().GetCore(c)
	cnfg := core.GetConfig()

	ps, err := cnfg.GetProjects()
	if err != nil {
		return nil
	}
	if p, ok := ps[c.Args().First()]; ok {
		cnfg.RemoveProject(p)
	}

	AListProject(c)
	return nil
}

// AAddProject ...
func AAddProject(c *cli.Context) error {
	core := helper.GetExtractor().GetCore(c)
	cnfg := core.GetConfig()

	arg := c.Args().First()
	ps, err := cnfg.GetProjects()
	if err != nil {
		return errors.New("bla")
	}
	if _, ok := ps[arg]; !ok && arg != "" {
		cnfg.AddProject(arg)
	}

	AListProject(c)
	return nil
}

// AEditProject ...
func AEditProject(c *cli.Context) error {
	core := helper.GetExtractor().GetCore(c)
	profile := core.GetConfig().GetProfile()
	UI := core.GetUI()

	UI.Println(profile)

	return nil
}
