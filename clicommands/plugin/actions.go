package plugin

import (
	"log"
	"sort"
	"strconv"
	"strings"

	"grgd/controller/helper"

	"github.com/gregod-com/grgdplugincontracts"
	"github.com/urfave/cli/v2"
)

// APluginList ...
func APluginList(c *cli.Context) error {
	core := helper.GetExtractor().GetCore(c)
	UI := core.GetUI()
	configObject := core.GetConfig()
	UI.Println(configObject, c)

	// index := pluginindex.CreatePluginIndexFromCLIContext(c)
	// head := []string{"Name", "Version", "Category", "Active", "Loaded", "URL"}

	// rows := sortPluginMetadataSlice(index.GetPluginListActive())

	// UI.Println("Active Plugins",c)
	// UI.PrintTable(c, head, rows)

	// rows = sortPluginMetadataSlice(index.GetPluginListInactive())

	// UI.Println(c, "")
	// UI.Println(c, "InActive Plugins")
	// UI.PrintTable(c, head, rows)

	// rows = sortPluginMetadataSlice(index.GetPluginListOffline())

	// UI.Println(c, "")
	// UI.Println(c, "Offline Plugins")
	// UI.PrintTable(c, head, rows)
	return nil
}

// APluginActivate ...
func APluginActivate(c *cli.Context) error {
	core := helper.GetExtractor().GetCore(c)
	UI := core.GetUI()

	var index grgdplugincontracts.IPluginIndex
	err := core.Get(&index)
	if err != nil {
		log.Fatalf("hellos %v", err.Error())
	}

	// index := pluginindex.CreatePluginIndexFromCLIContext(c)
	// plugname := c.Args().First()
	plugname := "franz"
	key := []string{}

	for _, v := range index.GetPluginList() {
		k := v.GetIdentifier()
		if strings.Contains(strings.ToLower(k), strings.ToLower(plugname)) {
			key = append(key, k)
		}
	}

	switch len(key) {
	case 0:
		UI.Println("No matching plugin names")
	case 1:
		UI.Println("Found Match", c)
		index.ToggleActive(key[0])
		APluginList(c)
	default:
		UI.Println("Multiple plugins found:")
		for _, k := range key {
			UI.Println(k)
		}
	}
	return nil
}

func sortPluginMetadataSlice(m []grgdplugincontracts.IPluginMetadata) [][]string {
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
		row = append(row, key.GetVersion())
		row = append(row, key.GetCategory())
		row = append(row, strconv.FormatBool(key.GetActive()))
		row = append(row, strconv.FormatBool(key.GetLoaded()))
		row = append(row, key.GetURL())
		rows = append(rows, row)
	}

	return rows
}
