package actions

import (
	"log"
	"strconv"

	idx "github.com/gregod-com/grgd/pluginindex"
	I "github.com/gregod-com/interfaces"
	cli "github.com/urfave/cli/v2"
)

func CreatePluginIndexFromCLIContext(c *cli.Context) I.IPluginIndex {
	pluginIndexPath, ok := c.App.Metadata["pluginIndex"].(string)
	if !ok {
		log.Fatal("Undefined Pluginindex")
	}
	return idx.CreatePluginIndex(pluginIndexPath)
}

// APluginList ...
func APluginList(c *cli.Context) error {
	index := CreatePluginIndexFromCLIContext(c)
	UI := c.App.Metadata["UIPlugin"].(I.IUIPlugin)

	head := []string{"Name", "Version", "Size", "URL", "Category", "Active"}
	rows := [][]string{}

	UI.Println(c, "Active Plugins")
	for _, plug := range index.GetPluginList() {
		if plug.GetActive() {
			row := []string{}
			row = append(row, plug.GetName())
			row = append(row, plug.GetVersion())
			row = append(row, strconv.FormatUint(plug.GetSize(), 10))
			row = append(row, plug.GetURL())
			row = append(row, plug.GetCategory())
			row = append(row, strconv.FormatBool(plug.GetActive()))
			rows = append(rows, row)
		}
	}

	UI.PrintTable(c, head, rows)

	head = []string{"Name", "Version", "Size", "URL", "Category", "Active"}
	rows = [][]string{}

	UI.Println(c, "\nInActive Plugins")
	for _, plug := range index.GetPluginList() {
		if plug.GetActive() == false {
			row := []string{}
			row = append(row, plug.GetName())
			row = append(row, plug.GetVersion())
			row = append(row, strconv.FormatUint(plug.GetSize(), 10))
			row = append(row, plug.GetURL())
			row = append(row, plug.GetCategory())
			row = append(row, strconv.FormatBool(plug.GetActive()))
			rows = append(rows, row)
		}
	}

	UI.PrintTable(c, head, rows)
	return nil
}

// APluginActivate ...
func APluginActivate(c *cli.Context) error {
	index := CreatePluginIndexFromCLIContext(c)
	index.GetPluginList()["commands-DNS"].SetActive(!index.GetPluginList()["commands-DNS"].GetActive())
	index.Update()
	APluginList(c)
	return nil
}
