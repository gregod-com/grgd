package interfaces

// IUpdater ...
type IUpdater interface {
	CheckUpdate(core ICore) error
}
