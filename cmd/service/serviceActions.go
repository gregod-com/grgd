package service

import (
	"errors"
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/gregod-com/grgd/interfaces"
	"github.com/gregod-com/grgd/pkg/helper"
	"github.com/gregod-com/grgd/pkg/project"
	"github.com/urfave/cli/v2"
)

// AListService ...
func AListService(c *cli.Context) error {
	core := helper.GetExtractor().GetCore(c)
	log := core.GetLogger()
	ui := core.GetUI()
	h := core.GetHelper()

	proj := core.GetConfig().GetActiveProfile().GetCurrentProject()
	if proj == nil {
		log.Warn("Current project is not set")
		return nil
	}

	obj, err := proj.ReadSettingsObject(h)
	if err != nil {
		return err
	}

	pm, ok := obj.(*project.ProjectMetadata)
	if !ok {
		return fmt.Errorf("Project Metadata not supported %T", obj)
	}

	for k, v := range pm.Services {
		mp, ok := v.(map[interface{}]interface{})
		if !ok {
			ui.Printf("Service %v is %T\n", k, v)
			continue
		}

		ui.Printf("%-20v\n", k)
		// service := proj.GetServiceByName(k)
		for key, value := range mp {
			if key == "active" {
				ui.Printf("%-4v%v: %v \n", "", key, value)
			}

			if key == "path" {
				yaml.Unmarshall
			}
		}
	}

	return nil
}

// ADeleteService ...
func ADeleteService(c *cli.Context) error {
	// var profile *persistence.Profile
	// helpers.ExtractMetadataFatal(c.App.Metadata, "profile", &profile)
	// current := getProjectByID(profile.Projects, profile.CurrentProjectID)

	// idx, err := getServiceIndexByName(current.Services, c.Args().First())
	// if err != nil {
	// 	AListService(c)
	// 	return nil
	// }

	// persistence.Remove(&current.Services[idx])
	// current.Services = append(current.Services[:idx], current.Services[idx+1:]...)

	AListService(c)
	return nil
}

// AAddService ...
func AAddService(c *cli.Context) error {
	// var profile *persistence.Profile
	// helpers.ExtractMetadataFatal(c.App.Metadata, "profile", &profile)
	// current := getProjectByID(profile.Projects, profile.CurrentProjectID)
	// UI := helpers.ExtractUI(c)

	// _, err := getServiceIndexByName(current.Services, c.Args().First())
	// if err == nil || c.Args().First() == "" {
	// 	UI.Println(c, "Service with name `"+c.Args().First()+"` already exists.")
	// 	AListService(c)
	// 	return nil
	// }

	// service := persistence.Service{Name: c.Args().First(), Path: current.Path}

	// // persistence.Init(&service)
	// current.Services = append(current.Services, service)

	AListService(c)
	return nil
}

// AEditService ...
func AEditService(c *cli.Context) error {

	return nil
}

// AToggleService ...
func AToggleService(c *cli.Context) error {

	return nil
}

// AGroupService ...
func AGroupService(c *cli.Context) error {

	return nil
}

func getServiceIndexByName(arr []interfaces.IService, name string) (int, error) {
	for k, v := range arr {
		if v.GetName() == name {
			return k, nil
		}
	}
	return 0, errors.New("Name not found")
}

func sortServiceMetadataSlice(m []interfaces.IService) [][]string {
	rows := [][]string{}

	sort.Slice(m, func(i, j int) bool {
		switch strings.Compare(m[i].GetName(), m[j].GetName()) {
		case -1:
			return true
		default:
			return false
		}
	})

	for _, key := range m {
		row := []string{}
		row = append(row, key.GetName())
		row = append(row, key.GetPath())
		row = append(row, strconv.FormatBool(key.GetActive()))
		rows = append(rows, row)
	}

	return rows
}
