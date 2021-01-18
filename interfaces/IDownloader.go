package interfaces

// IDownloader ...
type IDownloader interface {
	// Download the fiel located at arg `url` to the absolute path at arg `filepath`
	Load(filepath string, url string) error
}
