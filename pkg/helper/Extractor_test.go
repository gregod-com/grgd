package helper

import (
	"testing"

	"github.com/gregod-com/grgd/view"

	"github.com/tj/assert"
	"github.com/urfave/cli/v2"
)

func TestExtractMetadataWithValidString_1(t *testing.T) {
	// Given
	ext := GetExtractor()
	testmap := make(map[string]interface{}, 0)
	testmap["one"] = "Hello"

	// When
	var container string
	result := ext.GetMetadata(testmap, "one", &container)

	// Then
	assert.Nil(t, result, "Extract produces error")
	assert.Equal(t, "Hello", container)
}

func TestExtractMetadataWithValidString_2(t *testing.T) {
	// Given
	ext := GetExtractor()
	testmap := make(map[string]interface{}, 0)
	testmap["two"] = "World"

	// When
	var container string
	result := ext.GetMetadata(testmap, "two", &container)

	// Then
	assert.Nil(t, result, "Extract produces error")
	assert.Equal(t, "World", container)

}

func TestExtractMetadataWithValidInt_1(t *testing.T) {
	// Given
	ext := GetExtractor()
	testmap := make(map[string]interface{}, 0)
	testmap["three"] = 3

	// When
	var container int
	result := ext.GetMetadata(testmap, "three", &container)

	// Then
	assert.Nil(t, result, "Extract produces error")
	assert.Equal(t, 3, container)

}

func TestExtractMetadataWithInValidInt_1(t *testing.T) {
	// Given
	ext := GetExtractor()
	testmap := make(map[string]interface{}, 0)
	testmap["three"] = 3

	// When
	var container int
	result := ext.GetMetadata(testmap, "one", &container)

	// Then
	assert.Error(t, result, "Could not find key `one` in passed map", "Extract should produce error")

}

func TestExtractMetadataWithValidSturct_1(t *testing.T) {
	// Given
	ext := GetExtractor()
	testmap := make(map[string]interface{}, 0)
	testmap["one"] = &view.FallbackUI{}

	// When
	var container *view.FallbackUI
	result := ext.GetMetadata(testmap, "one", &container)

	// Then
	assert.Nil(t, result, "Extract produces error")

}

func TestExtractMetadataWithInValidSturct_1(t *testing.T) {
	// Given
	ext := GetExtractor()
	testmap := make(map[string]interface{}, 0)
	testmap["one"] = &view.FallbackUI{}

	// When
	var container *cli.Int64Flag
	result := ext.GetMetadata(testmap, "one", &container)

	// Then
	assert.Error(t, result, "Value at key `one` (type *view.FallbackUI) in passed map is not assignable to pointer (type `*cli.Int64Flag`)", "Does not produce error")
}
