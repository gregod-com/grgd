package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/common-nighthawk/go-figure"
	cli "github.com/urfave/cli/v2"
)

// ExecCommand ...
func ExecCommand(binary string, commandAndParams ...string) string {
	cmd := exec.Command(binary, commandAndParams...)
	stdout, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}

	return string(stdout)
}

// ExtractCliContext ...
func ExtractCliContext(i interface{}) *cli.Context {
	c, ok := i.(*cli.Context)
	if !ok {
		log.Fatal("unexpected implementation of cli.Context")
	}
	return c
}

type fallbackui struct {
}

// ClearScreen ...
func (p fallbackui) ClearScreen(i interface{}) interface{} {
	c := ExtractCliContext(i)
	//  do not clear screen if in debug mode
	if c.Bool("debug") {
		return nil
	}

	out := ExecCommand("clear", "cmd", "/c", "cls")
	fmt.Println(out)
	return nil
}

// PrintPercentOfScreen ...
func (p fallbackui) PrintPercentOfScreen(i interface{}, str string, percent int) interface{} {

	// percent int, c string
	// TODO find way to easyly get terminal width
	y := 100
	y = int(y*percent) / 100

	stringBuffer := strings.Repeat(str, y)
	fmt.Print(stringBuffer)
	return nil
}

// PrintBanner ...
func (p fallbackui) PrintBanner(i interface{}) interface{} {
	c := ExtractCliContext(i)

	iamASCII := figure.NewFigure(c.App.Name, "standard", true)
	fmt.Println(iamASCII.String())

	return nil
}

// PrintWorkloadOverview ...
func (p fallbackui) PrintWorkloadOverview(i interface{}) {
	// c := ExtractCliContext(i)
	// config := c.App.Metadata["iamconfig"].(I.IConfigObject)
}

// PrintTable ...
func (p fallbackui) PrintTable(i interface{}, heads []string, rows [][]string) interface{} {

	for _, v := range heads {
		fmt.Printf("| %-20v", v)
	}
	fmt.Println("")
	for range heads {
		fmt.Printf("| %-20v", "------")
	}
	fmt.Println("")
	for _, v := range rows {
		for _, r := range v {
			fmt.Printf("| %-20v", r)
		}
		fmt.Println()
	}

	return nil
}

// Println ...
func (p fallbackui) Println(i interface{}, str interface{}) interface{} {
	// c := ExtractCliContext(i)
	fmt.Println(str)
	return nil
}

// YesNoQuestion ...
func (p fallbackui) YesNoQuestion(i interface{}, question string) bool {
	reader := bufio.NewReader(os.Stdin)
	answer, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	answer = strings.Replace(answer, "\n", "", -1)
	if strings.Contains(answer, "y") {
		return true
	}
	return false
}

// // PrintActiveWorkload ...
// func (p uisimple) PrintActiveWorkload(c *cli.Context, w I.IWorkload, config I.IConfigObject, line int) {
// 	shared.Debug(c, "called")
// 	// ----------------------
// 	_, _, activeBool := w.GetActive()

// 	if activeBool {
// 		// imageAndTag := p.ReplaceRegistries(w.GetPod().GetMainContainer().GetImage().GetFullName(), config.GetRegistries(), w.GetName())

// 		p.PrintNetworkDetails(c, w.GetPod().GetMainContainer(), 4)
// 		p.PrintVolumeDetails(c, w)
// 		p.PrintSidecars(c, w, config)
// 	}
// }

// // PrintSidecars ...
// func (p uisimple) PrintSidecars(c *cli.Context, s I.IWorkload, config I.IConfigObject) {
// 	shared.Debug(c, "called")
// 	// ----------------------
// }

// // PrintNetworkDetails ...
// func (p uisimple) PrintNetworkDetails(c *cli.Context, s I.IContainer, indent int) {
// 	shared.Debug(c, "called")
// 	// ----------------------
// 	if c.Bool("network") {
// 		fmt.Println("TODO Network")
// 	}
// }

// // PrintVolumeDetails ...
// func (p uisimple) PrintVolumeDetails(c *cli.Context, w I.IWorkload) {
// 	shared.Debug(c, "called")
// 	// ----------------------
// }

// // PrintInactiveWorkload ...
// func (p uisimple) PrintInactiveWorkload(c *cli.Context, s I.IWorkload) {
// 	shared.Debug(c, "called")
// 	// ----------------------
// }

// // PrintExecutionTime ...
// func (p uisimple) PrintExecutionTime(d time.Duration) {
// }

// func sortMapAlphabetically(WorkloadMap map[string]I.IWorkload) []I.IWorkload {
// 	sorted := []I.IWorkload{}
// 	for _, v := range WorkloadMap {
// 		sorted = append(sorted, v)
// 	}
// 	sort.Slice(sorted, func(i, j int) bool {
// 		return sorted[i].GetName() < sorted[j].GetName()
// 	})
// 	return sorted
// }
