package interfaces

// IPinger ...
type IPinger interface {
	CheckConnections(conns map[string]interface{})
}
