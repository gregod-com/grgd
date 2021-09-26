package interfaces

// INetworker ...
type INetworker interface {
	CheckConnections(conns map[string]interface{})
	// Download the fiel located at arg `url` to the absolute path at arg `filepath`
	Load(filepath string, url string) error
}
