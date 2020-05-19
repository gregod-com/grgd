package actions

import (
	"fmt"

	I "github.com/gregorpirolt/interfaces"
	"github.com/urfave/cli"
)

// Up start the stack
func AUp(c *cli.Context) error {
	go fmt.Println("this is the up command")
	return nil
}

func ADown(c *cli.Context) error {
	go fmt.Println("this is the down command")
	return nil
}

func ARestart(c *cli.Context) error {
	go fmt.Println("this is the restart command")
	return nil
}

func ALogs(c *cli.Context) error {
	go fmt.Println("this is the logs command")
	return nil
}

func AConfig(c *cli.Context) error {
	go fmt.Println("this is the config command")
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
}

// SubAConfig ...
var SubAConfig = map[string]func(*cli.Context) error {
	"yaml": func(c *cli.Context) error {
		fmt.Println("This is the config yaml subcommand")
		configObject := c.App.Metadata["iamconfig"].(I.IConfigObject)
		configObject.PrintConfig()
		return nil
	},
}
