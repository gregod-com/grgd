package profile

import (
	"strings"

	"github.com/gregod-com/grgd/pkg/helper"
	"github.com/urfave/cli/v2"
)

// AListProfiles ...
func AListProfiles(c *cli.Context) error {
	core := helper.GetExtractor().GetCore(c)
	ui := core.GetUI()
	profiles, err := core.GetConfig().GetAllProfiles()
	if err != nil {
		return err
	}

	rows := [][]string{}
	for k, prof := range profiles {
		var subrows [][]string
		for _, v := range prof.GetValues() {
			subrows = append(subrows, append([]string{""}, strings.SplitN(v, ":", 2)...))
		}
		rows = append(rows, []string{k})
		rows = append(rows, subrows...)
		rows = append(rows, []string{"----"})
	}
	ui.PrintTable([]string{"Name", "Key", "Value"}, rows)
	return nil
}

// ADeleteProfile ...
func ADeleteProfile(c *cli.Context) error {
	core := helper.GetExtractor().GetCore(c)
	ui := core.GetUI()
	log := core.GetLogger()
	profiles, err := core.GetConfig().GetAllProfiles()
	if err != nil {
		return err
	}
	for _, v := range c.Args().Slice() {
		if p, ok := profiles[v]; ok {
			log.Debugf("found profile %s", p.GetName())
			delete(profiles, v)
			err = core.GetConfig().RemoveProfile(p)
			if err != nil {
				return err
			}
			ui.Printf("Profile %s has been deleted.\n", p.GetName())
		} else {
			ui.Printf("No matching profile with name `%s` found.\n", v)
		}
	}
	return nil
}
