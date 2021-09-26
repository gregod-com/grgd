package view

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"sort"
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
func ProvideFallbackUI(logger interfaces.ILogger) interfaces.IUIPlugin {
	ui := new(FallbackUI)
	logger.Tracef("provide %T", ui)
	ui.logger = logger
	return ui
}

// FallbackUI ...
type FallbackUI struct {
	logger interfaces.ILogger
}

// ClearScreen ...
func (ui FallbackUI) ClearScreen(i ...interface{}) interface{} {
	ui.logger.Tracef("")
	if ui.logger.GetLevel() != "trace" && ui.logger.GetLevel() != "debug" {
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
	}
	return nil
}

// PrintPercentOfScreen ...
func (ui FallbackUI) PrintPercentOfScreen(percentStart int, percentEnd int, str ...interface{}) interface{} {
	ui.logger.Tracef("")

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
func (ui FallbackUI) PrintBanner(i ...interface{}) interface{} {
	ui.logger.Tracef("")
	c := ExtractCliContext(i[0])
	core, ok := c.App.Metadata["core"].(interfaces.ICore)
	if !ok {
		ui.logger.Fatalf("Missing core in metadata!")
	}
	unsorted := []string{}
	longestMetaKey := 0
	longestMetaValue := 0
	for k, v := range core.GetConfig().GetActiveProfile().GetValuesAsMap() {
		unsorted = append(unsorted, k)
		if longestMetaKey < len(k) {
			longestMetaKey = len(k)
		}
		if longestMetaValue < len(v) {
			longestMetaValue = len(v)
		}
	}
	longestMetaKey++
	longestMetaValue++

	valueSpace := 30
	meta := []string{}
	currentProjName := "---"
	if p := core.GetConfig().GetActiveProfile().GetCurrentProject(); p != nil {
		currentProjName = p.GetName()
	}
	for key, value := range map[string]string{
		"version": c.App.Version,
		"profile": core.GetConfig().GetActiveProfile().GetName(),
		"project": currentProjName,
	} {
		meta = append(meta, fmt.Sprintf("%-*s \u001b[33m %-*s\u001b[0m", longestMetaKey, key, valueSpace, value))
	}

	sort.Strings(unsorted)
	for _, key := range unsorted {
		val := core.GetConfig().GetActiveProfile().GetMetaData(key)
		if len(val) > valueSpace {
			val = val[:valueSpace-3]
			val += "..."
		}
		if val == "" {
			val = "--"
		}

		meta = append(meta, fmt.Sprintf("%-*s \u001b[33m %-*s\u001b[0m", longestMetaKey, key, valueSpace, val))
	}

	// fmt.Srintf("version: \u001b[33m %s\u001b[0m", c.App.Version),
	// ||| profile: \u001b[33m%s\u001b[0m",
	// core.GetConfig().GetActiveProfile().GetName(),

	ASCII := figure.NewFigure(c.App.Name, "standard", true)
	longestLine := 0
	nrOfBannerLines := len(ASCII.Slicify())
	for _, line := range ASCII.Slicify() {
		if len(line) > longestLine {
			longestLine = len(line)
		}
	}
	for k, line := range ASCII.Slicify() {
		tag1 := ""
		tag2 := ""
		if k <= len(meta)-1 {
			tag1 = meta[k]
		}
		if nrOfBannerLines < len(meta) && k+nrOfBannerLines <= len(meta)-1 {
			tag2 = meta[k+nrOfBannerLines]
		}

		fmt.Printf("%-*s%s| %s\n", longestLine+2, line, tag1, tag2)
	}
	fmt.Println()
	return nil
}

// PrintTable ...
func (ui FallbackUI) PrintTable(heads []string, rows [][]string, i ...interface{}) interface{} {
	ui.logger.Tracef("")
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
func (ui FallbackUI) Println(i ...interface{}) (int, error) {
	ui.logger.Tracef("")
	return fmt.Println(i...)
}

// Printf ...
func (ui FallbackUI) Printf(format string, a ...interface{}) (int, error) {
	ui.logger.Tracef("")
	return fmt.Printf(format, a...)
}

// YesNoQuestion ...
func (ui FallbackUI) YesNoQuestion(question string, i ...interface{}) bool {
	ui.logger.Tracef("")
	fmt.Print(question + " [y/n] ")
	answer := readLine()
	if strings.Contains(answer, "y") {
		return true
	}
	return false
}

// YesNoQuestionf ...

func (ui FallbackUI) YesNoQuestionf(question string, i ...interface{}) bool {
	ui.logger.Tracef("")
	fmt.Printf(question+" [y/n] ", i...)
	answer := readLine()
	if strings.Contains(answer, "y") {
		return true
	}
	return false
}

// Question ...
func (ui FallbackUI) Question(question string, i ...interface{}) error {
	ui.logger.Tracef("")
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
func (ui FallbackUI) Questionf(question string, answer *string, i ...interface{}) error {
	ui.logger.Tracef("")
	fmt.Printf(question, i...)
	if input := readLine(); input != "" {
		*answer = input
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
// func (ui uisimple) PrintActiveWorkload(c *cli.Context, w I.IWorkload, config I.IConfigObject, line int) {
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
// func (ui uisimple) PrintSidecars(c *cli.Context, s I.IWorkload, config I.IConfigObject) {
// 	shared.Debug(c, "called")
// 	// ----------------------
// }

// // PrintNetworkDetails ...
// func (ui uisimple) PrintNetworkDetails(c *cli.Context, s I.IContainer, indent int) {
// 	shared.Debug(c, "called")
// 	// ----------------------
// 	if c.Bool("network") {
// 		fmt.Println("TODO Network")
// 	}
// }

// // PrintVolumeDetails ...
// func (ui uisimple) PrintVolumeDetails(c *cli.Context, w I.IWorkload) {
// 	shared.Debug(c, "called")
// 	// ----------------------
// }

// // PrintInactiveWorkload ...
// func (ui uisimple) PrintInactiveWorkload(c *cli.Context, s I.IWorkload) {
// 	shared.Debug(c, "called")
// 	// ----------------------
// }

// // PrintExecutionTime ...
// func (ui uisimple) PrintExecutionTime(d time.Duration) {
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
