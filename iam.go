// This package implements the cli for the iam-stack. The underlying framework is depnedent upon urfave/cli.
package main

import (
	"log"
	"os"
	"sort"

	at "github.com/gregorpirolt/animaterm"
	A "github.com/gregorpirolt/iamcli/actions"
	T "github.com/gregorpirolt/iamcli/templates"
	UI "github.com/gregorpirolt/iamcli/ui"
	Impl "github.com/gregorpirolt/implementations"
	I "github.com/gregorpirolt/interfaces"

	"github.com/urfave/cli"

	"time"

	tm "github.com/buger/goterm"
)

var iamUI = at.CreateUI()

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
			Name:    "sidecars, s",
			Aliases: []string{"sc"},
			Usage:   "show sidecars in overview",
		},
	}
	catStackControlls := "1) - " + tm.Color("stack controls", tm.BLUE)
	catWorkloadSpecific := "2) - " + tm.Color("workload specific", tm.RED)
	catUtils := "3) - " + tm.Color("utilities", tm.CYAN)

	app := cli.NewApp()
	app.Flags = myFlags
	app.Name = "iamCLI"
	app.Usage = "written in go. Can be used as a sidekick to iamMenu and iamDoctr"
	app.Version = "0.5.0"
	app.Metadata = make(map[string]interface{})
	app.Metadata["startTime"] = time.Now()
	app.CustomAppHelpTemplate = T.GetHelpTemplate()
	app.Before = func(c *cli.Context) error {
		userHome, err := os.UserHomeDir()
		if err != nil {
			return nil
		}
		c.App.Metadata["configLocation"] = userHome + "/.iam/iam_conf.yml"
		c.App.Metadata["iamconfig"] = Impl.CreateConfigObjectYaml(c.App.Metadata["configLocation"].(string))
		c.App.Metadata["workloads"] = c.App.Metadata["iamconfig"].(I.IConfigObject).GetWorkloads()
		c.App.Metadata["currentcontext"] = A.UtilGetCurrentKubeContext()
		iamUI.ClearScreen()
		iamUI.SetBoarder(5)
		c.App.Metadata["animation"] = int64(0)
		if c.App.Metadata["iamconfig"].(I.IConfigObject).
			GetLastUsed().Add(time.Duration(5) * time.Second).
			Before(time.Now()) {
			c.App.Metadata["animation"] = int64(1)
		}

		c.App.Metadata["iamui"] = iamUI
		UI.PrintBanner(c)
		return nil
	}
	app.Action = func(c *cli.Context) error {
		UI.PrintWorkloadOverview(c)
		return nil
	}
	app.After = func(c *cli.Context) error {
		c.App.Metadata["iamconfig"].(I.IConfigObject).Update()
		startTime := c.App.Metadata["startTime"].(time.Time)
		UI.PrintExecutionTime(time.Since(startTime))
		return nil
	}

	app.Commands = []*cli.Command{
		{
			Name:     "init",
			Category: catUtils,
			Usage:    "Initialze the " + app.Name,
			Before:   UI.CheckNewSkill,
			Action:   A.AInit,
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
			Before:   UI.CheckNewSkill,
			Action:   A.AUp,
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
			Before:   UI.CheckNewSkill,
			Action:   A.ADown,
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
			Before:   UI.CheckNewSkill,
			Action:   A.ARestart,
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
			Before:    UI.CheckNewSkill,
			Action:    A.ALogs,
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
					Before:  UI.CheckNewSkill,
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
			Before:   UI.CheckNewSkill,
			Action:   A.AEnter,
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
			Before:   UI.CheckNewSkill,
			Action:   A.AExecute,
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
			Before:    UI.CheckNewSkill,
			Action:    A.ATest,
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
			Before:    UI.CheckNewSkill,
			Action:    A.AActivate,
			After: func(c *cli.Context) error {
				UI.PrintWorkloadOverview(c)
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
					Before:  UI.CheckNewSkill,
					Description: `
					The 'shortcuts list' command is good and professional.
					`,
				},
				{
					Name:    "add",
					Aliases: []string{"a"},
					Usage:   "add new shortcut `sc add shortcut workload` ",
					Action:  A.SubAShortcut["add"],
					Before:  UI.CheckNewSkill,
					Description: `
					The 'shortcuts add' command is good and professional.
					`,
				},
				{
					Name:    "remove",
					Aliases: []string{"r"},
					Usage:   "remove a shortcut `ssc remove shortcut` ",
					Action:  A.SubAShortcut["remove"],
					Before:  UI.CheckNewSkill,
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
			Before:    UI.CheckNewSkill,
			Action:    A.AVolume,
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
			Before:    UI.CheckNewSkill,
			Action:    A.ACertificates,
			Description: `
			The 'certificates' command is good and professional.
			`,
		},
		{
			Name:     "dns",
			Category: catUtils,
			Usage:    "View and edit DNS routing",
			Before:   UI.CheckNewSkill,
			Action:   A.ADNS,
			Description: `
			The 'dns' command is good and professional.
			`,
		},
		{
			Name:     "context",
			Category: catUtils,
			Usage:    "View and edit kubernetes context",
			Aliases:  []string{"cont", "kubec"},
			Before:   UI.CheckNewSkill,
			Action:   A.AContext,
			After:    A.AfterContext,
			Description: `
			The 'context' command is good and professional.
			`,
		},
		{
			Name:     "helm",
			Category: catUtils,
			Usage:    "View and edit helm deployments",
			Aliases:  []string{"he"},
			Before:   UI.CheckNewSkill,
			Action:   A.AHelm,
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
