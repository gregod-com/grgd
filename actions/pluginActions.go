package actions

import (
	"log"
	"sort"
	"strconv"
	"strings"

	idx "github.com/gregod-com/grgd/pluginindex"
	plugContracts "github.com/gregod-com/grgdplugincontracts"
	cli "github.com/urfave/cli/v2"
)

// CreatePluginIndexFromCLIContext ...
func CreatePluginIndexFromCLIContext(c *cli.Context) plugContracts.IPluginIndex {
	pluginIndexPath, ok := c.App.Metadata["pluginIndex"].(string)
	if !ok {
		log.Fatal("Undefined Pluginindex")
	}
	return idx.CreatePluginIndex(pluginIndexPath)
}

// APluginList ...
func APluginList(c *cli.Context) error {
	index := CreatePluginIndexFromCLIContext(c)
	UI := c.App.Metadata["UIPlugin"].(plugContracts.IUIPlugin)
	head := []string{"Name", "Version", "URL", "Category", "Active"}

	rows := sortKeys(index.GetPluginListActive())

	if len(rows) > 0 {
		UI.Println(c, "Active Plugins")
		UI.PrintTable(c, head, rows)
	}

	rows = sortKeys(index.GetPluginListInactive())

	if len(rows) > 0 {
		UI.Println(c, "")
		UI.Println(c, "InActive Plugins")
		UI.PrintTable(c, head, rows)
	}
	return nil
}

// APluginActivate ...
func APluginActivate(c *cli.Context) error {
	UI := c.App.Metadata["UIPlugin"].(plugContracts.IUIPlugin)
	index := CreatePluginIndexFromCLIContext(c)
	plugname := c.Args().First()
	key := []string{}

	for k := range index.GetPluginList() {
		if strings.Contains(strings.ToLower(k), strings.ToLower(plugname)) {
			key = append(key, k)
		}
	}

	switch len(key) {
	case 0:
		UI.Println(c, "No matching plugin names")
	case 1:
		index.GetPluginList()[key[0]].ToggleActive()
		index.Update()
		APluginList(c)
	default:
		UI.Println(c, "Multiple plugins found:")
		for _, k := range key {
			UI.Println(c, k)
		}
	}
	return nil
}

func sortKeys(m map[string]plugContracts.IPluginMetadata) [][]string {
	rows := [][]string{}
	keys := []string{}
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, key := range keys {
		row := []string{}
		row = append(row, m[key].GetName())
		row = append(row, m[key].GetVersion())
		row = append(row, m[key].GetURL())
		row = append(row, m[key].GetCategory())
		row = append(row, strconv.FormatBool(m[key].GetActive()))
		rows = append(rows, row)
	}

	return rows
}
