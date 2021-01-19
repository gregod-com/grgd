package k8s

import (
	"context"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/gregod-com/grgd/controller/helper"
	"github.com/gregod-com/grgd/interfaces"
	"github.com/urfave/cli/v2"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

// ACertificate ...
func ACertificate(c *cli.Context) error {
	core := helper.GetExtractor().GetCore(c)
	UI := core.GetUI()
	logger := core.GetLogger()

	clientset := buildClientSet(c)
	tlsCerts := make(map[string]string)
	heads := []string{"Name", "Common Name"}
	rows := [][]string{}

	ctx := context.Background()

	secrets, err := clientset.CoreV1().Secrets("").List(ctx, metav1.ListOptions{})
	if err != nil {
		logger.Warnf(
			"%s\nAre you sure you are pointing to the correct k8s cluster?\n\n",
			err.Error())
		return cli.NewExitError("Count not connect to k8s cluster", 4)
	}

	if nrArgs := c.Args().Len(); nrArgs == 0 {
		for _, x := range secrets.Items {
			if x.Type == v1.SecretTypeTLS {
				if cn := x.GetAnnotations()["cert-manager.io/common-name"]; tlsCerts[x.GetName()] != "found" && cn != "" {
					tlsCerts[x.GetName()] = "found"
					rows = append(rows, []string{x.GetName(), cn})
				}
			}
		}
		UI.PrintTable(heads, rows, c)
		logger.Info("Add the certificates name as an argument")
		return cli.NewExitError("Missing argument", 2)
	}

	searchTerm := c.Args().First()
	UI.Println("Searching for certificate containing `" + searchTerm + "` ...")

	matchingSecrets := []v1.Secret{}

	for _, x := range secrets.Items {
		if x.Type == v1.SecretTypeTLS {
			if cn := x.GetAnnotations()["cert-manager.io/common-name"]; cn != "" && tlsCerts[x.GetName()] != "found" && strings.Contains(x.GetName(), searchTerm) {
				tlsCerts[x.GetName()] = "found"
				matchingSecrets = append(matchingSecrets, x)
			}
		}
	}

	if len(matchingSecrets) == 0 {
		UI.Println("Found no certificate containing `" + searchTerm + "`")
		UI.Println("  - if the search term is long, try to search for a shorter word")
		UI.Println("  - check if your internet is up")
		UI.Println("  - check if kubectl's context is pointing to the right cluster")
		UI.Println("  - execute the command without searchterm to list all possible tls certificates")
		return cli.NewExitError("", 0)
	}

	if len(matchingSecrets) > 1 {
		fmt.Printf("There were %v hits... please define certificate name closer\n", len(matchingSecrets))
		for _, i := range matchingSecrets {
			logger.Debug(i)
		}
		return cli.NewExitError("", 0)
	}

	if len(matchingSecrets) == 1 {
		certPEM := matchingSecrets[0].Data["tls.crt"]
		certDER, _ := pem.Decode(certPEM)

		certificate, err := x509.ParseCertificate(certDER.Bytes)
		if err != nil {
			return err
		}

		UI.Println("Found one certificate matching the search term: \n")
		UI.Println("Expires: " + certificate.NotAfter.String() + "\n")
		UI.Println("DNS Names: ")
		for _, v := range certificate.DNSNames {
			UI.Println(v)
		}
		UI.Println("")
		if UI.YesNoQuestion("Print as cleartext?") {
			UI.Println(string(matchingSecrets[0].Data["tls.crt"]))
			UI.Println(string(matchingSecrets[0].Data["tls.key"]))
		}
		if UI.YesNoQuestion("Print as base64?") {
			base64CRT := base64.StdEncoding.EncodeToString(matchingSecrets[0].Data["tls.crt"])
			base64KEY := base64.StdEncoding.EncodeToString(matchingSecrets[0].Data["tls.key"])
			UI.Printf("tls.crt: %s\n", base64CRT)
			UI.Printf("tls.crt: %s\n", base64KEY)
		}
	}

	return nil
}

// AContext ...
func AContext(c *cli.Context) error {
	core := helper.GetExtractor().GetCore(c)
	logger := core.GetLogger()
	var fsm interfaces.IFileSystemManipulator
	core.Get(&fsm)
	kubeConfigPath := fsm.HomeDir(".kube", "config")

	if c.NArg() > 0 {
		actualContextName, err := lookupContext(c.Args().First())
		if err != nil {
			logger.Fatal(err)
		}
		err = switchCurrentContext(actualContextName, kubeConfigPath)
		if err != nil {
			logger.Fatal(err)
		}
	}
	printKubeCTX(c)

	return nil
}

func lookupContext(ctx string) (string, error) {
	clientCfg, err := clientcmd.NewDefaultClientConfigLoadingRules().Load()
	if err != nil {
		return "", err
	}
	for k := range clientCfg.Contexts {
		if strings.Contains(k, ctx) {
			return k, nil
		}
	}
	return "", errors.New("context not found")
}

func printKubeCTX(c *cli.Context) {
	core := helper.GetExtractor().GetCore(c)
	UI := core.GetUI()
	logger := core.GetLogger()

	clientCfg, err := clientcmd.NewDefaultClientConfigLoadingRules().Load()
	if err != nil {
		logger.Fatal(err)
	}
	var rows [][]string
	var rowsUnsorted []string

	for k := range clientCfg.Contexts {
		spacing := "          "
		if clientCfg.CurrentContext == k {
			spacing = "(current) "
		}
		rowsUnsorted = append(rowsUnsorted, spacing+k)
	}
	sort.Strings(rowsUnsorted)
	for _, v := range rowsUnsorted {
		rows = append(rows, []string{v})
	}
	UI.PrintTable([]string{"Contexts"}, rows)
}

func buildClientSet(c *cli.Context) *kubernetes.Clientset {
	core := helper.GetExtractor().GetCore(c)
	// UI := core.GetUI()
	logger := core.GetLogger()

	var kubeconfig *string
	if homedir, err := os.UserHomeDir(); err != nil {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", filepath.Join(homedir, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	}
	flag.Parse()
	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		logger.Fatal(err.Error())
	}

	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		logger.Fatal(err.Error())
	}

	return clientset

}

func switchCurrentContext(context, kubeconfigPath string) error {
	cfg, _ := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
		&clientcmd.ClientConfigLoadingRules{ExplicitPath: kubeconfigPath},
		&clientcmd.ConfigOverrides{
			CurrentContext: context,
		}).RawConfig()
	cfg.CurrentContext = context
	return clientcmd.WriteToFile(cfg, kubeconfigPath)
}
