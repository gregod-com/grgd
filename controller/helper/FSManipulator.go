package helper

import (
	"grgd/interfaces"
	"os"
	"path"
)

type FSManipulator struct {
}

// HomeDir ...
func (h *FSManipulator) HomeDir(i ...string) string {
	dir, errHomeDir := os.UserHomeDir()
	if errHomeDir != nil {
		log.Fatal(errHomeDir)
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
