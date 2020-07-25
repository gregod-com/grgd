package helpers_test

import (
	"os"
	"testing"

	"github.com/gregod-com/grgd/helpers"
	"github.com/gregod-com/grgd/implementations/config"
	th "github.com/gregod-com/grgd/testhelpers"
	"github.com/gregod-com/interfaces"
	"github.com/urfave/cli/v2"
)

func TestHomeDirReturnsValidFolder(t *testing.T) {
	result := helpers.HomeDir()

	if result == "" {
		t.Errorf("HomeDir is empty string ")
	}
	if _, err := os.Stat(result); os.IsNotExist(err) {
		t.Errorf("returned HomeDir does not exists")
	}
}

func TestExtractMetadataWithValidString_1(t *testing.T) {
	// Given
	testmap := make(map[string]interface{}, 0)
	testmap["one"] = "Hello"

	// When
	var container string
	result := helpers.ExtractMetadata(testmap, "one", &container)

	// Then
	th.CheckErrorNil(t, result)
	th.AssertEqual(t, "Hello", container)
}

func TestExtractMetadataWithValidString_2(t *testing.T) {
	// Given
	testmap := make(map[string]interface{}, 0)
	testmap["two"] = "World"

	// When
	var container string
	result := helpers.ExtractMetadata(testmap, "two", &container)

	// Then
	th.CheckErrorNil(t, result)
	th.AssertEqual(t, "World", container)

}

func TestExtractMetadataWithValidInt_1(t *testing.T) {
	// Given
	testmap := make(map[string]interface{}, 0)
	testmap["three"] = 3

	// When
	var container int
	result := helpers.ExtractMetadata(testmap, "three", &container)

	// Then
	th.CheckErrorNil(t, result)
	th.AssertEqual(t, 3, container)

}

func TestExtractMetadataWithInValidInt_1(t *testing.T) {
	// Given
	testmap := make(map[string]interface{}, 0)
	testmap["three"] = 3

	// When
	var container int
	result := helpers.ExtractMetadata(testmap, "one", &container)

	// Then
	th.CheckErrorNotNil(t, result)
}

func TestExtractMetadataWithValidSturct_1(t *testing.T) {
	// Given
	testmap := make(map[string]interface{}, 0)
	testmap["one"] = &helpers.FallbackUI{}

	// When
	var container *helpers.FallbackUI
	result := helpers.ExtractMetadata(testmap, "one", &container)

	// Then
	th.CheckErrorNil(t, result)
}

func TestExtractMetadataWithInValidSturct_1(t *testing.T) {
	// Given
	testmap := make(map[string]interface{}, 0)
	testmap["one"] = &helpers.FallbackUI{}

	// When
	var container *cli.Int64Flag
	result := helpers.ExtractMetadata(testmap, "one", &container)

	// Then
	th.CheckErrorNotNil(t, result)
}

func TestExtractMetadataWithValidInterface_1(t *testing.T) {
	// Given
	testmap := make(map[string]interface{}, 0)
	testmap["one"] = &config.ConfigObjectYaml{ProjectDirectory: "test/value"}

	// When
	var target interfaces.IConfigObject
	result := helpers.ExtractMetadata(testmap, "one", &target)

	// Then
	th.CheckErrorNil(t, result)
	th.AssertEqual(t, "test/value", target.GetProjectDir())
	th.AssertEqual(t, testmap["one"], target)
}
