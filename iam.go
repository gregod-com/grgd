package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"sort"
	"strings"

	"github.com/gregorpirolt/iamutils"

	"github.com/gregorpirolt/iamcli/templates"
	"gopkg.in/yaml.v3"

	"github.com/gregorpirolt/iamcli/ui"

	"github.com/urfave/cli"

	"time"

	tm "github.com/buger/goterm"
)

func main() {
	myFlags := []cli.Flag{
		cli.BoolFlag{
			Name:  "debug, d",
			Usage: "run the command in debug mode",
		},
		cli.BoolFlag{
			Name:  "network, n",
			Usage: "show network details in overview",
		},
		cli.BoolFlag{
			Name:  "mounts, m",
			Usage: "show network details in overview",
		},
		cli.BoolFlag{
			Name:  "sidecars, s",
			Usage: "show network details in overview",
		},
	}

	app := cli.NewApp()
	app.Flags = myFlags
	app.Name = "iamCLI"
	app.Usage = "written in go. Can be used as a sidekick to iamMenu and iamDoctr"
	app.Version = "1.0.0"
	app.Email = "gregor.pirolt@me.com"
	app.Copyright = "copyright yo"
	app.Author = "Gregor Pirolt"
	app.Metadata = make(map[string]interface{})
	app.Metadata["startTime"] = time.Now()
	app.Before = func(c *cli.Context) error {

		configObject := iamutils.IamConfigYaml{}
		configObject.InitFromFile("/.iam/iam_conf.yml")
		c.App.Metadata["iamconfig"] = configObject

		c.App.Metadata["services"] = iamutils.GenerateCLIServices(configObject)

		homeDir, err := os.UserHomeDir()
		if err != nil {
			return err
		}

		shortcutsYAMLContent, err := ioutil.ReadFile(homeDir + "/.iam/iam_shortcuts.yml")
		if err != nil {
			log.Println(err.Error())
			return err
		}

		shortcuts := map[string]string{}
		err = yaml.Unmarshal(shortcutsYAMLContent, &shortcuts)
		if err != nil {
			return err
		}

		c.App.Metadata["shortcuts"] = shortcuts

		out, err := exec.Command("kubectl", "config", "current-context").Output()
		if err != nil {
			log.Fatal(err)
		}

		c.App.Metadata["currentcontext"] = string(out)

		// always print banner before running the app
		ui.PrintBanner(c)

		app.CustomAppHelpTemplate = templates.GetHelpTemplate()
		return nil
	}
	app.HideHelp = true
	app.Action = func(c *cli.Context) error {
		ui.PrintServiceOverview(c)
		return nil
	}
	app.After = func(c *cli.Context) error {

		conf := c.App.Metadata["iamconfig"].(iamutils.IamConfigYaml)
		conf.Update()

		// shortcutsYAMLContent, err := yaml.Marshal(c.App.Metadata["shortcuts"].(map[string]string))
		// if err != nil {
		// 	return err
		// }

		// homeDir, err := os.UserHomeDir()
		// if err != nil {
		// 	return err
		// }

		// err = ioutil.WriteFile(homeDir+"/.iam/iam_shortcuts.yml", shortcutsYAMLContent, 644)
		// if err != nil {
		// 	log.Println(err.Error())
		// 	return err
		// }

		startTime := c.App.Metadata["startTime"].(time.Time)
		ui.PrintExecutionTime(time.Since(startTime))
		return nil
	}

	catStackControlls := "1) - " + tm.Color("stack controls", tm.BLUE)
	catServiceSpecific := "2) - " + tm.Color("service specific", tm.RED)
	catUtils := "3) - " + tm.Color("utilities", tm.CYAN)

	app.Commands = []cli.Command{
		{
			Name:     "up",
			Category: catStackControlls,
			Usage:    "Start dev stack with current config",
			Aliases:  []string{"u"},
			Flags:    myFlags,
			Action: func(c *cli.Context) error {
				return nil
			},
		},
		{
			Name:     "down",
			Category: catStackControlls,
			Usage:    "Stop dev stack (or single service)",
			Aliases:  []string{"d"},
			Flags:    myFlags,
			Action: func(c *cli.Context) error {
				return nil
			},
		},
		{
			Name:     "restart",
			Category: catStackControlls,
			Usage:    "Restart dev stack (or single service)",
			Aliases:  []string{"r"},
			Flags:    myFlags,
			Action: func(c *cli.Context) error {
				return nil
			},
		},
		{
			Name:        "logs",
			Category:    catStackControlls,
			Usage:       "Show all logs for running stack (or single service)",
			Aliases:     []string{"l"},
			Flags:       myFlags,
			UsageText:   "so this is the usage text",
			Description: "and what is this then?",
			ArgsUsage:   "Args usage",
			Action: func(c *cli.Context) error {
				return nil
			},
		},
		{
			Name:     "config",
			Category: catStackControlls,
			Usage:    "Configuration for current stack (type `config help` for possible subcommands)",
			Aliases:  []string{"conf", "c", "."},
			Flags:    myFlags,
			Subcommands: []cli.Command{
				{
					Name:    "yaml",
					Usage:   "Show iam_config.yaml file",
					Aliases: []string{"y"},
					Action: func(c *cli.Context) error {
						configObject := c.App.Metadata["iamconfig"].(iamutils.IamConfigYaml)
						configObject.PrintSourceYaml()
						return nil
					},
				},
			},
			Action: func(c *cli.Context) error {
				// services  := c.App.Metadata["services"].(map[string]iamutils.CliService)

				// composeChain := concatDockerComposeFiles(services)

				// // ctx := context.Background()
				// mycli, err := client.NewClientWithOpts(client.WithAPIVersionNegotiation())
				// if err != nil {
				// 	return nil
				// }

				// mycli.Close()
				// images, err := mycli.ImageList(context.Background(), types.ImageListOptions{})
				// if err != nil {
				// 	return nil
				// }

				// for _, image := range images {
				// 	fmt.Println(image.ID)
				// }

				// os.Stdin, os.Stdout, os.Stderr
				// cli2, err := command.NewDockerCli()
				// if err != nil {
				// 	return err
				// }
				// cli2.ClientInfo()
				// // fmt.Println(flags.DefaultCaFile)
				// cli2.Initialize(flags.NewClientOptions())
				// cmd := stack.NewStackCommand(cli2)

				// // the command package will pick up these, but you could override if you need to
				// cmd.SetArgs(append([]string{"deploy", "mystack"}, composeChain...))

				// cmd.Execute()
				// // project, err := docker.NewProject(&ctx.Context{
				// 	Context: project.Context{
				// 		ComposeFiles: composeChain,
				// 		ProjectName:  "my-compose",
				// 	},
				// }, nil)

				// if err != nil {
				// 	log.Fatal(err)
				// }

				// err = project.Up(context.Background(), options.Up{})

				// if err != nil {
				// 	log.Fatal(err)
				// }

				// fmt.Println(composeChain)

				// cmd := exec.Command("docker-compose", "-v")
				// // cmd := exec.Command("docker-compose", composeChain...)
				// cmd.Stdout = os.Stdout
				// cmd.Stderr = os.Stderr
				// err := cmd.Run()
				// if err != nil {
				// 	log.Fatalf("cmd.Run() failed with %s\n", err)
				// }

				return nil
			},
		},
		{
			Name:     "enter",
			Category: catServiceSpecific,
			Usage:    "Enter a service",
			Aliases:  []string{"en"},
			Flags:    myFlags,
			Action: func(c *cli.Context) error {
				str := ""
				for {
					// tm.Clear()
					str = str + "="

					// By moving cursor to top-left position we ensure that console output
					// will be overwritten each time, instead of adding new.
					tm.MoveCursor(1, 1)
					tm.Println("Current Time:", time.Now().Format(time.RFC1123))

					tm.Println(str + ">")
					tm.Println("\nCancel with Ctrl + c")
					tm.Flush() // Call it every time at the end of rendering
					time.Sleep(time.Millisecond * 15)
				}
				return nil
			},
		},
		{
			Name:     "execute",
			Category: catServiceSpecific,
			Usage:    "Execute a command in service and view output",
			Aliases:  []string{"exec", "ex"},
			Flags:    myFlags,
			Action: func(c *cli.Context) error {
				return nil
			},
		},
		{
			Name:        "test",
			Category:    catServiceSpecific,
			Usage:       "Run unittest inside container",
			Aliases:     []string{"t"},
			Flags:       myFlags,
			UsageText:   "TODO",
			Description: "TODO",
			ArgsUsage:   "TODO",
			Action: func(c *cli.Context) error {
				fmt.Println("")
				fmt.Println("ACTION: (test)")
				fmt.Println("TODO")
				return nil
			},
		},
		{
			Name:        "activate",
			Category:    catServiceSpecific,
			Usage:       "Activate a service",
			Aliases:     []string{"act", "a"},
			Flags:       myFlags,
			UsageText:   "TODO",
			Description: "TODO",
			ArgsUsage:   "TODO",
			After: func(c *cli.Context) error {
				ui.PrintServiceOverview(c)
				return nil
			},
			Action: func(c *cli.Context) error {
				services := TranslateShortcuts(c)
				configObject := c.App.Metadata["iamconfig"].(iamutils.IamConfigYaml)

				for _, serviceToActivate := range services {
					for _, s := range configObject.IamServiceSettings {
						if serviceToActivate == s.Name {
							s.ToggleActive()
							configObject.IamServiceSettings[s.Name] = s
							c.App.Metadata["services"] = iamutils.GenerateCLIServices(configObject)
						}
					}
				}
				return nil
			},
		},
		{
			Name:     "shortcuts",
			Category: catUtils,
			Usage:    "Show and edit shortcut names for services",
			Aliases:  []string{"sc"},
			Flags:    myFlags,
			After: func(c *cli.Context) error {
				PrintShortcuts(c)
				return nil
			},
			Subcommands: []cli.Command{
				{
					Name:  "add",
					Usage: "add new shortcut `ssc add shortcut service` ",
					Action: func(c *cli.Context) error {
						if c.NArg() != 2 {
							return cli.NewExitError("You should enter a service and a shortcut", 5)
						}
						shortcut := c.Args()[0]
						service := c.Args()[1]

						shortcuts := c.App.Metadata["shortcuts"].(map[string]string)

						if val := shortcuts[shortcut]; val == "" {
							for _, s := range c.App.Metadata["services"].([]iamutils.DockerComposePod) {
								fmt.Println(s)
								// if s.GetName() == service {
								// 	shortcuts[shortcut] = service
								// 	fmt.Println("Added new shortcut: " + tm.Color(shortcut, tm.RED) + " -> " + tm.Color(shortcuts[shortcut], tm.CYAN))

								// 	return nil
								// }
							}
							return cli.NewExitError("The service "+tm.Color(service, tm.CYAN)+" is not part of your stack. You can list all services with the command `iam config`", 6)
						}
						return cli.NewExitError(tm.Color(shortcut, tm.RED)+" already exists and points to "+tm.Color(shortcuts[shortcut], tm.CYAN), 7)
					},
				},
				{
					Name:  "remove",
					Usage: "remove a shortcut `ssc remove shortcut` ",
					Action: func(c *cli.Context) error {
						if c.NArg() > 1 {
							return cli.NewExitError("You should only one shortcut at a time", 5)
							return nil
						}
						shortcut := c.Args()[0]

						shortcuts := c.App.Metadata["shortcuts"].(map[string]string)

						if val := shortcuts[shortcut]; val == "" {
							return cli.NewExitError("There is no shortcut "+tm.Color(shortcut, tm.RED), 8)
						}
						fmt.Println("Removed shortcut: " + tm.Color(shortcut, tm.RED) + " -> " + tm.Color(shortcuts[shortcut], tm.CYAN))
						delete(shortcuts, shortcut)

						return nil
					},
				},
			},
			Action: func(c *cli.Context) error {
				return nil
			},
		},
		{
			Name:     "volume",
			Category: catUtils,
			Usage:    "View and edit service attached volumes",
			Aliases:  []string{"vol"},
			Flags: append(myFlags, cli.BoolFlag{
				Name:  "print_volume",
				Usage: "run the command without coloring output",
			},
			),
			UsageText:   "TODO",
			Description: "TODO",
			ArgsUsage:   "TODO",
			Action: func(c *cli.Context) error {
				c.Set("print_volume", "true")
				ui.PrintServiceOverview(c)
				fmt.Println("")
				fmt.Println("ACTION: (volume)")
				fmt.Println("TODO")
				return nil
			},
		},
		{
			Name:        "certificates",
			Category:    catUtils,
			Usage:       "View and fetch certificates from cluster",
			Aliases:     []string{"cert"},
			UsageText:   "TODO",
			Description: "TODO",
			ArgsUsage:   "TODO",
			Action: func(c *cli.Context) error {
				c.Set("print_volume", "true")
				ui.PrintServiceOverview(c)
				fmt.Println("")
				fmt.Println("ACTION: (volume)")
				fmt.Println("TODO")
				return nil
			},
		},
		{
			Name:     "dns",
			Category: catUtils,
			Usage:    "View and edit DNS routing",
			Action: func(c *cli.Context) error {
				return nil
			},
		},
		{
			Name:     "settings",
			Category: catUtils,
			Usage:    "View and edit global settings",
			Action: func(c *cli.Context) error {
				return nil
			},
		},
		{
			Name:     "context",
			Category: catUtils,
			Usage:    "View and edit kubernetes context",
			Aliases:  []string{"cont", "kubec"},
			Action: func(c *cli.Context) error {
				out, err := exec.Command("kubectl", "config", "get-contexts").Output()
				if err != nil {
					log.Fatal(err)
				}
				fmt.Println(string(out))
				if c.NArg() > 0 {
					for _, word := range strings.Fields(string(out)) {
						if strings.Contains(word, c.Args().First()) {
							fmt.Println("Setting current context to '" + word + "'")
							out, err := exec.Command("kubectl", "config", "use-context", word).Output()
							if err != nil {
								log.Fatal(err)
							}
							fmt.Println(string(out))
							return nil
						}
					}
				}
				return nil
			},
		},
		{
			Name:     "helm",
			Category: catUtils,
			Usage:    "View and edit helm deployments",
			Aliases:  []string{"he"},
			Action: func(c *cli.Context) error {
				out, err := exec.Command("helm", "ls", "--all-namespaces").Output()
				if err != nil {
					log.Fatal(err)
				}
				fmt.Println(string(out))
				return nil
			},
		},
	}

	sort.Sort(cli.FlagsByName(app.Flags))

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

// TranslateShortcuts ...
func TranslateShortcuts(c *cli.Context) []string {
	shortcuts := c.App.Metadata["shortcuts"].(map[string]string)
	services := make([]string, c.NArg())

	for k, v := range c.Args() {
		services[k] = v
		if service := shortcuts[v]; service != "" {
			services[k] = service
		}
	}
	return services
}

// PrintShortcuts ...
func PrintShortcuts(c *cli.Context) {
	shortcuts := c.App.Metadata["shortcuts"].(map[string]string)
	fmt.Println("\nShortcuts: ")
	sorted := []string{}
	for scs, service := range shortcuts {
		sorted = append(sorted, fmt.Sprintf("\t%-15v->%-15v", scs, service))
	}
	sort.Strings(sorted)
	for _, s := range sorted {
		fmt.Println(s)
	}
}

func concatDockerComposeFiles(s map[string]iamutils.CliService) []string {
	composeChain := []string{}
	for _, service := range s {
		_, _, active := service.GetActive()
		if active {
			composeChain = append(composeChain, "-c")
			composeChain = append(composeChain, service.Pod.Path)
		}
	}
	return composeChain
}
