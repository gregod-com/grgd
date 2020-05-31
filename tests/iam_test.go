package test

import (
	"bytes"
	"log"
	"os"
	"testing"
)

// capture output of a function and return the string
func captureOutput(f func()) string {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	f()
	log.SetOutput(os.Stderr)
	return buf.String()
}

func TestHelloWorld(t *testing.T) {
	WelloHorld := "Wello Horld"
	if WelloHorld != "Wello Horld" {
		t.Errorf("wrong")
	}
}

func TestPrintBanner(t *testing.T) {
	PrintBanner()
}

func TestMainFunction(t *testing.T) {
	// GIVEN
	output := captureOutput(func() { PrintBanner() })

	// WHEN
	want := "iamCLI"

	// THEN
	assert.Containsf(t, output, want, "Does not contain %s", want)

}
