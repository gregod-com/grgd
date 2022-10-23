package core

import (
	"testing"

	"github.com/gregod-com/grgd/interfaces"

	"github.com/gregod-com/grgd/pkg/helper"
	"github.com/gregod-com/grgd/pkg/logger"
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
			"ILogger": logger.ProvideLogrusLogger,
			"IHelper": helper.ProvideHelper,
		})
}

func TestRegisterDependecies_With_Helper_2(t *testing.T) {
	// Given

	// When
	mycore, err := RegisterDependecies(
		map[string]interface{}{
			"ILogger": logger.ProvideLogrusLogger,
			"IHelper": helper.ProvideHelper,
		})

	// Then
	assert.NoError(t, err)
	assert.NotNil(t, mycore.GetHelper(), "core returned nil for helper")
	hlpr := mycore.GetHelper()

	var helperint interfaces.IHelper
	assert.Implements(t, &helperint, hlpr)
}

// func TestRegisterDependecies_With_Helper_3(t *testing.T) {
// 	// Given
// 	// When
// 	mycore, err := RegisterDependecies(
// 		map[string]interface{}{
// 			"ILogger": logger.ProvideLogrusLogger,
// 			"hans":    nil,
// 		})

// 	// Then
// 	assert.NoError(t, err)
// 	assert.NotNil(t, mycore.GetHelper(), "core returned nil for helper")
// 	hlpr := mycore.GetHelper()
// 	var helperint interfaces.IHelper
// 	assert.Implements(t, &helperint, hlpr)

// }

func TestRegisterDependecies_With_Helper_4(t *testing.T) {
	// Given
	// When
	mycore, err := RegisterDependecies(
		map[string]interface{}{
			"ILogger":             logger.ProvideLogrusLogger,
			"interfaces.IHelper":  helper.ProvideHelper,
			"interfaces.IHelper2": helper.ProvideHelper,
			"interfaces.IHelper3": helper.ProvideHelper,
			"interfaces.IHelper4": helper.ProvideHelper,
		})

	// Then
	assert.NoError(t, err)
	assert.NotNil(t, mycore.GetHelper(), "core returned nil for helper")
	hlpr := mycore.GetHelper()
	var helperint interfaces.IHelper
	assert.Implements(t, &helperint, hlpr)
}

func TestRegisterDependecies_With_Helper_5(t *testing.T) {
	// Given
	// When
	mycore, err := RegisterDependecies(
		map[string]interface{}{
			"ILogger":            logger.ProvideLogrusLogger,
			"interfaces.IHelper": helper.ProvideHelper,
		})

	// Then
	assert.NoError(t, err)
	assert.NotNil(t, mycore.GetHelper(), "core returned nil for helper")
	hlpr := mycore.GetHelper()
	var helperint interfaces.IHelper
	assert.Implements(t, &helperint, hlpr)
}

// func TestRegisterDependecies_With_Variadric_Provider(t *testing.T) {
// 	// Given
// 	// When
// 	mycore, err := RegisterDependecies(
// 		map[string]interface{}{
// 			"ILogger":     logger.ProvideLogrusLogger,
// 			"fmt.Println": fmt.Println,
// 		})

// 	// Then
// 	assert.NoError(t, err)
// 	assert.NotNil(t, mycore.GetHelper(), "core returned nil for helper")
// 	hlpr := mycore.GetHelper()
// 	var helperint interfaces.IHelper
// 	assert.Implements(t, &helperint, hlpr)
// }

// func TestRegisterDependecies_With_Struct(t *testing.T) {
// 	// Given
// 	configStruct := config.ConfigDatabase{}
// 	// When
// 	mycore, err := RegisterDependecies(
// 		map[string]interface{}{
// 			"ILogger": logger.ProvideLogrusLogger,
// 			"test":    configStruct,
// 		})

// 	// Then
// 	assert.NoError(t, err)
// 	assert.NotNil(t, mycore.GetHelper(), "core returned nil for helper")
// 	hlpr := mycore.GetHelper()
// 	var helperint interfaces.IHelper
// 	assert.Implements(t, &helperint, hlpr)
// }
