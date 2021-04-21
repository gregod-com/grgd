package helper

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/gregod-com/grgd/interfaces"
	"golang.org/x/mod/semver"
)

// ProvideNetworker ...
func ProvideNetworker(logger interfaces.ILogger) interfaces.INetworker {
	networker := new(Networker)
	networker.logger = logger
	logger.Tracef("provide %T", networker)
	return networker
}

// IndexObject ...
type IndexObject struct {
	Releases map[string]Category `yaml:"category"`
}

// Category ...
type Category struct {
	Targets map[string]Target `yaml:"target"`
}

type Target struct {
	Versions map[string]DownloadMetadata `yaml:"version"`
}

// Version ...
// type Version struct {
// 	DownloadMetadata map[string]DownloadMetadata
// }

// DownloadMetadata ...
type DownloadMetadata struct {
	Author       string
	Description  string
	Md5          string
	Released     time.Time
	Size         int
	URL          string
	ReleaseNotes string
}

// Connection ...
type Connection struct {
	Endpoint string
	TimeOut  int
	Success  bool
}

// Networker ...
type Networker struct {
	logger interfaces.ILogger
}

func catchOut(binPath string, args ...string) string {
	cmd := exec.Command(binPath, args...)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal("Error executing: " + err.Error())
	}
	return strings.TrimSuffix(out.String(), "\n")
}

func sortSemverSlice(semverSlice []string) error {
	// sort.Strings(sortedScriptVersions)
	sort.Slice(semverSlice, func(i, j int) bool {
		switch semver.Compare(semverSlice[i], semverSlice[j]) {
		case 1:
			return true
		default:
			return false
		}
	})
	return nil
}

var downloadSize uint64

// Load ...
func (n *Networker) Load(filepath string, url string) error {

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
	fmt.Println()
	return err
}

// WriteCounter ...
type WriteCounter struct {
	Total uint64
}

func (wc *WriteCounter) Write(p []byte) (int, error) {
	n := len(p)
	wc.Total += uint64(n)
	wc.PrintProgress()
	return n, nil
}

// PrintProgress ...
func (wc WriteCounter) PrintProgress() {
	// Clear the line by using a character return to go back to the start and remove
	// the remaining characters by filling it with spaces
	fmt.Printf("\r%s", strings.Repeat(" ", 35))

	// Return again and print current status of download
	// We use the humanize package to print the bytes in a meaningful way (e.g. 10 MB)
	fmt.Printf("\rDownloading... %d (%d MB) (%d%%) ", wc.Total/1024/1024, downloadSize/1024/1024, (wc.Total * 100 / downloadSize))
}

// CheckConnections ...
func (n *Networker) CheckConnections(conns map[string]interface{}) {
	for k := range conns {
		if conn, ok := conns[k].(Connection); ok {
			_, err := http.Get(conn.Endpoint)
			if err != nil {
				log.Fatal(err)
			}
			conn.Success = true
		}
	}
	return
}
