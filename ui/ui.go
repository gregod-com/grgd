package ui

import (
	"fmt"
	"log"
	"os"
	"reflect"
	"runtime"
	"sort"
	"strconv"
	"strings"
	"time"

	at "github.com/gregod-com/animaterm"
	I "github.com/gregod-com/interfaces"

	"github.com/urfave/cli/v2"

	figure "github.com/common-nighthawk/go-figure"
)

var p = fmt.Print
var pl = fmt.Println
var pf = fmt.Printf
var colorOffset = 14
var globalOffset = 0
var tempOffset = 0

// PrintPercentOfScreen ...
func PrintPercentOfScreen(percent int, c string) {
	y := at.Width()
	y = int(y*percent) / 100

	stringBuffer := strings.Repeat(c, y)
	p(stringBuffer)
}

// PrintFullScreenRepeatingChar ...
func PrintFullScreenRepeatingChar(left string, c string, right string, singleCharLeftRight bool) {
	p := fmt.Println
	y := at.Width()
	remain := y - (len(left) + len(right))
	if singleCharLeftRight {
		remain = y - 2
	}
	stringBuffer := strings.Repeat(c, remain)

	p(left + stringBuffer + right)
}

// PrintFullScreenCenteredString ...
func PrintFullScreenCenteredString(left string, c string, right string) {
	p := fmt.Println
	y := at.Width()
	remain := y - (len(left) + len(right) + len(c))
	remain /= 2

	stringBuffer := strings.Repeat(" ", remain)

	p(left + stringBuffer + c + stringBuffer + right)

}

// PrintFullScreenColored ...
func PrintFullScreenColored(parts []string, colors []int) {
	p := fmt.Println

	stringline := ""

	for k := range parts {
		stringline += at.Color(parts[k], colors[k])
	}

	p(stringline)
}

// PrintBanner ...
func PrintBanner(c *cli.Context) {
	config := c.App.Metadata["iamconfig"].(I.IConfigObject)
	userInterface := c.App.Metadata["iamui"].(at.IUserInterface)
	ch, wg := c.App.Metadata["iamui"].(at.IUserInterface).StartDrawLoop(25)

	iamASCII := figure.NewFigure(c.App.Name, "shadow", true)
	defaultAnimation := at.Animation{
		AnimationType: at.Ikea,
		Duration:      1000 * c.App.Metadata["animation"].(int64),
		Direction:     at.Right,
		GradientV:     true,
		GradientH:     true,
	}

	userInterface.MoveElement(at.CreatePos(50, 5), at.CreatePos(0, 5),
		iamASCII.String(),
		at.COLORPATTERNMEADOWS1,
		defaultAnimation)

	userInterface.MoveElement(at.CreatePos(70, 5), at.CreatePos(45, 5),
		"v"+c.App.Version+" (go edition)\n"+
			config.GetProjectDir()+"\n"+
			c.App.Metadata["configLocation"].(string)+"\n"+
			"UI Implementation: "+reflect.TypeOf(c.App.Metadata["iamui"]).String()+"\n",
		at.COLORPATTERNMEADOWS1, defaultAnimation)

	go userInterface.MoveElement(at.CreatePos(85, 5), at.CreatePos(75, 5),
		"Connected Cluster: "+c.App.Metadata["currentcontext"].(string)+
			runtime.Version()+"\n"+
			"OS Architecture:   "+runtime.GOOS+"\n"+
			"todo\n",
		at.COLORPATTERNMEADOWS1, defaultAnimation)

	userInterface.DrawPattern(at.CreatePos(0, 18), 100, "─", at.COLORPATTERNLIME, defaultAnimation)

	close(ch)
	wg.Wait()
}

// PrintWorkloadOverview ...
func PrintWorkloadOverview(c *cli.Context) {

	config := c.App.Metadata["iamconfig"].(I.IConfigObject)
	userInterface := c.App.Metadata["iamui"].(at.IUserInterface)
	workloadMap := c.App.Metadata["workloads"].(map[string]I.IWorkload)
	activeWorkloads := [][]string{}
	workloadsInactive := []string{}

	k := 0
	// for _, s := range sortMapAlphabetically(workloadMap) {
	// 	log.Println(s.GetPod().GetMainContainer())
	// }

	for _, s := range sortMapAlphabetically(workloadMap) {
		if _, _, a := s.GetActive(); a {

			activeWorkloads = append(activeWorkloads, []string{})
			activeWorkloads[k] = append(activeWorkloads[k], s.GetName())
			activeWorkloads[k] = append(activeWorkloads[k], "down")
			image := s.GetPod().GetMainContainer().GetImage().GetRepositoryAsString()
			for regName, reg := range config.GetRegistries() {
				if reg == "" {
					continue
				}
				if strings.Contains(image, reg) {
					image = strings.Replace(image, reg, "["+regName+"]", 1)
				}
			}
			activeWorkloads[k] = append(activeWorkloads[k], image)
			activeWorkloads[k] = append(activeWorkloads[k], s.GetPod().GetMainContainer().GetImage().GetTagAsString())
			ports := ""
			for _, p := range s.GetPod().GetMainContainer().GetPortsAsString() {
				ports += p + " "
			}
			activeWorkloads[k] = append(activeWorkloads[k], ports)
			k++
			continue
		}
		workloadsInactive = append(workloadsInactive, "//"+s.GetName())
		// PrintActiveWorkload(c, s, config, k*10)
	}

	defaultAnimation := at.Animation{
		AnimationType: at.Ikea,
		Duration:      1000 * c.App.Metadata["animation"].(int64),
		Direction:     at.Right,
		GradientV:     true,
		GradientH:     true,
	}

	regs := []string{}
	for regName, reg := range config.GetRegistries() {
		if reg != "" {
			regs = append(regs, "["+regName+"] -> "+reg)
		}
	}

	ch, wg := c.App.Metadata["iamui"].(at.IUserInterface).StartDrawLoop(90)

	tempOffset = userInterface.DrawElement(at.CreatePos(0, 25), "Registries:", at.GREY)

	tempOffset = userInterface.DrawElementsHorizontal(
		at.CreatePos(0, 0).SetOffset(getGlobalOffset(tempOffset)+1),
		regs,
		[]int{0, 50},
		[]int{at.BLUE, at.YELLOW, at.COLORPATTERNSKYLIGHT, at.PINK, at.GREY})

	tempOffset = userInterface.DrawPattern(
		at.CreatePos(0, 0).SetOffset(getGlobalOffset(tempOffset)+1),
		100,
		"\\",
		at.COLORPATTERNLIME,
		defaultAnimation)

	tempOffset = userInterface.DrawElementsHorizontal(
		at.CreatePos(0, 0).SetOffset(getGlobalOffset(tempOffset)+1),
		[]string{"Workload", "Status", "Image", "Tags", "Ports"},
		[]int{0, 25, 35, 70, 90},
		[]int{at.BLUE, at.YELLOW, at.COLORPATTERNSKYLIGHT, at.PINK, at.GREY})

	tempOffset = userInterface.DrawPattern(
		at.CreatePos(0, 0).SetOffset(getGlobalOffset(tempOffset)+1),
		100,
		"─",
		at.COLORPATTERNLIME,
		defaultAnimation)

	tempOffset = userInterface.DrawTable(
		at.CreatePos(0, 0).SetOffset(getGlobalOffset(tempOffset)+1),
		activeWorkloads,
		[]int{0, 25, 35, 70, 90},
		[]int{at.BLUE, at.YELLOW, at.COLORPATTERNSKYLIGHT, at.PINK, at.GREY})

	tempOffset = userInterface.DrawPattern(
		at.CreatePos(0, 0).SetOffset(getGlobalOffset(tempOffset)+1),
		100,
		"/",
		at.COLORPATTERNMEADOWS1,
		defaultAnimation)

	// tempOffset = userInterface.DrawElement(
	// 	at.CreatePos(0, 0).SetOffset(getGlobalOffset(tempOffset)),
	// 	"Inactive workloads:",
	// 	at.YELLOW)

	tempOffset = userInterface.DrawElementsHorizontal(
		at.CreatePos(0, 0).SetOffset(getGlobalOffset(tempOffset)+1),
		append([]string{"Inactive workloads: "}, workloadsInactive...),
		[]int{0, 20, 40, 60, 80},
		[]int{at.LIGHTGREY, at.LIGHTGREY, at.LIGHTGREY, at.LIGHTGREY, at.LIGHTGREY})

	// userInterface.DrawElement(at.CreatePos(0, 62), workloads[1], at.COLORPATTERNNEON1)
	// userInterface.DrawPattern(at.CreatePos(0, 55), 100, "─", at.COLORPATTERNLIME, defaultAnimation)
	time.Sleep(time.Duration(100) * time.Millisecond)
	close(ch)
	wg.Wait()
}

func getGlobalOffset(tempOffset int) int {
	if tempOffset != 0 {
		globalOffset = tempOffset
	}

	return globalOffset
}

// PrintActiveWorkload ...
func PrintActiveWorkload(c *cli.Context, w I.IWorkload, config I.IConfigObject, line int) {
	// userInterface := c.App.Metadata["iamui"].(at.IUserInterface)

	// ch, wg := c.App.Metadata["iamui"].(at.IUserInterface).StartDrawLoop(80)
	y := at.Width()
	_, _, activeBool := w.GetActive()

	if activeBool {
		// y := at.Width()
		imageAndTag := ReplaceRegistries(w.GetPod().GetMainContainer().GetImage().GetFullName(), config.GetRegistries(), w.GetName())

		yth := y / 8
		fmt.Printf("    ")
		fmt.Printf("%-*v", yth*2+colorOffset, at.Color(w.GetName(), at.BLUE))
		fmt.Printf("%-*v", yth*1, "down")
		fmt.Printf("%-*v", yth*2+colorOffset, imageAndTag["Image"])
		fmt.Printf("%-*v", yth*2, imageAndTag["Tag"])
		for _, p := range w.GetPod().GetMainContainer().GetPortsAsString() {
			fmt.Printf("%v ", p)
		}
		fmt.Println()

		PrintNetworkDetails(c, w.GetPod().GetMainContainer(), 4)
		PrintVolumeDetails(c, w)
		PrintSidecars(c, w, config)
	}
}

// ReplaceRegistries ...
func ReplaceRegistries(imageFullName string, regs map[string]string, workload string) map[string]string {
	workloadUPPER := strings.ToUpper("IMAGE_" + workload)
	workloadUPPER = strings.ReplaceAll(workloadUPPER, "-", "_")

	imageIsOverwritten := ""
	val, found := os.LookupEnv(workloadUPPER)
	if found && val != "" {
		imageIsOverwritten = " - (overwritten)"
		imageFullName = val
	}

	imageParts := strings.Split(imageFullName, ":")
	if len(imageParts) != 2 {
		fmt.Println(at.Color("There seems to be an error with this image: "+imageFullName+", is the Tag missing?", at.RED))
		imageParts = append(imageParts, "NO TAG!!!")
	}

	imageParts[0] += imageIsOverwritten

	foundReg := false
	for regName, reg := range regs {
		if reg != "" && strings.Contains(imageParts[0], reg) {
			imageParts[0] = strings.ReplaceAll(imageParts[0], reg, at.Color("["+regName+"]", at.GREEN))
			foundReg = true
		}
	}
	if !foundReg {
		imageParts[0] = at.Color(imageParts[0], at.GREEN)
	}
	return map[string]string{"Image": imageParts[0], "Tag": imageParts[1]}
}

// PrintSidecars ...
func PrintSidecars(c *cli.Context, s I.IWorkload, config I.IConfigObject) {
	y := at.Width()
	if c.Bool("sidecars") {
		if len(s.GetPod().GetSidecars()) > 0 {
			PrintFullScreenCenteredString("   └> Sidecars: ", "", "")
		}
		for _, sc := range s.GetPod().GetSidecars() {
			imageAndTag := ReplaceRegistries(sc.GetImage().GetFullName(), config.GetRegistries(), sc.GetName())
			yth := y / 8
			fmt.Printf("     ")
			fmt.Printf("%-*v", yth*2+colorOffset, at.Color(sc.GetName(), at.BLUE))
			fmt.Printf("%-*v", yth*1, "down")
			fmt.Printf("%-*v", yth*2+colorOffset, imageAndTag["Image"])
			fmt.Printf("%-*v", yth*2, imageAndTag["Tag"])
			for _, p := range sc.GetPortsAsString() {
				fmt.Printf("%-*v", yth, p)
			}
			fmt.Println()

			// fmt.Printf(
			// 	"        └> %-28v│ %-13v│ %-42v %-21v %-2v \n",
			// 	at.Color(sc.GetName(), at.CYAN),
			// 	"TODO",
			// 	imageAndTag["Image"],
			// 	imageAndTag["Tag"],
			// 	sc.GetPortsAsString(),
			// )
			PrintNetworkDetails(c, sc, 11)
		}
	}
}

// PrintNetworkDetails ...
func PrintNetworkDetails(c *cli.Context, s I.IContainer, indent int) {
	if c.Bool("network") {
		fmt.Println("TODO Network")
		// aliases := s.Get Networks["default"]["aliases"]
		// if len(aliases) > 1 {
		// fmt.Println(strings.Repeat(" ", indent) + "└> " + strings.Join(aliases[1:], ",\n"+strings.Repeat(" ", indent)+"└> "))
		// }
	}
}

// PrintVolumeDetails ...
func PrintVolumeDetails(c *cli.Context, w I.IWorkload) {
	if len(w.GetPod().GetMainContainer().GetVolumes()) > 0 && c.Bool("mounts") {
		PrintFullScreenCenteredString("     └Volumes: ", "", "")
		for _, vol := range w.GetPod().GetMainContainer().GetVolumes() {
			fmt.Println(vol)
			// hostContainerVol := strings.Split(vol, ":")

			// if strings.HasPrefix(hostContainerVol[0], "..") || strings.HasPrefix(hostContainerVol[0], "/") {
			// 	// hostpath
			// 	hostpath := filepath.Join(w.GetPath(), hostContainerVol[0])

			// 	PrintFullScreenColored([]string{"      └>: " + hostpath, " -> " + hostContainerVol[1]}, []int{at.BLUE, at.CYAN})
			// } else {
			// 	// docker volume
			// 	PrintFullScreenColored([]string{"      └>: " + vol}, []int{at.RED})
			// }
		}
	}
}

// PrintInactiveWorkload ...
func PrintInactiveWorkload(c *cli.Context, s I.IWorkload) {
	_, _, activeBool := s.GetActive()
	if !activeBool {
		fmt.Print(at.Color(" //"+s.GetName(), at.GREEN))
	}
}

// PrintExecutionTime ...
func PrintExecutionTime(d time.Duration) {
	fmt.Print(at.Color("Execution took: ", at.RED))
	fmt.Println(d.Round(time.Millisecond))
}

func sortMapAlphabetically(WorkloadMap map[string]I.IWorkload) []I.IWorkload {
	sorted := []I.IWorkload{}
	for _, v := range WorkloadMap {
		sorted = append(sorted, v)
	}
	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i].GetName() < sorted[j].GetName()
	})
	return sorted
}

// CheckNewSkill ...
func CheckNewSkill(c *cli.Context) error {
	if c.Command.FullName() != "init" && !c.App.Metadata["iamconfig"].(I.IConfigObject).WasCommandUsed("init") {
		log.Fatal("Before doing anything else you should run the init command.")
	}
	if c.App.Metadata["iamconfig"].(I.IConfigObject).WasCommandUsed(c.Command.FullName()) {
		return nil
	}
	c.App.Metadata["iamconfig"].(I.IConfigObject).MarkCommandLerned(c.Command.FullName())

	totalNrOfCommands := len(c.App.Commands)
	commandsLearned := c.App.Metadata["iamconfig"].(I.IConfigObject).LearnedCommands()

	// var wg sync.WaitGroup
	// wg.Add(1)
	// ch := make(chan int)
	c.App.Metadata["iamui"].(at.IUserInterface).ClearScreen()

	ch, wg := c.App.Metadata["iamui"].(at.IUserInterface).StartDrawLoop(95)
	// Draw(at.ReducedHeight(), at.Width(), ch, &wg)

	c.App.Metadata["iamui"].(at.IUserInterface).MoveElement(
		at.CreatePos(30, 5),
		at.CreatePos(0, 5),
		figure.NewFigure("New Command:", "slant", true).String(),
		at.COLORPATTERNNEON1,
		at.Animation{
			AnimationType: at.Ikea,
			Duration:      0,
		})

	go c.App.Metadata["iamui"].(at.IUserInterface).MoveElement(
		at.CreatePos(90, 2),
		at.CreatePos(80, 2),
		"Level "+strconv.Itoa(commandsLearned)+"/"+strconv.Itoa(totalNrOfCommands),
		at.COLORPATTERNNEON1,
		at.Animation{
			AnimationType: at.Ikea,
			Duration:      1000,
		})

	c.App.Metadata["iamui"].(at.IUserInterface).DrawPattern(
		at.CreatePos(100, 20), 100, "█\n█\n█\n█\n█\n", at.COLORPATTERNNEON1,
		at.Animation{
			AnimationType: at.Ikea,
			Duration:      900,
			Direction:     at.Left,
			GradientV:     true,
			GradientH:     true,
		})

	// PrintFullScreenRepeatingChar("┌", "─", "┐", true)
	c.App.Metadata["iamui"].(at.IUserInterface).MoveElement(
		at.CreatePos(100, 20),
		at.CreatePos(0, 20),
		figure.NewFigure(c.Command.FullName(), "standard", true).String(),
		at.COLORPATTERNSKYLIGHT,
		at.Animation{
			AnimationType: at.Ikea,
			Duration:      1000,
		})

	// PrintFullScreenRepeatingChar("┌", "─", "┐", true)
	c.App.Metadata["iamui"].(at.IUserInterface).MoveElement(
		at.CreatePos(70, 45),
		at.CreatePos(2, 45),
		c.Command.Description,
		at.COLORPATTERNSKYLIGHT,
		at.Animation{
			AnimationType: at.Ikea,
			Duration:      1000,
		})

	close(ch)
	wg.Wait()
	os.Exit(0)

	return nil
}
