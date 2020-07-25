// This package implements the cli for the iam-stack. The underlying framework is depnedent upon urfave/cli.
package helpers

import (
	"io"
	"net/http"
	"os"
)

// DownloadFile updates and plugins from s3
func DownloadFile(filepath string, url string) error {

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
