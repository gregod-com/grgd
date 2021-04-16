package project

import (
	"github.com/gregod-com/grgd/pkg/helper"
	"github.com/urfave/cli/v2"
)

// AListProject ...
func AListProject(c *cli.Context) error {
	core := helper.GetExtractor().GetCore(c)
	UI := core.GetUI()
	profile := core.GetConfig().GetActiveProfile()

	rows := profile.GetProjectsTable()
	UI.PrintTable(rows[0], rows[1:], c)

	return nil
}

// ASwitchProject ...
func ASwitchProject(c *cli.Context) error {
	core := helper.GetExtractor().GetCore(c)
	prof := core.GetConfig().GetActiveProfile()

	if p, ok := prof.GetProjects()[c.Args().First()]; ok {
		prof.SetCurrentProjectID(p.GetID())
	}

	AListProject(c)
	return nil
}

// ADeleteProject ...
func ADeleteProject(c *cli.Context) error {
	core := helper.GetExtractor().GetCore(c)
	conf := core.GetConfig()
	ui := core.GetUI()
	prof := core.GetConfig().GetActiveProfile()

	ps := prof.GetProjects()
	AListProject(c)
	name := c.Args().First()
	for name == "" {
		ui.Question("What project do you want to delete? ", &name)
	}
	if _, ok := ps[name]; !ok {
		ui.Printf("The project `%s` does not exists, skipping delete.\n", name)
		return nil
	}
	conf.Remove(ps[name])
	prof.RemoveProject(ps[name])
	AListProject(c)
	return nil
}

// AAddProject ...
func AAddProject(c *cli.Context) error {
	core := helper.GetExtractor().GetCore(c)
	ui := core.GetUI()
	prof := core.GetConfig().GetActiveProfile()
	ps := prof.GetProjects()

	AListProject(c)

	name := c.Args().First()
	for name == "" {
		ui.Question("What's the name of the new project? ", &name)
	}
	if _, ok := ps[name]; ok {
		ui.Printf("The project `%s` already exists, please use unique name.\n", name)
		return nil
	}
	prof.AddProjectByName(name)
	AListProject(c)
	if ui.YesNoQuestionf("Init project `%s` now?", name) {
		if err := prof.GetProjects()[name].Init(core); err != nil {
			return err
		}
		AListProject(c)
	}
	return nil
}

// AEditProject ...
func AEditProject(c *cli.Context) error {
	core := helper.GetExtractor().GetCore(c)
	ui := core.GetUI()
	prof := core.GetConfig().GetActiveProfile()
	ps := prof.GetProjects()
	AListProject(c)

	name := c.Args().First()
	for name == "" {
		ui.Question("What's the name of the new project? ", &name)
	}
	if _, ok := ps[name]; !ok {
		ui.Printf("The project `%s` does not exists. If you want to add a project use the `add` command.\n", name)
		return nil
	}
	err := ps[name].Init(core)
	AListProject(c)
	return err
}
