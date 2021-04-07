package interfaces

import "os"

// IHelper ...
type IHelper interface {
	CheckUserProfile() string
	CheckFlag(flag string) bool
	CheckFlagArg(flag string) string
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
