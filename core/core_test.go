package core

import (
	"fmt"
	"grgd/interfaces"
	"testing"

	"github.com/gregod-com/grgd/controller/helper"
	"github.com/stretchr/testify/assert"
)

func TestRegisterDependecies_With_Nil(t *testing.T) {
	// Given
	// When
	RegisterDependecies(
		map[string]interface{}{})
}

func TestRegisterDependecies_With_Helper(t *testing.T) {
	// Given
	// When
	RegisterDependecies(
		map[string]interface{}{
			"interfaces.IHelper": helper.ProvideHelper,
		})
}

func TestRegisterDependecies_With_Helper_2(t *testing.T) {
	// Given

	// When
	mycore := RegisterDependecies(
		map[string]interface{}{
			"interfaces.IHelper": helper.ProvideHelper,
		})

	// Then
	assert.NotNil(t, mycore.GetHelper(), "core returned nil for helper")
	hlpr := mycore.GetHelper()

	var helperint interfaces.IHelper
	assert.Implements(t, &helperint, hlpr)
}

func TestRegisterDependecies_With_Helper_3(t *testing.T) {
	// Given
	// When
	mycore := RegisterDependecies(
		map[string]interface{}{
			"hans": nil,
		})

	// Then
	assert.NotNil(t, mycore.GetHelper(), "core returned nil for helper")
	hlpr := mycore.GetHelper()
	var helperint interfaces.IHelper
	assert.Implements(t, &helperint, hlpr)

}

func TestRegisterDependecies_With_Helper_4(t *testing.T) {
	// Given
	// When
	mycore := RegisterDependecies(
		map[string]interface{}{
			"interfaces.IHelper":  helper.ProvideHelper,
			"interfaces.IHelper2": helper.ProvideHelper,
			"interfaces.IHelper3": helper.ProvideHelper,
			"interfaces.IHelper4": helper.ProvideHelper,
		})

	// Then
	assert.NotNil(t, mycore.GetHelper(), "core returned nil for helper")
	hlpr := mycore.GetHelper()
	var helperint interfaces.IHelper
	assert.Implements(t, &helperint, hlpr)
}

func TestRegisterDependecies_With_Helper_5(t *testing.T) {
	// Given
	// When
	mycore := RegisterDependecies(
		map[string]interface{}{
			"interfaces.IHelper": helper.ProvideHelper,
		})

	// Then
	assert.NotNil(t, mycore.GetHelper(), "core returned nil for helper")
	hlpr := mycore.GetHelper()
	var helperint interfaces.IHelper
	assert.Implements(t, &helperint, hlpr)
}

func TestRegisterDependecies_With_Variadric_Provider(t *testing.T) {
	// Given
	// When
	mycore := RegisterDependecies(
		map[string]interface{}{
			"fmt.Println": fmt.Println,
		})

	// Then
	assert.NotNil(t, mycore.GetHelper(), "core returned nil for helper")
	hlpr := mycore.GetHelper()
	var helperint interfaces.IHelper
	assert.Implements(t, &helperint, hlpr)
}