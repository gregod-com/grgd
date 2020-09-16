package helper

import (
	"grgd/interfaces"
	"io"
	"net/http"
	"os"
)

// Downloader ...
type Downloader struct{}

// Load ...
func (d *Downloader) Load(filepath string, url string) error {

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}
