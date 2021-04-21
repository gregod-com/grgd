package grgdk3d

import (
	"context"
	"fmt"

	"github.com/gregod-com/grgd/pkg/helper"

	"github.com/rancher/k3d/v4/pkg/client"
	"github.com/rancher/k3d/v4/pkg/runtimes"
	"github.com/rancher/k3d/v4/pkg/types"

	cli "github.com/urfave/cli/v2"
)

type CMD struct {
}

func (cmd *CMD) GetCommands(i interface{}) interface{} {
	app, ok := i.(*cli.App)
	if !ok {
		return fmt.Errorf("did not pass *cli.App")
	}
	return []*cli.Command{
		{
			Name:            "k3d",
			Category:        "grgd-native",
			Usage:           "handle local k3d-clusters",
			HideHelpCommand: true,
			Before:          nil,
			Flags:           app.Flags,
			Subcommands: []*cli.Command{
				{
					Name:    "up",
					Usage:   "start default cluster",
					Aliases: []string{"u"},
					Flags:   app.Flags,
					Action:  up,
				},
				{
					Name:    "down",
					Usage:   "stop the default cluster",
					Aliases: []string{"d"},
					Flags:   app.Flags,
					Before:  app.Before,
					Action:  down,
				},
				{
					Name:    "list",
					Usage:   "list clusters",
					Aliases: []string{"ls"},
					Flags:   app.Flags,
					Action:  list,
				},
				{
					Name:    "bootstrap",
					Usage:   "bootstrap and register new cluster for current project",
					Aliases: []string{"boot", "new"},
					Flags:   app.Flags,
					Action:  bootstrap,
				},
				{
					Name:    "delete",
					Usage:   "remove and unregister current for current project",
					Aliases: []string{"del", "rm"},
					Flags:   app.Flags,
					Action:  delete,
				},
			},
		},
	}
}

func up(ctx *cli.Context) error {
	core := helper.GetExtractor().GetCore(ctx)
	log := core.GetLogger()

	tempctx := context.Background()
	startClusterOpts := types.ClusterStartOpts{}
	clusters, err := client.ClusterList(tempctx, runtimes.SelectedRuntime)
	if err != nil {
		log.Fatal(err)
	}

	for _, c := range clusters {
		if err := client.ClusterStart(tempctx, runtimes.SelectedRuntime, c, startClusterOpts); err != nil {
			log.Fatal(err)
		}
	}
	return nil
}

func down(ctx *cli.Context) error {
	core := helper.GetExtractor().GetCore(ctx)
	log := core.GetLogger()

	tempctx := context.Background()
	clusters, err := client.ClusterList(ctx.Context, runtimes.SelectedRuntime)
	if err != nil {
		log.Fatal(err)
	}

	for _, c := range clusters {
		if err := client.ClusterStop(tempctx, runtimes.SelectedRuntime, c); err != nil {
			log.Fatal(err)
		}
	}
	return nil
}

func list(ctx *cli.Context) error {
	tempctx := context.Background()
	core := helper.GetExtractor().GetCore(ctx)
	ui := core.GetUI()
	log := core.GetLogger()
	clusters, err := client.ClusterList(tempctx, runtimes.SelectedRuntime)
	if err != nil {
		log.Fatal(err)
	}

	for _, c := range clusters {
		ui.Printf("===== k3d-%-45v\n", c.Name)
		ui.Printf("%-25v %-15v %-20v| Network: %v\n", "Name", "State", "Role", c.Network.Name)
		for k, node := range c.Nodes {
			ui.Printf("%-25v %-15v %-20v", node.Name, node.State.Status, node.Role)
			switch k {
			case 0:
				ui.Printf("| Token: %-25v\n", c.Token)
			case 1:
				ui.Printf("| ImageVolume: %-25v\n", c.ImageVolume)
			default:
				ui.Println()
			}
		}
	}
	return nil
}

func bootstrap(ctx *cli.Context) error {
	core := helper.GetExtractor().GetCore(ctx)
	ui := core.GetUI()
	log := core.GetLogger()
	currentProj := core.GetConfig().GetActiveProfile().GetCurrentProject()
	if currentProj == nil {
		core.GetConfig().GetActiveProfile().AddProjectByName("testa")
	}
	currentProj = core.GetConfig().GetActiveProfile().GetCurrentProject()
	if currentProj == nil {
		log.Fatal("here1")
	}

	ui.Printf("So here we are proj %s\n", currentProj.GetName())

	// FILE=iamk3d.yaml

	// if k3d node ls | grep -q ' running';
	// then
	// 	echo "\nThere is already a k3d cluster running. Exiting script now.\n\n"
	// 	exit 1
	// fi

	// if test -f "$FILE"; then
	//   echo "Using $FILE as config for new cluster."
	// else
	//   echo "$FILE does not exist in current path. Move to iamk3d project or create $FILE"
	//   exit 1
	// fi

	// if [ -z "$3" ]
	// then
	//   CLUSTERNAME=$USER
	// else
	//   CLUSTERNAME=$3
	// fi

	// if [ -z "$4" ]
	// then
	// 	k3sversion=latest
	// else
	// 	k3sversion=$4
	// fi

	// if docker pull rancher/k3s:$k3sversion;
	// then
	//   echo "\n\nAll good, starting k3d cluster now...\n\n"
	// else
	//   echo "\n\nInvalid k3s version. Check releases for valid tags: https://hub.docker.com/r/rancher/k3s/tags\n\n"
	//   exit 1
	// fi

	// k3d cluster create $CLUSTERNAME -c $FILE \
	// 	--image rancher/k3s:$k3sversion \
	// 	--volume $(PWD)/iammanifests/:/var/lib/rancher/k3s/server/manifests/iammanifests@all \
	// 	--volume $(PWD)/registries.yaml:/etc/rancher/k3s/registries.yaml@all

	// if k3d node ls | grep -q ' running';
	// then
	// 	resetRancherContext
	// 	echo "\n\nImporting your k3d cluster into your rancher instance...\n\n"
	// 	rancher cluster create --import k3d-$CLUSTERNAME
	// 	clusterID=$(rancher cluster ls --format json | jq -cr --arg CN "k3d-$CLUSTERNAME" '. | select( .Name == $CN ) | .ID' )
	// 	rancher cluster import -q $clusterID | head -n 1 | sh -
	// else
	// 	echo "\n\n hmmm cluster not running... skipping import to rancher\n\n"
	// fi

	return nil
}

func delete(ctx *cli.Context) error {
	core := helper.GetExtractor().GetCore(ctx)
	ui := core.GetUI()
	log := core.GetLogger()
	ui.Println("tba - please use k3d cli as before")
	log.Fatal("here")

	return nil
}
