// This package implements the cli for the iam-stack. The underlying framework is depnedent upon urfave/cli.
package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"time"

	plugContracts "github.com/gregod-com/grgdplugincontracts"
	I "github.com/gregod-com/interfaces"

	A "github.com/gregod-com/grgd/actions"
	T "github.com/gregod-com/grgd/templates"
	Impl "github.com/gregod-com/implementations"

	cli "github.com/urfave/cli/v2"
)

// PLUGINSKEY ...
const PLUGINSKEY = "grgdplugins"

// STARTTIMEKEY ...
const STARTTIMEKEY = "startTime"

// CONFIGPATH ...
const CONFIGPATH = "configLocation"

// CONFIG ...
const CONFIG = "iamconfig"

func main() {
	app := cli.NewApp()

	homedir := HomeDir()
	CMDPlugins, UIPlugin := LoadPlugins(homedir + "/.grgd/plugins/")

	for _, plug := range CMDPlugins {
		cmdplug, ok := plug.(plugContracts.ICMDPlugin)
		if !ok {
			fmt.Println("Wrong impl")
		}
		app.Commands = append(app.Commands, cmdplug.GetCommands(nil)...)
	}

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

	app.Flags = myFlags
	app.Name = "grgd"
	app.Usage = "written in go. Can be used as a sidekick to gregod-menu and gregod-doctor"
	app.Version = "0.6.0"
	app.Metadata = make(map[string]interface{})
	app.Metadata[STARTTIMEKEY] = time.Now()
	app.CustomAppHelpTemplate = T.GetHelpTemplate()
	app.Commands = append(app.Commands, []*cli.Command{
		{
			Name:        "init",
			Usage:       "Initialze the " + app.Name,
			Action:      A.AInit,
			Description: T.Description(app, "init"),
		},
		{
			Name:        "plugins",
			Usage:       "Configuration for plugins",
			Aliases:     []string{"p"},
			Flags:       myFlags,
			Description: T.Description(app, "config"),
			Subcommands: []*cli.Command{
				{
					Name:        "list",
					Usage:       "Show all plugins",
					Aliases:     []string{"ls"},
					Action:      A.APluginList,
					Description: T.Description(app, "config-yaml"),
				},
				{
					Name:        "activate",
					Usage:       "Activate plugins",
					Aliases:     []string{"a"},
					Action:      A.APluginActivate,
					Description: T.Description(app, "config-yaml"),
				},
			},
		},
		{
			Name:        "config",
			Usage:       "Configuration for current stack",
			Aliases:     []string{"conf", "c"},
			Flags:       myFlags,
			Description: T.Description(app, "config"),
			Subcommands: []*cli.Command{
				{
					Name:        "yaml",
					Usage:       "Show iam_config.yaml file",
					Aliases:     []string{"y"},
					Action:      A.SubAConfig["yaml"],
					Description: T.Description(app, "config-yaml"),
				},
			},
		},
		{
			Name:        "shortcuts",
			Usage:       "Show and edit shortcut names for workloads",
			Aliases:     []string{"sc"},
			Flags:       myFlags,
			Description: T.Description(app, "shortcuts"),
			Subcommands: []*cli.Command{
				{
					Name:        "list",
					Aliases:     []string{"l", "ls"},
					Usage:       "list all shortcuts",
					Action:      A.SubAShortcut["list"],
					Description: T.Description(app, "shortcuts-list"),
				},
				{
					Name:        "add",
					Aliases:     []string{"a"},
					Usage:       "add new shortcut `sc add shortcut workload` ",
					Action:      A.SubAShortcut["add"],
					Description: T.Description(app, "shortcuts-add"),
				},
				{
					Name:        "remove",
					Aliases:     []string{"r"},
					Usage:       "remove a shortcut `sc remove shortcut` ",
					Action:      A.SubAShortcut["remove"],
					Description: T.Description(app, "shortcuts-remove"),
				},
			},
			After: func(c *cli.Context) error {
				A.PrintShortcuts(c)
				return nil
			},
		},
	}...)

	app.Before = func(c *cli.Context) error {
		c.App.Metadata[CONFIGPATH] = homedir + "/.grgd/config.yml"
		c.App.Metadata[CONFIG] = Impl.CreateConfigObjectYaml(c.App.Metadata[CONFIGPATH].(string))
		c.App.Metadata["pluginIndex"] = homedir + "/.grgd/plugins/index.yaml"
		c.App.Metadata["repoIndex"] = "https://s3.gregod.com/public/plugins/index.yaml"
		c.App.Metadata["AWS-REGION"] = "eu-central-1"
		c.App.Metadata["updatecheckinterval"] = time.Millisecond * 50
		c.App.Metadata["currentcontext"] = A.UtilGetCurrentKubeContext()
		c.App.Metadata["UIPlugin"] = UIPlugin

		SystemCheck(c)

		UIPlugin.ClearScreen(c)
		UIPlugin.PrintBanner(c)
		return nil
	}
	// app.Action = func(c *cli.Context) error {
	// 	// UIPlugin.PrintWorkloadOverview(c)
	// 	return nil
	// }
	app.After = func(c *cli.Context) error {
		c.App.Metadata[CONFIG].(I.IConfigObject).Update()
		// startTime := c.App.Metadata[STARTTIMEKEY].(time.Time)
		// fmt.Println(startTime)
		// UIPlugin.PrintExecutionTime(time.Since(startTime))
		return nil
	}

	sort.Sort(cli.FlagsByName(app.Flags))

	apperr := app.Run(os.Args)
	if apperr != nil {
		log.Fatal(apperr)
	}
}

// HomeDir ...
func HomeDir() string {
	dir, errHomeDir := os.UserHomeDir()
	if errHomeDir != nil {
		log.Fatal(errHomeDir)
	}
	return dir
}
