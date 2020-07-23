// This package implements the cli for the iam-stack. The underlying framework is depnedent upon urfave/cli.
package main

import (
	"log"
	"os"
	"sort"
	"time"

	I "github.com/gregod-com/interfaces"

	A "github.com/gregod-com/grgd/actions"
	T "github.com/gregod-com/grgd/templates"

	cli "github.com/urfave/cli/v2"
)

// PLUGINSKEY ...
const PLUGINSKEY = "grgdplugins"

// STARTTIMEKEY ...
const STARTTIMEKEY = "startTime"

// CONFIGPATH ...
const CONFIGPATH = "configLocation"

// CONFIG ...
const CONFIG = "config"

func main() {
	app := cli.NewApp()

	homedir := HomeDir()
	CMDPlugins, UIPlugin := LoadPlugins(homedir + "/.grgd/plugins/")

	for _, plug := range CMDPlugins {
		app.Commands = append(app.Commands, plug.GetCommands(nil)...)
	}

	myFlags := []cli.Flag{
		&cli.BoolFlag{
			Name:    "debug",
			Aliases: []string{"d"},
			Usage:   "run the command in debug mode",
		},
		&cli.BoolFlag{
			Name:    "verbose",
			Aliases: []string{"vvv"},
			Usage:   "run the command in verbose mode",
		},
		&cli.BoolFlag{
			Name:    "silent",
			Aliases: []string{"s"},
			Usage:   "mute all outputs",
		},
	}

	app.Flags = myFlags
	app.Name = "grgd"
	app.Usage = "written in go. Can be used as a sidekick to gregod-menu and gregod-doctor"
	app.Version = "0.7.0"
	app.Metadata = make(map[string]interface{})
	app.Metadata[STARTTIMEKEY] = time.Now()
	app.CustomAppHelpTemplate = T.GetHelpTemplate()
	app.HideHelpCommand = true
	app.Commands = append(app.Commands, []*cli.Command{
		{
			Name:        "init",
			Usage:       "Initialze the " + app.Name,
			Flags:       myFlags,
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
			Name:            "config",
			Usage:           "view and edit current configuration",
			Aliases:         []string{"conf", "c"},
			Flags:           app.Flags,
			HideHelpCommand: true,
			// Action:      A.AConfig,
			// Description: A.AConfigDescription,
			Subcommands: []*cli.Command{
				{
					Name:    "yaml",
					Usage:   "print config file in yaml format",
					Aliases: []string{"y"},
					Action:  A.SubAConfigYAML,
				},
				{
					Name:    "edit",
					Usage:   "edit the config file",
					Aliases: []string{"e"},
					Action:  A.SubAConfigEdit,
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
					Aliases:     []string{"ls"},
					Usage:       "list all shortcuts",
					Action:      A.SubAShortcutList,
					Description: A.SubAShortcutListDescription,
				},
				{
					Name:        "add",
					Aliases:     []string{"a"},
					Usage:       "add new shortcut `sc add shortcut workload` ",
					Action:      A.SubAShortcutAdd,
					Description: A.SubAShortcutAddDescription,
				},
				{
					Name:        "remove",
					Aliases:     []string{"r"},
					Usage:       "remove a shortcut `sc remove shortcut` ",
					Action:      A.SubAShortcutRemove,
					Description: A.SubAShortcutRemoveDescription,
				},
			},
			After: func(c *cli.Context) error {
				A.SubAShortcutList(c)
				return nil
			},
		},
	}...)

	app.Before = func(c *cli.Context) error {
		c.App.Metadata[CONFIGPATH] = homedir + "/.grgd/config.yml"
		c.App.Metadata[CONFIG] = CreateConfigObjectYaml(c.App.Metadata[CONFIGPATH].(string))
		c.App.Metadata["pluginIndex"] = homedir + "/.grgd/plugins/index.yaml"
		c.App.Metadata["repoIndex"] = "https://s3.gregod.com/public/plugins/index.yaml"
		c.App.Metadata["AWS-REGION"] = "eu-central-1"
		c.App.Metadata["updatecheckinterval"] = time.Millisecond * 50
		// c.App.Metadata["currentcontext"] = A.UtilGetCurrentKubeContext()

		c.App.Metadata["UIPlugin"] = UIPlugin

		SystemCheck(c)
		CheckUpdate(c)

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
