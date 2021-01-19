package interfaces

// IUpdater ...
type IUpdater interface {
	CheckUpdate(version string, core ICore) error
}
