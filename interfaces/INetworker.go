package interfaces

// INetworker ...
type INetworker interface {
	CheckUpdate(version string, core ICore) error
	CheckConnections(conns map[string]interface{})
	// Download the fiel located at arg `url` to the absolute path at arg `filepath`
	Load(filepath string, url string) error
}
