package view

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"github.com/common-nighthawk/go-figure"
	"github.com/gregod-com/grgd/interfaces"
	cli "github.com/urfave/cli/v2"
	"golang.org/x/sys/unix"
	"golang.org/x/term"
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
		if i == nil {
			return &cli.Context{}
		}
		log.Fatal("unexpected implementation of cli.Context")
	}
	return c
}

// ProvideFallbackUI ...
func ProvideFallbackUI() interfaces.IUIPlugin {
	return new(FallbackUI)
}

// FallbackUI ...
type FallbackUI struct {
}

// ClearScreen ...
func (p FallbackUI) ClearScreen(i ...interface{}) interface{} {
	switch runtime.GOOS {
	case "linux", "darwin":
		out := ExecCommand("clear")
		fmt.Println(out)
	case "windows":
		out := ExecCommand("cmd", "/c", "cls")
		fmt.Println(out)
	default:
		fmt.Printf("ClearScreen not implemented for %v\n", runtime.GOOS)
	}

	return nil
}

// PrintPercentOfScreen ...
func (p FallbackUI) PrintPercentOfScreen(percentStart int, percentEnd int, str ...interface{}) interface{} {
	y := 100
	ws, err := getWinsize()
	if err == nil {
		y = int(ws.Col)
	}

	if percentStart > 100 {
		percentStart = 100
	}
	if percentStart < 0 {
		percentStart = 0
	}
	if percentEnd > 100 {
		percentEnd = 100
	}
	if percentEnd < 0 {
		percentEnd = 0
	}

	if percentStart > percentEnd {
		temp := percentEnd
		percentEnd = percentStart
		percentStart = temp
	}

	distance := percentEnd - percentStart

	start := int(y*percentStart) / 100
	width := int(y*distance) / 100
	for _, k := range str {
		s, ok := k.(string)
		if !ok {
			continue
		}
		times := width / len(s)
		stringBuffer := strings.Repeat(" ", start)
		stringBuffer += strings.Repeat(s, times)
		fmt.Print(stringBuffer)
	}
	return nil
}

// PrintBanner ...
func (p FallbackUI) PrintBanner(i ...interface{}) interface{} {
	c := ExtractCliContext(i[0])
	core := c.App.Metadata["core"].(interfaces.ICore)

	iamASCII := figure.NewFigure(c.App.Name, "standard", true)
	fmt.Println(iamASCII.String())
	fmt.Printf("version: \u001b[33m %s\u001b[0m ||| profile: \u001b[33m%s\u001b[0m ||| k8s-ctx: \u001b[33m%s\u001b[0m \n\n",
		c.App.Version,
		core.GetConfig().GetProfile().GetName(),
		c.App.Metadata["kubeContext"])
	return nil
}

// PrintTable ...
func (p FallbackUI) PrintTable(heads []string, rows [][]string, i ...interface{}) interface{} {
	w, _, err := term.GetSize(2)
	if err != nil {
		fmt.Println(err)
	}
	colwidth := fmt.Sprint(w/len(heads) - len(heads) - 1)

	for _, v := range heads {
		fmt.Printf("| %-"+colwidth+"v", v)
	}
	fmt.Println("")
	for range heads {
		fmt.Printf("| %-"+colwidth+"v", "------")
	}
	fmt.Println("")
	for _, v := range rows {
		for _, r := range v {
			fmt.Printf("| %-"+colwidth+"v", r)
		}
		fmt.Println()
	}
	fmt.Println()
	return nil
}

// Println ...
func (p FallbackUI) Println(i ...interface{}) (int, error) {
	return fmt.Println(i...)
}

// Printf ...
func (p FallbackUI) Printf(format string, a ...interface{}) (int, error) {
	return fmt.Printf(format, a...)
}

// YesNoQuestion ...
func (p FallbackUI) YesNoQuestion(question string, i ...interface{}) bool {
	fmt.Print(question + " [y/n] ")
	answer := readLine()
	if strings.Contains(answer, "y") {
		return true
	}
	return false
}

// YesNoQuestionf ...
func (p FallbackUI) YesNoQuestionf(question string, i ...interface{}) bool {
	fmt.Printf(question+" [y/n] ", i...)
	answer := readLine()
	if strings.Contains(answer, "y") {
		return true
	}
	return false
}

// Question ...
func (p FallbackUI) Question(question string, i ...interface{}) error {
	fmt.Print(question)
	if len(i) > 0 {
		answer, ok := i[0].(*string)
		if input := readLine(); ok && input != "" {
			*answer = input
		}
	}
	return nil
}

// Questionf ...
func (p FallbackUI) Questionf(question string, i ...interface{}) error {
	fmt.Printf(question, i[1:])
	if len(i) > 0 {
		answer, ok := i[0].(*string)
		if input := readLine(); ok && input != "" {
			*answer = input
		}
	}
	return nil
}

func readLine() string {
	reader := bufio.NewReader(os.Stdin)
	answer, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	answer = strings.Replace(answer, "\n", "", -1)
	return answer
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

func getWinsize() (*unix.Winsize, error) {

	ws, err := unix.IoctlGetWinsize(int(os.Stdout.Fd()), unix.TIOCGWINSZ)
	if err != nil {
		return nil, os.NewSyscallError("GetWinsize", err)
	}

	return ws, nil
}
