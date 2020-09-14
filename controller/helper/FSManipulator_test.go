package helper

import (
	"os"
	"path"
	"testing"

	"gotest.tools/v3/assert"
)

func TestHomeDirReturnsValidFolder(t *testing.T) {
	// Given
	ext := &FSManipulator{}

	// When
	result := ext.HomeDir()

	// Then
	assert.Assert(t, result != "", "HomeDir is empty string")
	// assert.NilError(t, err, "HomeDir does not exists")
}

func TestHomeDirWithArgs(t *testing.T) {
	// Given
	ext := &FSManipulator{}
	actualHomedir, err := os.UserHomeDir()
	folders := []string{actualHomedir, ".grgd", "Plugins"}

	// When
	result := ext.HomeDir(folders[1:]...)

	// Then
	assert.NilError(t, err, "Error getting homedir in test preparation")

	assert.Assert(t, result != "", "HomeDir is empty string")
	assert.Equal(t, path.Join(folders...), result, "Path was constructed wrong")

}
