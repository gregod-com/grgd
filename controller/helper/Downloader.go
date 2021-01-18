package helper

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gregod-com/grgd/interfaces"
)

// ProvideDownloader ...
func ProvideDownloader() interfaces.IDownloader {
	return new(Downloader)
}

// Downloader ...
type Downloader struct{}

var downloadSize uint64

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

	size, _ := strconv.Atoi(resp.Header.Get("Content-Length"))
	downloadSize = uint64(size)

	counter := &WriteCounter{}
	// Write the body to file
	_, err = io.Copy(out, io.TeeReader(resp.Body, counter))
	return err
}

type WriteCounter struct {
	Total uint64
}

func (wc *WriteCounter) Write(p []byte) (int, error) {
	n := len(p)
	wc.Total += uint64(n)
	wc.PrintProgress()
	return n, nil
}

func (wc WriteCounter) PrintProgress() {
	// Clear the line by using a character return to go back to the start and remove
	// the remaining characters by filling it with spaces
	fmt.Printf("\r%s", strings.Repeat(" ", 35))

	// Return again and print current status of download
	// We use the humanize package to print the bytes in a meaningful way (e.g. 10 MB)
	fmt.Printf("\rDownloading... %d (%d MB) (%d%%) ", wc.Total/1024/1024, downloadSize/1024/1024, (wc.Total * 100 / downloadSize))
}
