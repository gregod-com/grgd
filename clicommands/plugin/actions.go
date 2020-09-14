package plugin

import (
	"sort"
	"strconv"
	"strings"

	"grgd/controller/helper"
	"grgd/controller/pluginindex"

	"github.com/gregod-com/grgdplugincontracts"
	cli "github.com/urfave/cli/v2"
)

// APluginList ...
func APluginList(c *cli.Context) error {
	ext := helper.GetExtractor()
	UI := ext.GetCore(c).GetUI()
	configObject := ext.GetCore(c).GetConfig()
	UI.Println(c, configObject)

	// index := pluginindex.CreatePluginIndexFromCLIContext(c)
	// head := []string{"Name", "Version", "Category", "Active", "Loaded", "URL"}

	// rows := sortPluginMetadataSlice(index.GetPluginListActive())

	// UI.Println(c, "Active Plugins")
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
	ext := helper.GetExtractor()
	UI := ext.GetCore(c).GetUI()

	index := pluginindex.CreatePluginIndexFromCLIContext(c)
	plugname := c.Args().First()
	key := []string{}

	for _, v := range index.GetPluginList() {
		k := v.GetIdentifier()
		if strings.Contains(strings.ToLower(k), strings.ToLower(plugname)) {
			key = append(key, k)
		}
	}

	switch len(key) {
	case 0:
		UI.Println(c, "No matching plugin names")
	case 1:
		UI.Println(c, "Found Match")
		index.ToggleActive(key[0])
		APluginList(c)
	default:
		UI.Println(c, "Multiple plugins found:")
		for _, k := range key {
			UI.Println(c, k)
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
