package interfaces

import "os"

// IFileSystemManipulator ...
type IFileSystemManipulator interface {
	// HomeDir ...
	HomeDir(i ...string) string

	// CheckOrCreateFolder ...
	CheckOrCreateFolder(path string, permissions os.FileMode)

	// CheckOrCreateParentFolder ...
	CheckOrCreateParentFolder(path string, permissions os.FileMode)
	// PathExists ...
	PathExists(path string) bool

	LoadBootConfig() *Bootconfig
}

// Bootconfig ...
type Bootconfig struct {
	DatabasePath string
}
