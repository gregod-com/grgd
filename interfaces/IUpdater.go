package interfaces

type IUpdater interface {
	CheckUpdate(version string, core ICore) error
	CheckSinceLastUpdate(version string, core ICore) error
}
