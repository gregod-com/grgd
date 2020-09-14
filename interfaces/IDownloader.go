package interfaces

// IDownloader ...
type IDownloader interface {
	Load(filepath string, url string) error
}
