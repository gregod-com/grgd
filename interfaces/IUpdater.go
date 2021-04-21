package interfaces

type IUpdater interface {
	CheckUpdate(version string, core ICore) error
}
