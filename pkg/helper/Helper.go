package helper

import (
	"bytes"
	"io"
	"os"
	"os/exec"
	"path"
	"strings"

	"github.com/gregod-com/grgd/interfaces"
	"gopkg.in/yaml.v2"
)

// ProvideHelper ...
func ProvideHelper(logger interfaces.ILogger) interfaces.IHelper {
	h := new(Helper)
	logger.Tracef("provide %T", h)
	h.logger = logger
	h.basedir = h.HomeDir(".grgd")
	h.bootconfigname = "bootconfig.yaml"
	h.CheckOrCreateFolder(h.basedir, os.FileMode(uint32(0760)))
	return h
}

// Helper ...
type Helper struct {
	logger         interfaces.ILogger
	basedir        string
	bootconfigname string
}

// CheckUserProfile ...
func (h *Helper) CheckUserProfile() string {
	h.logger.Tracef("")
	var profilename string
	u, ok := os.LookupEnv("USER")
	if !ok {
		h.logger.Fatal("failed to lookup USER ENV VAR! Exiting")
	}

	if u == "root" {
		h.logger.Fatal("You should not run this app as root! Exiting")
	}

	profilename = u

	if p := h.CheckFlagArg("profile"); p != "" {
		profilename = p
	}
	h.logger.Trace("Found profile")
	return profilename
}

// CheckFlagArg ...
func (h *Helper) CheckFlagArg(flag string) string {
	h.logger.Tracef("")
	for k, v := range os.Args {
		if v == "--"+flag && len(os.Args) > k+1 {
			return os.Args[k+1]
		}
	}
	return ""
}

// CheckFlag ...
func (h *Helper) CheckFlag(flag string) bool {
	h.logger.Tracef("")
	for _, v := range os.Args {
		if v == "-"+flag {
			return true
		}
		if v == "--"+flag {
			return true
		}
	}
	return false
}

// HomeDir ...
func (h *Helper) HomeDir(i ...string) string {
	h.logger.Tracef("")
	dir, err := os.UserHomeDir()
	if err != nil {
		h.logger.Fatal(err)
	}
	for _, v := range i {
		dir = path.Join(dir, v)
	}
	return dir
}

// CurrentWorkdir ...
func (h *Helper) CurrentWorkdir(i ...string) string {
	h.logger.Tracef("")
	dir, err := os.Getwd()
	if err != nil {
		h.logger.Fatal(err)
	}
	for _, v := range i {
		dir = path.Join(dir, v)
	}
	return dir
}

// CheckOrCreateFolder ...
func (h *Helper) CheckOrCreateFolder(pathToCheck string, permissions os.FileMode) {
	h.logger.Tracef("")
	if !h.PathExists(pathToCheck) {
		os.MkdirAll(pathToCheck, permissions)
	}
}

// CheckOrCreateParentFolder ...
func (h *Helper) CheckOrCreateParentFolder(pathToCheck string, permissions os.FileMode) {
	h.logger.Tracef("")
	dir, _ := path.Split(pathToCheck)
	h.CheckOrCreateFolder(dir, permissions)
}

// PathExists ...
func (h *Helper) PathExists(path string) bool {
	h.logger.Tracef("")
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}

func (h *Helper) ReadFile(path string) ([]byte, error) {
	return os.ReadFile(path)
}

func (h *Helper) UpdateOrWriteFile(path string, content []byte, permissions os.FileMode) error {
	return os.WriteFile(path, content, permissions)
}

// LoadBootConfig ...
func (h *Helper) LoadBootConfig() *interfaces.Bootconfig {
	h.logger.Tracef("")
	bootconfigpath := path.Join(h.basedir, h.bootconfigname)
	bootconfig := &interfaces.Bootconfig{}
	if !h.PathExists(bootconfigpath) {
		h.createDefaultConfig(bootconfigpath)
	}
	dat, err := h.ReadFile(bootconfigpath)
	if err != nil {
		h.logger.Fatal("Error reading bootconfig yaml")
	}

	if err := yaml.Unmarshal(dat, bootconfig); err != nil {
		h.logger.Fatal("Error unmarshalling bootconfig yaml")
	}
	return bootconfig
}

func (h *Helper) CatchOutput(script string, silent bool, args ...string) (string, error) {
	cmd := exec.Command(script, args...)
	var out, errout bytes.Buffer
	cmd.Stdout = io.MultiWriter(os.Stdout, &out)
	cmd.Stderr = io.MultiWriter(os.Stderr, &errout)
	if silent {
		cmd.Stdout = &out
		cmd.Stderr = &errout
	}
	err := cmd.Run()
	stdout := strings.Trim(out.String(), "\n")
	stderr := strings.Trim(errout.String(), "\n")
	return stdout + stderr, err
}

func (h *Helper) createDefaultConfig(bootconfigpath string) {
	h.logger.Tracef("")
	newbootconfig := &interfaces.Bootconfig{
		DatabasePath: path.Join(h.basedir, "grgd.db"),
	}
	dat, err := yaml.Marshal(newbootconfig)
	if err != nil {
		h.logger.Fatal("Error writing bootconfig yaml")
	}
	h.UpdateOrWriteFile(bootconfigpath, dat, os.FileMode(0760))
}
