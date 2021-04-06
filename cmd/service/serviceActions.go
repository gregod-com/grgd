package service

import (
	"errors"
	"sort"
	"strconv"
	"strings"

	"github.com/gregod-com/grgd/interfaces"
	"github.com/gregod-com/grgd/pkg/helper"
	"github.com/urfave/cli/v2"
)

// AListService ...
func AListService(c *cli.Context) error {
	core := helper.GetExtractor().GetCore(c)
	logger := core.GetLogger()
	logger.Trace("This is the list service command")

	// var profile *persistence.Profile
	// helpers.ExtractMetadataFatal(c.App.Metadata, "profile", &profile)
	// current := getProjectByID(profile.Projects, profile.CurrentProjectID)

	// head := []string{"Name", "Path", "Description"}
	// rows := sortServiceMetadataSlice(current.Services)
	// logger.Trace(profile.Projects)

	// UI.PrintTable(c, head, rows)

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
