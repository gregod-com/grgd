package helper

import (
	"grgd/interfaces"
	"os"
	"path"
)

// ProvideFSManipulator ...
func ProvideFSManipulator(logger interfaces.ILogger) interfaces.IFileSystemManipulator {
	fsm := new(FSManipulator)
	fsm.logger = logger
	fsm.CheckOrCreateFolder(fsm.HomeDir(".grgd"), os.FileMode(uint32(0760)))
	return fsm
}

// FSManipulator ...
type FSManipulator struct {
	logger interfaces.ILogger
}

// HomeDir ...
func (h *FSManipulator) HomeDir(i ...string) string {
	dir, errHomeDir := os.UserHomeDir()
	if errHomeDir != nil {
		h.logger.Fatal(errHomeDir)
	}
	for _, v := range i {
		dir = path.Join(dir, v)
	}
	return dir
}

// CheckOrCreateFolder ...
func (h *FSManipulator) CheckOrCreateFolder(path string, permissions os.FileMode) {
	if !h.PathExists(path) {
		os.MkdirAll(path, permissions)
	}
}

// PathExists ...
func (h *FSManipulator) PathExists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}
