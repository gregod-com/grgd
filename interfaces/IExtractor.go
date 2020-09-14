package interfaces

// IExtractor ...
type IExtractor interface {
	// ExtractLogger ...
	GetCore(i interface{}) ICore
	// ExtractMetadata ...
	GetMetadata(m map[string]interface{}, key string, container interface{}) error
	// ExtractMetadataFatal calls ExtractMetadata but fails fataly before returning to caller
	// if extraction has error. This allows for less lines in calling code for essential
	// extractions that would need to interrupt the application anyways
	GetMetadataFatal(m map[string]interface{}, key string, container interface{})
}
