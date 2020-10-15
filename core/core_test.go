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
		[]interface{}{
			nil,
		})
}

func TestRegisterDependecies_With_Helper(t *testing.T) {
	// Given
	// When
	RegisterDependecies(
		[]interface{}{
			helper.ProvideHelper,
		})
}

func TestRegisterDependecies_With_Helper_2(t *testing.T) {
	// Given

	// When
	mycore := RegisterDependecies(
		[]interface{}{
			helper.ProvideHelper,
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
		[]interface{}{
			nil,
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
		[]interface{}{
			helper.ProvideHelper,
			helper.ProvideHelper,
			helper.ProvideHelper,
			helper.ProvideHelper,
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
		[]interface{}{
			helper.ProvideHelper(),
		})

	// Then
	assert.NotNil(t, mycore.GetHelper(), "core returned nil for helper")
	hlpr := mycore.GetHelper()
	var helperint interfaces.IHelper
	assert.Implements(t, &helperint, hlpr)
}

func TestRegisterDependecies_With_Random_Values_1(t *testing.T) {
	// Given
	// When
	mycore := RegisterDependecies(
		[]interface{}{
			fmt.Println,
		})

	// Then
	assert.NotNil(t, mycore.GetHelper(), "core returned nil for helper")
	hlpr := mycore.GetHelper()
	var helperint interfaces.IHelper
	assert.Implements(t, &helperint, hlpr)
}
