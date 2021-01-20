package interfaces

// ICMDPlugin ...
type ICMDPlugin interface {
	// GetCommands returns all urfave cli commands implemented by this plugin
	GetCommands(i interface{}) interface{}
}
