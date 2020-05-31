package actions

import (
	"fmt"
	"log"
	"os/exec"
	"sort"
	"strings"

	tm "github.com/buger/goterm"
	UI "github.com/gregod-com/grgd/ui"
	I "github.com/gregod-com/interfaces"
	"github.com/urfave/cli/v2"
)

func AInit(c *cli.Context) error {
	return nil
}

// SubAShortcut ...
var SubAShortcut = map[string]func(*cli.Context) error{
	"add": func(c *cli.Context) error {
		if c.NArg() != 2 {
			return cli.NewExitError("You should enter a workload and a shortcut", 5)
		}
		shortcut, workload := c.Args().Get(0), c.Args().Get(1)

		err := c.App.Metadata["iamconfig"].(I.IConfigObject).AddWorkloadShortcut(shortcut, workload)
		if err != nil {
			switch err.Error() {
			case "WorkloadNotFound":
				return cli.NewExitError("The workload "+tm.Color(workload, tm.CYAN)+" is not part of your stack.", 6)
			case "ShortcutExists":
				return cli.NewExitError("Shortcut '"+tm.Color(shortcut, tm.RED)+"' already exists and points to "+tm.Color(c.App.Metadata["iamconfig"].(I.IConfigObject).GetWorkloadByShortcut(shortcut), tm.CYAN), 7)
			default:
				return cli.NewExitError("Unexpected error occured", 0)
			}
		}
		fmt.Println("Added new shortcut: " + tm.Color(shortcut, tm.RED) + " -> " + tm.Color(workload, tm.CYAN))
		return nil
	},
	"remove": func(c *cli.Context) error {
		if c.NArg() > 1 {
			return cli.NewExitError("You should only one shortcut at a time", 5)
		}
		shortcut := c.Args().Get(0)
		err := c.App.Metadata["iamconfig"].(I.IConfigObject).RemoveWorkloadShortcut(shortcut)
		if err != nil {
			switch err.Error() {
			case "ShortcutNotFound":
				return cli.NewExitError("There is no shortcut "+tm.Color(shortcut, tm.RED), 8)
			default:
				return cli.NewExitError("Unexpected error occured", 0)
			}
			fmt.Println("Removed shortcut: " + tm.Color(shortcut, tm.RED))
		}
		return nil

	},
}

func AVolume(c *cli.Context) error {
	c.Set("print_volume", "true")
	UI.PrintWorkloadOverview(c)
	fmt.Println("")
	fmt.Println("ACTION: (volume)")
	fmt.Println("TODO")
	return nil
}

func ACertificates(c *cli.Context) error {
	c.Set("print_volume", "true")
	UI.PrintWorkloadOverview(c)
	fmt.Println("")
	fmt.Println("ACTION: (volume)")
	fmt.Println("TODO")
	return nil
}

func ADNS(c *cli.Context) error {
	return nil
}

func ASettings(c *cli.Context) error {
	return nil
}

func AContext(c *cli.Context) error {
	out, err := exec.Command("kubectl", "config", "get-contexts").Output()
	if err != nil {
		log.Fatal(err)
	}
	if c.NArg() > 0 {
		for _, word := range strings.Fields(string(out)) {
			if strings.Contains(word, c.Args().First()) {
				// fmt.Println("Setting current context to '" + word + "'")
				out, err := exec.Command("kubectl", "config", "use-context", word).Output()
				if err != nil {
					log.Fatal(err)
				}
				fmt.Println(string(out))
				break
			}
		}
	}
	return nil
}

func AfterContext(c *cli.Context) error {
	out, err := exec.Command("kubectl", "config", "get-contexts").Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(out))
	return nil
}

func AHelm(c *cli.Context) error {
	out, err := exec.Command("helm", "ls", "--all-namespaces").Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(out))
	return nil
}

// TranslateShortcuts ...
func TranslateShortcuts(c *cli.Context) []string {
	shortcuts := c.App.Metadata["iamconfig"].(I.IConfigObject).GetWorkloadShortcuts()
	workloads := make([]string, c.NArg())

	// for k, v := range c.Args().Slice() {
	// 	workloads[k] = v
	// 	if workload := shortcuts[v]; workload != "" {
	// 		workloads[k] = workload
	// 	}
	// }
	log.Println(shortcuts)
	return workloads
}

// UtilGetCurrentKubeContext ...
func UtilGetCurrentKubeContext() string {
	out, err := exec.Command("kubectl", "config", "current-context").Output()
	if err != nil {
		log.Fatal(err)
	}
	return string(out)
}

// PrintShortcuts ...
func PrintShortcuts(c *cli.Context) {
	fmt.Println("\nShortcuts: ")
	overviewmap := map[string][]string{}
	for shortcut, workload := range c.App.Metadata["iamconfig"].(I.IConfigObject).GetWorkloadShortcuts() {
		overviewmap[workload] = append(overviewmap[workload], shortcut)
	}
	sorted := []string{}

	for wl, shortcuts := range overviewmap {
		workloadline := ""
		workloadline += fmt.Sprintf("  %-*v=>", 20, wl)
		sort.Strings(shortcuts)
		for _, s := range shortcuts {
			workloadline += fmt.Sprintf(" '%v',", s)
		}
		workloadline += fmt.Sprintln()
		sorted = append(sorted, workloadline)
	}

	sort.Strings(sorted)
	for _, s := range sorted {
		fmt.Println(s)
	}
}
