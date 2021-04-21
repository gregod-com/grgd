package grgdk3d

import (
	"bytes"
	"fmt"
	"io"
	"math/rand"
	"os"
	"os/exec"
	"path"
	"strconv"
	"strings"
	"time"

	"github.com/gregod-com/grgd/interfaces"
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
	core := helper.GetExtractor().GetCore(ctx)
	log := core.GetLogger()

	startClusterOpts := types.ClusterStartOpts{}
	clusters, err := client.ClusterList(ctx.Context, runtimes.SelectedRuntime)
	if err != nil {
		log.Fatal(err)
	}

	for _, c := range clusters {
		if err := client.ClusterStart(ctx.Context, runtimes.SelectedRuntime, c, startClusterOpts); err != nil {
			log.Fatal(err)
		}
	}
	return nil
}

func down(ctx *cli.Context) error {
	core := helper.GetExtractor().GetCore(ctx)
	log := core.GetLogger()

	clusters, err := client.ClusterList(ctx.Context, runtimes.SelectedRuntime)
	if err != nil {
		log.Fatal(err)
	}

	for _, c := range clusters {
		if err := client.ClusterStop(ctx.Context, runtimes.SelectedRuntime, c); err != nil {
			log.Fatal(err)
		}
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
	ui := core.GetUI()
	log := core.GetLogger()
	h := core.GetHelper()
	currentProj := core.GetConfig().GetActiveProfile().GetCurrentProject()
	if currentProj == nil {
		return fmt.Errorf("Current project not defined")
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
	manifests := k3dMeta["manifestsdir"].(string)

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
			return fmt.Errorf("Cluster %s is running, please stop before booting a new cluster", c.Name)
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

	vol1 := v1alpha2.VolumeWithNodeFilters{
		Volume:      manifests + ":/var/lib/rancher/k3s/server/manifests/mountedManifests",
		NodeFilters: []string{"server[*]", "agent[*]"},
	}
	conf.Volumes = []v1alpha2.VolumeWithNodeFilters{vol1}

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
		return fmt.Errorf("Failed to create cluster '%s' because a cluster with that name already exists", clusterConfig.Cluster.Name)
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

	/**************
	* Kubeconfig *
	**************/

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

	importToRancher(clusterConfig.Cluster.Name, log)

	return nil
}

func delete(ctx *cli.Context) error {
	core := helper.GetExtractor().GetCore(ctx)
	// ui := core.GetUI()
	log := core.GetLogger()

	clusters := []*types.Cluster{}
	clusternames := []string{}
	if ctx.NArg() != 0 {
		clusternames = ctx.Args().Slice()
	}

	for _, name := range clusternames {
		c, err := client.ClusterGet(ctx.Context, runtimes.SelectedRuntime, &types.Cluster{Name: name})
		if err != nil {
			if err == client.ClusterGetNoNodesFoundError {
				continue
			}
			return err
		}
		clusters = append(clusters, c)
	}

	if len(clusters) == 0 {
		return fmt.Errorf("No clusters found")
	}

	log.Infof("Checking for clusters on Rancher...")

	clustermap, err := getClusterIDMap()
	if err != nil {
		return err
	}

	for _, c := range clusters {
		if c.Name == "local" || strings.Contains(c.Name, "prod") {
			log.Fatal("You really should not try to delete those clusters...")
		}
		if clusterID, ok := clustermap[c.Name]; !ok {
			log.Warnf("Could not find cluster %s on Rancher. Cannot delete automatically.", c.Name)
		} else {
			log.Infof("Removing cluster %s from Rancher...", c.Name)
			outDeleteCommand, err := catchOutput(true, "rancher", "cluster", "delete", clusterID)
			if err != nil {
				return err
			}
			log.Infof("%s", outDeleteCommand)
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
	}

	return nil
}

func importToRancher(clustername string, log interfaces.ILogger) error {
	log.Info("Importing your k3d cluster into your rancher instance...")
	outCreate, err := catchOutput(true, "rancher", "cluster", "create", "--import", clustername)
	if err != nil {
		return err
	}
	log.Warnf("%s", outCreate)
	clustermap, err := getClusterIDMap()
	if err != nil {
		return err
	}
	clusterID := clustermap[clustername]

	outImportCommand, err := catchOutput(true, "rancher", "cluster", "import", "-q", clusterID)
	if err != nil {
		return err
	}
	commands := strings.Split(outImportCommand, "\n")

	cmdAndArgs := strings.Split(commands[0], " ")
	_, err = catchOutput(true, cmdAndArgs[0], cmdAndArgs[1:]...)
	if err != nil {
		return err
	}
	return nil
}

func getClusterIDMap() (map[string]string, error) {
	var clustermap map[string]string
	outList, err := catchOutput(true, "rancher", "cluster", "ls", "--format", "{{.Cluster.Name}}: {{.Cluster.ID}}")
	if err != nil {
		return nil, err
	}
	yaml.Unmarshal([]byte(outList), &clustermap)
	return clustermap, nil
}

func catchOutput(silent bool, script string, args ...string) (string, error) {
	cmd := exec.Command(script, args...)
	var out, errout bytes.Buffer
	cmd.Stdout = io.MultiWriter(os.Stdout, &out)
	cmd.Stderr = io.MultiWriter(os.Stderr, &errout)
	if silent {
		cmd.Stdout = &out
		cmd.Stderr = &errout
	}
	err := cmd.Run()
	return out.String() + errout.String(), err
}
