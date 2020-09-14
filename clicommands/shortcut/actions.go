package shortcut

import (
	"grgd/controller/helper"

	"github.com/urfave/cli/v2"
)

// SubAShortcutAddDescription ...
const SubAShortcutAddDescription = `This is the description for shortcut add`

// SubAShortcutAdd ...
func SubAShortcutAdd(c *cli.Context) error {
	core := helper.GetExtractor().GetCore(c)
	UI := core.GetUI()
	// config := core.GetConfig()

	UI.Println(c, "\nShortcuts: [add]")

	// if c.NArg() != 2 {
	// 	return cli.NewExitError("You should enter a workload and a shortcut", 5)
	// }
	// shortcut, workload := c.Args().Get(0), c.Args().Get(1)

	// err := config.AddWorkloadShortcut(shortcut, workload)
	// if err != nil {
	// 	switch err.Error() {
	// 	case "WorkloadNotFound":
	// 		return cli.NewExitError("The workload "+workload+" is not part of your stack.", 6)
	// 	case "ShortcutExists":
	// 		return cli.NewExitError("Shortcut '"+shortcut+"' already exists and points to "+config.GetWorkloadByShortcut(shortcut), 7)
	// 	default:
	// 		return cli.NewExitError("Unexpected error occured", 0)
	// 	}
	// }
	// UI.Println(c, "Added new shortcut: "+shortcut+" -> "+workload)
	return nil
}

// SubAShortcutRemoveDescription ...
const SubAShortcutRemoveDescription = `This is the description for shortcut remove`

// SubAShortcutRemove ...
func SubAShortcutRemove(c *cli.Context) error {
	core := helper.GetExtractor().GetCore(c)
	UI := core.GetUI()
	// config := core.GetConfig()

	UI.Println(c, "\nShortcuts: [remove]")

	// if c.NArg() > 1 {
	// 	return cli.NewExitError("You should only one shortcut at a time", 5)
	// }
	// shortcut := c.Args().Get(0)
	// err := config.RemoveWorkloadShortcut(shortcut)
	// if err != nil {
	// 	switch err.Error() {
	// 	case "ShortcutNotFound":
	// 		return cli.NewExitError("There is no shortcut "+shortcut, 8)
	// 	default:
	// 		return cli.NewExitError("Unexpected error occured", 0)
	// 	}
	// }
	// UI.Println(c, "Removed shortcut: "+shortcut)
	return nil

}

// SubAShortcutListDescription ...
const SubAShortcutListDescription = `This is the description for shortcut add`

// SubAShortcutList ...
func SubAShortcutList(c *cli.Context) error {
	core := helper.GetExtractor().GetCore(c)
	UI := core.GetUI()
	// config := core.GetConfig()

	UI.Println(c, "\nShortcuts: ")
	// overviewmap := map[string][]string{}
	// for shortcut, workload := range config.GetWorkloadShortcuts() {
	// 	overviewmap[workload] = append(overviewmap[workload], shortcut)
	// }
	// sorted := []string{}

	// for wl, shortcuts := range overviewmap {
	// 	workloadline := ""
	// 	workloadline += fmt.Sprintf("  %-*v=>", 20, wl)
	// 	sort.Strings(shortcuts)
	// 	for _, s := range shortcuts {
	// 		workloadline += fmt.Sprintf(" '%v',", s)
	// 	}
	// 	workloadline += fmt.Sprintln()
	// 	sorted = append(sorted, workloadline)
	// }

	// sort.Strings(sorted)
	// for _, s := range sorted {
	// 	UI.Println(c, s)
	// }
	return nil
}

// TranslateShortcuts ...
func TranslateShortcuts(c *cli.Context) []string {
	core := helper.GetExtractor().GetCore(c)
	UI := core.GetUI()
	// config := core.GetConfig()

	UI.Println(c, "\nShortcuts: Translate")

	// shortcuts := config.GetWorkloadShortcuts()
	workloads := make([]string, c.NArg())

	// for k, v := range c.Args().Slice() {
	// 	workloads[k] = v
	// 	if workload := shortcuts[v]; workload != "" {
	// 		workloads[k] = workload
	// 	}
	// }
	// UI.Println(c, shortcuts)
	return workloads
}
