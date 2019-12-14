package ui

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/gregorpirolt/iamutils"

	"github.com/urfave/cli"

	tm "github.com/buger/goterm"
	"github.com/common-nighthawk/go-figure"
)

// PrintCurrentStats ...
func PrintCurrentStats() error {
	//0= month?, 1= month, 2 == day, 3 == h; 4 == min 5 == secs;
	// pl := fmt.Println
	// p := fmt.Print
	// t := time.Now()
	// p(dockerstate)
	return nil
}

// PrintPercentOfScreen ...
func PrintPercentOfScreen(percent int, c string) {
	p := fmt.Print
	y := tm.Width()
	y = int(y*percent) / 100

	stringBuffer := strings.Repeat(c, y)
	p(stringBuffer)
}

// PrintFullScreenRepeatingChar ...
func PrintFullScreenRepeatingChar(left string, c string, right string, singleCharLeftRight bool) {
	p := fmt.Println
	y := tm.Width()
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
	y := tm.Width()
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
		stringline += tm.Color(parts[k], colors[k])
	}

	p(stringline)
}

// PrintBanner ...
func PrintBanner(c *cli.Context) {

	cmd := exec.Command("clear", "cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()

	PrintFullScreenRepeatingChar("┌", "─", "┐", true)
	
	iamASCII := figure.NewFigure(c.App.Name, "standard", true)
	iamconfig := c.App.Metadata["iamconfig"].(iamutils.IamConfigYaml)
	properties := []string{"go edition", "v" + c.App.Version, iamconfig.IamDir, "Todo", "Connected Cluster: " + c.App.Metadata["currentcontext"].(string),"",""}

	for k := range iamASCII.Slicify() {
		PrintFullScreenColored([]string{"\t", iamASCII.Slicify()[k] + "\t", properties[k]}, []int{tm.RED, tm.RED, tm.CYAN})
	}
	PrintFullScreenRepeatingChar("└", "─", "┘", true)
}

// PrintServiceOverview ...
func PrintServiceOverview(c *cli.Context) {
	y := tm.Width()
	serviceMap := c.App.Metadata["services"].(map[string]iamutils.CliService)
	// make array out of services, since they are not modifies after printing here and need to be sorted alphabetically
	sortedServices := sortMapAlphabetically(serviceMap)
	config := c.App.Metadata["iamconfig"].(iamutils.IamConfigYaml)

	regs := "Registries:"
	for regName, reg := range config.Registries {
		if reg != "" {
			regs = strings.Join([]string{regs, " ", "[" + tm.Color(regName, tm.GREEN) + "]", "->", reg}, "")
		}
	}
	fmt.Println("  " + regs)
	PrintFullScreenRepeatingChar("┌", "─", "┐", true)

	yth := y / 20
	fmt.Printf("  %-*v", yth*4, "Service:")
	fmt.Printf("|%-*v|", yth*2, "Status:")
	fmt.Printf("%-*v", yth*5, "Image:")
	fmt.Printf("%-*v", yth*4, "Tag:")
	fmt.Printf("%-*v\n", yth*4, "Ports: [host->container]")
	PrintFullScreenRepeatingChar("│", "─", "│", true)
	for _, s := range sortedServices {
		PrintActiveService(c, s, config)
	}

	fmt.Println()
	fmt.Println("Inactive services:")
	for _, s := range sortedServices {
		PrintInactiveService(c, s)
	}
	fmt.Println()
	PrintFullScreenRepeatingChar("└", "─", "┘", true)
}

// PrintActiveService ...
func PrintActiveService(c *cli.Context, s iamutils.CliService, config iamutils.IamConfigYaml) {
	_, _, activeBool := s.GetActive()
	if activeBool {
		y := tm.Width()

		imageAndTag := ReplaceRegistries(s.GetMaster().Image, config.Registries, s.GetName())

		yth := y / 20
		col := 9
		fmt.Printf("  %-*v", yth*4+col, tm.Color(s.GetName(), tm.RED))
		fmt.Printf("|%-*v|", yth*2, " down ")
		fmt.Printf("%-*v", yth*5+col, imageAndTag["Image"])
		fmt.Printf("%-*v", yth*4,imageAndTag["Tag"])
		fmt.Printf("%-*v\n", yth*1, s.GetMaster().Ports)

		PrintNetworkDetails(c, s.GetMaster(), 4)
		PrintVolumeDetails(c, s)
		PrintSidecars(c, s, config)
	}
}

// ReplaceRegistries ...
func ReplaceRegistries(imageFullName string, regs map[string]string, service string) map[string]string {
	serviceUPPER := strings.ToUpper("IMAGE_" + service)
	serviceUPPER = strings.ReplaceAll(serviceUPPER, "-", "_")

	imageIsOverwritten := ""
	val, found := os.LookupEnv(serviceUPPER)
	if found && val != "" {
		imageIsOverwritten = " - (overwritten)"
	}

	imageParts := strings.Split(imageFullName, ":")
	if len(imageParts) != 2 {
		fmt.Println(tm.Color("There seems to be an error with this image: "+imageFullName+", is the Tag missing?", tm.RED))
		imageParts = append(imageParts, "NO TAG!!!")
	}

	imageParts[0] += imageIsOverwritten

	foundReg := false
	for regName, reg := range regs {
		if reg != "" && strings.Contains(imageParts[0], reg) {
			imageParts[0] = strings.ReplaceAll(imageParts[0], reg, tm.Color("["+regName+"]", tm.GREEN))
			foundReg = true
		}
	}
	if !foundReg {
		imageParts[0] = tm.Color(imageParts[0], tm.YELLOW)
	}
	return map[string]string{"Image": imageParts[0], "Tag": imageParts[1]}
}

// PrintSidecars ...
func PrintSidecars(c *cli.Context, s iamutils.CliService, config iamutils.IamConfigYaml) {
	if c.Bool("sidecars") {
		if len(s.GetSidecars()) > 0 {
			PrintFullScreenCenteredString("     └> Sidecars: ", "", "")
		}
		for scName, sc := range s.GetSidecars() {
			imageAndTag := ReplaceRegistries(sc.Image, config.Registries, scName)
			fmt.Printf(
				"        └> %-36v│ %-2v│ %-45v %-25v %s\n",
				tm.Color(scName, tm.CYAN),
				s.Setting.GetEnvAsEmoji(),
				imageAndTag["Image"],
				imageAndTag["Tag"],
				sc.Ports,
			)
			PrintNetworkDetails(c, sc, 11)
		}
	}
}

// PrintNetworkDetails ...
func PrintNetworkDetails(c *cli.Context, s iamutils.DockerComposeContainer, indent int) {
	if c.Bool("network") {
		aliases := s.Networks["default"]["aliases"]
		if len(aliases) > 1 {
			fmt.Println(strings.Repeat(" ", indent) + "└> " + strings.Join(aliases[1:], ",\n"+strings.Repeat(" ", indent)+"└> "))
		}
	}
}

// PrintVolumeDetails ...
func PrintVolumeDetails(c *cli.Context, s iamutils.CliService) {
	if len(s.GetMaster().Volumes) > 0 && c.Bool("mounts") {
		PrintFullScreenCenteredString("     └Volumes: ", "", "")
		for _, vol := range s.GetMaster().Volumes {
			hostContainerVol := strings.Split(vol, ":")

			if strings.HasPrefix(hostContainerVol[0], "..") || strings.HasPrefix(hostContainerVol[0], "/") {
				// hostpath
				hostpath := filepath.Join(s.Path, hostContainerVol[0])

				PrintFullScreenColored([]string{"      └>: " + hostpath, " -> " + hostContainerVol[1]}, []int{tm.BLUE, tm.CYAN})
			} else {
				// docker volume
				PrintFullScreenColored([]string{"      └>: " + vol}, []int{tm.RED})
			}
		}
	}
}

// PrintInactiveService ...
func PrintInactiveService(c *cli.Context, s iamutils.CliService) {
	_, _, activeBool := s.GetActive()
	if !activeBool {
		fmt.Print(tm.Color(" //"+s.GetName(), tm.GREEN))
	}
}

// PrintExecutionTime ...
func PrintExecutionTime(d time.Duration) {
	fmt.Print(tm.Color("Execution took: ", tm.RED))
	// fmt.Println(tm.Color(fmt.Sprint(d.Round(time.Millisecond)), tm.CYAN))
	fmt.Println(d.Round(time.Millisecond))

}

func sortMapAlphabetically(serviceMap map[string]iamutils.CliService) []iamutils.CliService {
	sorted := []iamutils.CliService{}
	for _, v := range serviceMap {
		sorted = append(sorted, v)
	}
	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i].Name < sorted[j].Name
	})

	return sorted
}

// ┌ ┐ └ ┘ ─ │
