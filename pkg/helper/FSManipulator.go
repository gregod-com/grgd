package helper

import (
	"io/ioutil"
	"os"
	"path"

	"github.com/gregod-com/grgd/interfaces"
	"gopkg.in/yaml.v2"
)

// ProvideFSManipulator ...
func ProvideFSManipulator(logger interfaces.ILogger) interfaces.IFileSystemManipulator {
	fsm := new(FSManipulator)
	logger.Tracef("provide %T", fsm)
	fsm.logger = logger
	fsm.basedir = fsm.HomeDir(".grgd")
	fsm.bootconfigname = "bootconfig.yaml"
	fsm.CheckOrCreateFolder(fsm.basedir, os.FileMode(uint32(0760)))
	return fsm
}

// FSManipulator ...
type FSManipulator struct {
	logger         interfaces.ILogger
	basedir        string
	bootconfigname string
}

// HomeDir ...
func (fsm *FSManipulator) HomeDir(i ...string) string {
	fsm.logger.Tracef("")
	dir, errHomeDir := os.UserHomeDir()
	if errHomeDir != nil {
		fsm.logger.Fatal(errHomeDir)
	}
	for _, v := range i {
		dir = path.Join(dir, v)
	}
	return dir
}

// CheckOrCreateFolder ...
func (fsm *FSManipulator) CheckOrCreateFolder(pathToCheck string, permissions os.FileMode) {
	fsm.logger.Tracef("")
	if !fsm.PathExists(pathToCheck) {
		os.MkdirAll(pathToCheck, permissions)
	}
}

// CheckOrCreateParentFolder ...
func (fsm *FSManipulator) CheckOrCreateParentFolder(pathToCheck string, permissions os.FileMode) {
	fsm.logger.Tracef("")
	dir, _ := path.Split(pathToCheck)
	fsm.CheckOrCreateFolder(dir, permissions)
}

// PathExists ...
func (fsm *FSManipulator) PathExists(path string) bool {
	fsm.logger.Tracef("")
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}

// LoadBootConfig ...
func (fsm *FSManipulator) LoadBootConfig() *interfaces.Bootconfig {
	fsm.logger.Tracef("")
	bootconfigpath := path.Join(fsm.basedir, fsm.bootconfigname)
	bootconfig := &interfaces.Bootconfig{}
	if !fsm.PathExists(bootconfigpath) {
		fsm.createDefaultConfig(bootconfigpath)
	}
	dat, err := ioutil.ReadFile(bootconfigpath)
	if err != nil {
		fsm.logger.Fatal("Error reading bootconfig yaml")
	}

	if err := yaml.Unmarshal(dat, bootconfig); err != nil {
		fsm.logger.Fatal("Error unmarshalling bootconfig yaml")
	}
	return bootconfig
}

func (fsm *FSManipulator) createDefaultConfig(bootconfigpath string) {
	fsm.logger.Tracef("")
	newbootconfig := &interfaces.Bootconfig{
		DatabasePath: path.Join(fsm.basedir, "grgd.db"),
	}
	dat, err := yaml.Marshal(newbootconfig)
	if err != nil {
		fsm.logger.Fatal("Error writing bootconfig yaml")
	}
	ioutil.WriteFile(bootconfigpath, dat, os.FileMode(0760))
}
