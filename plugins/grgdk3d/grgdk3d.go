package grgdk3d

import (
	"fmt"
	"math/rand"
	"os"
	"path"
	"strconv"
	"strings"
	"time"

	"github.com/gregod-com/grgd/pkg/helper"
	"gopkg.in/yaml.v2"

	"github.com/rancher/k3d/v4/pkg/client"
	"github.com/rancher/k3d/v4/pkg/config"
	"github.com/rancher/k3d/v4/pkg/config/v1alpha2"
	"github.com/rancher/k3d/v4/pkg/runtimes"
	"github.com/rancher/k3d/v4/pkg/types"
	"github.com/rancher/k3d/v4/pkg/util"

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
			Aliases:         []string{"cluster", "k3s"},
			Category:        "grgd-native",
			Usage:           "handle local k3d-clusters",
			HideHelpCommand: true,
			Before:          nil,
			Flags:           app.Flags,
			Subcommands: []*cli.Command{
				{
					Name:    "up",
					Usage:   "start default cluster",
					Aliases: []string{"u", "start"},
					Flags:   app.Flags,
					Action:  up,
				},
				{
					Name:    "down",
					Usage:   "stop the default cluster",
					Aliases: []string{"d", "stop"},
					Flags:   app.Flags,
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
	c, err := getDefaultClusterOfFromFirstArgument(ctx)
	if err != nil {
		return err
	}
	startClusterOpts := types.ClusterStartOpts{
		WaitForServer: true,
	}
	if err := client.ClusterStart(ctx.Context, runtimes.SelectedRuntime, c, startClusterOpts); err != nil {
		return err
	}
	return nil
}

func down(ctx *cli.Context) error {
	c, err := getDefaultClusterOfFromFirstArgument(ctx)
	if err != nil {
		return err
	}
	if err := client.ClusterStop(ctx.Context, runtimes.SelectedRuntime, c); err != nil {
		return err
	}
	return nil
}

func list(ctx *cli.Context) error {
	core := helper.GetExtractor().GetCore(ctx)
	ui := core.GetUI()
	log := core.GetLogger()
	clusters, err := client.ClusterList(ctx.Context, runtimes.SelectedRuntime)
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
	core.CallPreHook(ctx)
	ui := core.GetUI()
	log := core.GetLogger()
	h := core.GetHelper()

	currentProj := core.GetConfig().GetActiveProfile().GetCurrentProject()
	if currentProj == nil {
		return fmt.Errorf("current project not defined")
	}
	obj, err := currentProj.ReadSettingsObject(h)
	if err != nil {
		return err
	}

	k3dMeta, ok := obj.Meta["k3d"].(map[interface{}]interface{})
	if !ok {
		return fmt.Errorf("k3dMeta not defined in project settings. (%T) ", obj.Meta["k3d"])
	}

	k3dyaml := k3dMeta["yaml"].(string)
	k3dbase := k3dMeta["basedir"].(string)

	if !h.PathExists(fmt.Sprint(k3dyaml)) {
		return fmt.Errorf("k3d yaml was defined but could not be found at path %s", k3dyaml)
	}

	clusters, err := client.ClusterList(ctx.Context, runtimes.SelectedRuntime)
	if err != nil {
		return err
	}

	for _, c := range clusters {
		totalA, runningA := c.AgentCountRunning()
		totalS, runningS := c.ServerCountRunning()
		running := runningA + runningS
		total := totalA + totalS
		ui.Printf("%v of %v nodes are running for cluster %s\n", running, total, c.Name)
		if running > 0 {
			return fmt.Errorf("cluster %s is running, please stop before booting a new cluster", c.Name)
		}
	}

	var conf v1alpha2.SimpleConfig

	out, err := h.ReadFile(k3dyaml)
	if err != nil {
		log.Error(err)
		return err
	}
	err = yaml.Unmarshal(out, &conf)
	if err != nil {
		log.Error(err)
		return err
	}

	conf.Name = h.CheckUserProfile()
	if ctx.NArg() > 0 {
		conf.Name = ctx.Args().First()
	}

	for k := range conf.Volumes {
		conf.Volumes[k].Volume = strings.Replace(conf.Volumes[k].Volume, "BASE", k3dbase, 1)
		if strings.Contains(conf.Volumes[k].Volume, "server-*") {
			conf.Volumes[k].Volume = strings.Replace(conf.Volumes[k].Volume, "server-*", "server-0", 1)
			conf.Volumes[k].NodeFilters = []string{"server[0]"}
			for i := 1; i < conf.Servers; i++ {
				conf.Volumes = append(conf.Volumes, v1alpha2.VolumeWithNodeFilters{
					Volume:      strings.Replace(conf.Volumes[k].Volume, "server-*", "server-"+fmt.Sprint(i), 1),
					NodeFilters: []string{"server[" + fmt.Sprint(i) + "]"},
				})
			}
		}
		if strings.Contains(conf.Volumes[k].Volume, "agent-*") {
			if conf.Agents > 0 {
				conf.Volumes[k].Volume = strings.Replace(conf.Volumes[k].Volume, "agent-*", "agent-0", 1)
				conf.Volumes[k].NodeFilters = []string{"agent[0]"}
				for i := 1; i < conf.Agents; i++ {
					conf.Volumes = append(conf.Volumes, v1alpha2.VolumeWithNodeFilters{
						Volume:      strings.Replace(conf.Volumes[k].Volume, "agent-*", "agent-"+fmt.Sprint(i), 1),
						NodeFilters: []string{"agent[" + fmt.Sprint(i) + "]"},
					})
				}
			} else {
				conf.Volumes = append(conf.Volumes[:k], conf.Volumes[k+1:]...)
			}
		}
	}

	conf.ExposeAPI.Host = "localhost"
	conf.ExposeAPI.HostIP = "127.0.0.1"
	conf.ExposeAPI.HostPort = strconv.Itoa(6444 + rand.Intn(100))

	clusterConfig, err := config.TransformSimpleToClusterConfig(ctx.Context, runtimes.SelectedRuntime, conf)
	if err != nil {
		log.Error(err)
		return err
	}

	clusterConfig, err = config.ProcessClusterConfig(*clusterConfig)
	if err != nil {
		log.Error(err)
		return err
	}
	if err := config.ValidateClusterConfig(ctx.Context, runtimes.SelectedRuntime, *clusterConfig); err != nil {
		log.Error(err)
		return err
	}

	// check if a cluster with that name exists already
	if _, err := client.ClusterGet(ctx.Context, runtimes.SelectedRuntime, &clusterConfig.Cluster); err == nil {
		return fmt.Errorf("failed to create cluster '%s' because a cluster with that name already exists", clusterConfig.Cluster.Name)
	}

	if clusterConfig.KubeconfigOpts.UpdateDefaultKubeconfig {
		log.Debug("'--kubeconfig-update-default set: enabling wait-for-server")
		clusterConfig.ClusterCreateOpts.WaitForServer = true
	}

	// create cluster
	if err := client.ClusterRun(ctx.Context, runtimes.SelectedRuntime, clusterConfig); err != nil {
		log.Error(err)
		return err
	}
	log.Infof("Cluster '%s' created successfully!", clusterConfig.Cluster.Name)

	if !clusterConfig.KubeconfigOpts.UpdateDefaultKubeconfig && clusterConfig.KubeconfigOpts.SwitchCurrentContext {
		log.Info("--kubeconfig-update-default=false --> sets --kubeconfig-switch-context=false")
		clusterConfig.KubeconfigOpts.SwitchCurrentContext = false
	}

	if clusterConfig.KubeconfigOpts.UpdateDefaultKubeconfig {
		log.Debugf("Updating default kubeconfig with a new context for cluster %s", clusterConfig.Cluster.Name)
		if _, err := client.KubeconfigGetWrite(ctx.Context, runtimes.SelectedRuntime, &clusterConfig.Cluster, "", &client.WriteKubeConfigOptions{UpdateExisting: true, OverwriteExisting: false, UpdateCurrentContext: conf.Options.KubeconfigOptions.SwitchCurrentContext}); err != nil {
			log.Warn(err)
		}
	}

	for tot, run := clusterConfig.Cluster.AgentCountRunning(); run < tot; tot, run = clusterConfig.Cluster.AgentCountRunning() {
		time.Sleep(time.Duration(time.Second * 2))
		log.Info("Waiting for nodes to be ready...")
	}

	ctx.App.Metadata["newCluster"] = clusterConfig.Cluster.Name

	return core.CallPostHook(ctx)
}

func delete(ctx *cli.Context) error {
	core := helper.GetExtractor().GetCore(ctx)
	core.CallPreHook(ctx)
	log := core.GetLogger()

	c, err := getDefaultClusterOfFromFirstArgument(ctx)
	if err != nil {
		return err
	}

	if err := client.ClusterDelete(ctx.Context, runtimes.SelectedRuntime, c, types.ClusterDeleteOpts{SkipRegistryCheck: false}); err != nil {
		log.Fatal(err)
	}

	log.Info("Removing cluster details from default kubeconfig...")
	if err := client.KubeconfigRemoveClusterFromDefaultConfig(ctx.Context, c); err != nil {
		log.Warn("Failed to remove cluster details from default kubeconfig")
		return err
	}
	log.Info("Removing standalone kubeconfig file (if there is one)...")
	configDir, err := util.GetConfigDirOrCreate()
	if err != nil {
		log.Warnf("Failed to delete kubeconfig file: %+v", err)
		return err
	} else {
		kubeconfigfile := path.Join(configDir, fmt.Sprintf("kubeconfig-%s.yaml", c.Name))
		if err := os.Remove(kubeconfigfile); err != nil {
			if !os.IsNotExist(err) {
				log.Warnf("Failed to delete kubeconfig file '%s'", kubeconfigfile)
				return err
			}
		}
	}

	log.Infof("Successfully deleted cluster %s!", c.Name)

	return core.CallPostHook(ctx)
}

func getDefaultClusterOfFromFirstArgument(ctx *cli.Context) (*types.Cluster, error) {
	core := helper.GetExtractor().GetCore(ctx)
	name := core.GetConfig().GetActiveProfile().GetName()
	if ctx.NArg() > 0 {
		name = ctx.Args().First()
	}

	c, err := client.ClusterGet(ctx.Context, runtimes.SelectedRuntime, &types.Cluster{Name: name})
	if err != nil {
		return nil, err
	}

	return c, nil
}
