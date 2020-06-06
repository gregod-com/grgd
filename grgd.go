// This package implements the cli for the iam-stack. The underlying framework is depnedent upon urfave/cli.
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"plugin"
	"sort"
	"time"
	"reflect"

	I "github.com/gregod-com/interfaces"

	A "github.com/gregod-com/grgd/actions"
	PlugIndex "github.com/gregod-com/grgd/pluginindex"
	T "github.com/gregod-com/grgd/templates"
	Impl "github.com/gregod-com/implementations"

	"github.com/urfave/cli/v2"
)

const PLUGINSKEY = "grgdplugins"
const STARTTIMEKEY = "startTime"
const CONFIGPATH = "configLocation"
const CONFIG = "iamconfig"

var UI I.IUI

// Entrypoint for the iam cli
func main() {
	myFlags := []cli.Flag{
		&cli.BoolFlag{
			Name:  "debug, d",
			Usage: "run the command in debug mode",
		},
		&cli.BoolFlag{
			Name:  "network, n",
			Usage: "show network details in overview",
		},
		&cli.BoolFlag{
			Name:  "mounts, m",
			Usage: "show mounting details in overview",
		},
		&cli.BoolFlag{
			Name:  "sidecars, s, sc",
			Usage: "show sidecars in overview",
		},
	}
	catStackControlls := "1) - stack controls"
	catWorkloadSpecific := "2) - workload specific"
	catUtils := "3) - utilities"

	app := cli.NewApp()
	app.Flags = myFlags
	app.Name = "grgd"
	app.Usage = "written in go. Can be used as a sidekick to gregod-menu and gregod-doctor"
	app.Version = "0.6.0"
	app.Metadata = make(map[string]interface{})
	app.Metadata[STARTTIMEKEY] = time.Now()
	app.CustomAppHelpTemplate = T.GetHelpTemplate()
	app.Before = func(c *cli.Context) error {
		userHome, err := os.UserHomeDir()
		if err != nil {
			return nil
		}

		c.App.Metadata["repoIndex"] = "https://s3.gregod.com/public/plugins/index.yaml"
		c.App.Metadata["updatecheckinterval"] = time.Millisecond * 200
		c.App.Metadata[CONFIGPATH] = userHome + "/.iam/iam_conf.yml"
		c.App.Metadata["currentcontext"] = A.UtilGetCurrentKubeContext()
		c.App.Metadata[CONFIG] = Impl.CreateConfigObjectYaml(c.App.Metadata[CONFIGPATH].(string))
		c.App.Metadata["animation"] = int64(300)

		SystemCheck(c)
		LoadPlugins(c)

		UI.ClearScreen(c)

		UI.PrintBanner(c)
		// UI.PrintPercentOfScreen(c, "h", 20)
		UI.PrintWorkloadOverview(c)
		return nil
	}
	app.Action = func(c *cli.Context) error {
		// UI.PrintWorkloadOverview(c)
		return nil
	}
	app.After = func(c *cli.Context) error {
		c.App.Metadata[CONFIG].(I.IConfigObject).Update()
		// startTime := c.App.Metadata[STARTTIMEKEY].(time.Time)
		// fmt.Println(startTime)
		// UI.PrintExecutionTime(time.Since(startTime))
		return nil
	}
	app.Commands = []*cli.Command{
		{
			Name:     "init",
			Category: catUtils,
			Usage:    "Initialze the " + app.Name,
			// Before:   UI.CheckNewSkill,
			Action: A.AInit,
			Description: `
			Welcome to the ` + app.Name + ` !!! 🍄 🍄 🍄

			It looks like you just unlocked your first command. 🤗 🎉 🎉 🎉
			Sadly you are not going to use this one as often as the other ones.
			But still it is an important one.

			When ever you are ready, lets start setting up the ` + app.Name + ` by defining
			the base path for your projects:
			`,
		},
		{
			Name:     "up",
			Category: catStackControlls,
			Usage:    "Start dev stack with current config",
			Aliases:  []string{"u"},
			Flags:    myFlags,
			// Before:   UI.CheckNewSkill,
			Action: A.AUp,
			Description: `
The 'up' command allows you to start a single workload in your stack or even the
whole stack at once. All services that are currently active can be started with the command:
iam up [workload name] (i.e iam up database)
This does not nessesarily mean that the workload is actually started as a container
but that 'a' workload is made available to your stack via it's DNS name.
You can for example start a workload like a database and just wire up the connection to
an external database hosted as a process on your local machine, on a nearby dev server, or even
with tunneling or ingress routing on a remote kubernetes cluster. If you omitt the workload name
the cli assumes you want to start all defined workloads. If some or all of them are already running,
the command is ignored for them. If you need to restart a service have a look at iam restart.
			`,
		},
		{
			Name:     "down",
			Category: catStackControlls,
			Usage:    "Stop dev stack (or single workload)",
			Aliases:  []string{"d"},
			Flags:    myFlags,
			// Before:   UI.CheckNewSkill,
			Action: A.ADown,
			Description: `
			The 'down' command is good and professional.
			`,
		},
		{
			Name:     "restart",
			Category: catStackControlls,
			Usage:    "Restart dev stack (or single workload)",
			Aliases:  []string{"r"},
			Flags:    myFlags,
			// Before:   UI.CheckNewSkill,
			Action: A.ARestart,
			Description: `
			The 'restart' command is good and professional.
			`,
		},
		{
			Name:      "logs",
			Category:  catStackControlls,
			Usage:     "Show all logs for running stack (or single workload)",
			Aliases:   []string{"l"},
			Flags:     myFlags,
			UsageText: "Show logs for a single Workload or all workloads combined",
			ArgsUsage: "Args usage",
			// Before:    UI.CheckNewSkill,
			Action: A.ALogs,
			Description: `
			The 'logs' command is good and professional.
			`,
		},
		{
			Name:     "config",
			Category: catStackControlls,
			Usage:    "Configuration for current stack (type `config help` for possible subcommands)",
			Aliases:  []string{"conf", "c", "."},
			Flags:    myFlags,
			Subcommands: []*cli.Command{
				{
					Name:    "yaml",
					Usage:   "Show iam_config.yaml file",
					Aliases: []string{"y"},
					Action:  A.SubAConfig["yaml"],
					// Before:  UI.CheckNewSkill,
					Description: `
					The 'config yaml' command is good and professional.
					`,
				},
			},
		},
		{
			Name:     "enter",
			Category: catWorkloadSpecific,
			Usage:    "Enter a workload",
			Aliases:  []string{"en"},
			Flags:    myFlags,
			// Before:   UI.CheckNewSkill,
			Action: A.AEnter,
			Description: `
			The 'enter' command is good and professional.
			`,
		},
		{
			Name:     "execute",
			Category: catWorkloadSpecific,
			Usage:    "Execute a command in workload and view output",
			Aliases:  []string{"exec", "ex"},
			Flags:    myFlags,
			// Before:   UI.CheckNewSkill,
			Action: A.AExecute,
			Description: `
			The 'execute' command is good and professional.
			`,
		},
		{
			Name:      "test",
			Category:  catWorkloadSpecific,
			Usage:     "Run unittest inside container",
			Aliases:   []string{"t"},
			Flags:     myFlags,
			UsageText: "TODO",
			ArgsUsage: "TODO",
			// Before:    UI.CheckNewSkill,
			Action: A.ATest,
			Description: `
			The 'test' command is good and professional.
			`,
		},
		{
			Name:      "activate",
			Category:  catWorkloadSpecific,
			Usage:     "Activate a workload",
			Aliases:   []string{"act", "a"},
			Flags:     myFlags,
			UsageText: "TODO",
			ArgsUsage: "TODO",
			// Before:    UI.CheckNewSkill,
			Action: A.AActivate,
			After: func(c *cli.Context) error {
				// UI.PrintWorkloadOverview(c)
				return nil
			},
			Description: `
			The 'activate' command is good and professional.
			`,
		},
		{
			Name:     "shortcuts",
			Category: catUtils,
			Usage:    "Show and edit shortcut names for workloads",
			Aliases:  []string{"sc", "shortcut"},
			Flags:    myFlags,
			Subcommands: []*cli.Command{
				{
					Name:    "list",
					Aliases: []string{"l", "ls"},
					Usage:   "list all shortcuts",
					Action:  A.SubAShortcut["list"],
					// Before:  UI.CheckNewSkill,
					Description: `
					The 'shortcuts list' command is good and professional.
					`,
				},
				{
					Name:    "add",
					Aliases: []string{"a"},
					Usage:   "add new shortcut `sc add shortcut workload` ",
					Action:  A.SubAShortcut["add"],
					// Before:  UI.CheckNewSkill,
					Description: `
					The 'shortcuts add' command is good and professional.
					`,
				},
				{
					Name:    "remove",
					Aliases: []string{"r"},
					Usage:   "remove a shortcut `ssc remove shortcut` ",
					Action:  A.SubAShortcut["remove"],
					// Before:  UI.CheckNewSkill,
					Description: `
					The 'shortcuts remove' command is good and professional.
					`,
				},
			},
			After: func(c *cli.Context) error {
				A.PrintShortcuts(c)
				return nil
			},
		},
		{
			Name:     "volume",
			Category: catUtils,
			Usage:    "View and edit workload attached volumes",
			Aliases:  []string{"vol"},
			Flags: append(
				myFlags,
				&cli.BoolFlag{
					Name:  "print_volume",
					Usage: "run the command without coloring output",
				},
			),
			UsageText: "TODO",
			ArgsUsage: "TODO",
			// Before:    UI.CheckNewSkill,
			Action: A.AVolume,
			Description: `
			The 'volume' command is good and professional.
			`,
		},
		{
			Name:      "certificates",
			Category:  catUtils,
			Usage:     "View and fetch certificates from cluster",
			Aliases:   []string{"cert"},
			UsageText: "TODO",
			ArgsUsage: "TODO",
			// Before:    UI.CheckNewSkill,
			Action: A.ACertificates,
			Description: `
			The 'certificates' command is good and professional.
			`,
		},
		{
			Name:     "dns",
			Category: catUtils,
			Usage:    "View and edit DNS routing",
			// Before:   UI.CheckNewSkill,
			Action: A.ADNS,
			Description: `
			The 'dns' command is good and professional.
			`,
		},
		{
			Name:     "context",
			Category: catUtils,
			Usage:    "View and edit kubernetes context",
			Aliases:  []string{"cont", "kubec"},
			// Before:   UI.CheckNewSkill,
			Action: A.AContext,
			After:  A.AfterContext,
			Description: `
			The 'context' command is good and professional.
			`,
		},
		{
			Name:     "helm",
			Category: catUtils,
			Usage:    "View and edit helm deployments",
			Aliases:  []string{"he"},
			// Before:   UI.CheckNewSkill,
			Action: A.AHelm,
			Description: `
			The 'helm' command is good and professional.
			`,
		},
	}

	sort.Sort(cli.FlagsByName(app.Flags))
	// c.App.Metadata["iamui"].(at.IUserInterface).ClearScreen()
	// c.App.Metadata["iamui"].(at.IUserInterface).SetBoarder(0)
	// go c.App.Metadata["iamui"].(at.IUserInterface).Draw(at.ReducedHeight()/3, at.Width())

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

// SystemCheck ...
func SystemCheck(ctx *cli.Context) error {
	homedir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	path := homedir + "/.grgd"
	if _, notexistserr := os.Stat(path); os.IsNotExist(notexistserr) {
		os.Mkdir(path, os.FileMode(uint32(0760)))
	}
	ctx.App.Metadata["grgdhome"] = path

	pluginPath := homedir + "/.grgd/plugins"
	if _, notexistserr := os.Stat(pluginPath); os.IsNotExist(notexistserr) {
		os.Mkdir(pluginPath, os.FileMode(uint32(0760)))
	}
	ctx.App.Metadata[PLUGINSKEY] = pluginPath
	updateCheckInterval := ctx.App.Metadata["updatecheckinterval"].(time.Duration)
	pl := PlugIndex.CreatePluginIndex(pluginPath + "/index.yaml")
	if time.Now().After(pl.GetLastChecked().Add(updateCheckInterval)) {
		err := CheckUpdate(ctx)
		if err != nil {
			log.Println("Looks like there was an error fetching updates... skipping this update-cycle")
			pl.Update()
		}
	}

	return err
}

// CheckUpdate ...
func CheckUpdate(ctx *cli.Context) error {
	// reader := bufio.NewReader(os.Stdin)
	repoIndex := ctx.App.Metadata["repoIndex"].(string)
	// currentIndex := ctx.App.Metadata[PLUGINSKEY].(string) + "/index.yaml"
	remoteIndex := ctx.App.Metadata[PLUGINSKEY].(string) + "/index-remote.yaml"
	err := DownloadFile(remoteIndex, repoIndex)
	if err != nil {
		return err
	}
	// pluginsCurrent := PlugIndex.CreatePluginIndex(currentIndex)
	// pluginsRemote := PlugIndex.CreatePluginIndex(remoteIndex)

	// fmt.Println("Downloaded: " + repoIndex)
	// for _, plremote := range pluginsRemote.GetPluginList().([]PlugIndex.PluginMetadata) {
	// 	for _, pllocal := range pluginsCurrent.GetPluginList().([]PlugIndex.PluginMetadata) {
	// 		if plremote.Name == pllocal.Name {
	// 			vlocal := semver.New(pllocal.Version)
	// 			vremote := semver.New(plremote.Version)
	// 			if vlocal.LessThan(*vremote) {
	// 				fmt.Printf("Update plugin    %-15s to v%v (current %v)? [y/n]", plremote.Name, vremote, vlocal)
	// 				yes, _ := reader.ReadString('\n')
	// 				if yes == "y\n" {
	// 					fmt.Println(plremote.Sha)
	// 					DownloadFile("filepath", plremote.URL)
	// 				}
	// 			}
	// 		}
	// 	}

	// 	// fmt.Println(p.Name)
	// 	// fmt.Println(p.Version)
	// 	// fmt.Println(p.Size)
	// 	// fmt.Println(p.Sha)
	// 	// fmt.Println(p.URL)
	// }
	// pl.Update()
	return nil
}

// DownloadFile updates and plugins from s3
func DownloadFile(filepath string, url string) error {

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}

// LoadPlugins ...
func LoadPlugins(ctx *cli.Context) error {

	// TODO: the categories should be fetched from config
	var categories = []string{"test", "ui", "executor", "dns"}

	pluginFolder, ok := ctx.App.Metadata[PLUGINSKEY].(string)
	if !ok {
		log.Fatal(ok)
	}

	// iterate over cartegories in plugin folder
	for _, c := range categories {
		fileinfo, err := ioutil.ReadDir(pluginFolder + "/" + c)
		if err != nil {
			fmt.Println(err)
		}

		// iterate over plugin implementations
		for _, f := range fileinfo {
			pluginPath := pluginFolder + "/" + c + "/" + f.Name()
			if ctx.Bool("debug") {
				log.Printf("Loading plugin at %s\n", pluginPath)
			}

			// open .so file and error if something goes wrong
			pluginImpl, err := plugin.Open(pluginPath)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			// check if there is a var of func called `Plugin` in .so file
			symPlugin, err := pluginImpl.Lookup("Plugin")
			if err != nil {
				fmt.Println(err)
				os.Exit(2)
			}

			// check if the var/func is implemneting the grgd plugin interface
			grgdplugin, ok := symPlugin.(I.IGrgdPlugin)
			if !ok {
				fmt.Println(reflect.TypeOf(symPlugin))
				cli.Exit("unexpected type from module symbol", 3)
			}

			// check if the init method returns a valid plugin metatda interface
			metadata, ok := grgdplugin.Init(pluginPath).(I.IPluginMetadata)
			if !ok {
				fmt.Println("Looks like the metadata is shiat")
				os.Exit(4)
			}

			ctx.App.Metadata[metadata.Name()] = grgdplugin

			uiplugin, ok := symPlugin.(I.IUI)
			if ok {
				if ctx.Bool("debug") {
					log.Printf("Assigned %s as default UI\n", metadata.Name())
				}
				UI = uiplugin
			}

		}

	}
	return nil
}
