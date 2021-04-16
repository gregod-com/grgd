package config

import (
	"testing"

	"github.com/gregod-com/grgd/interfaces/mocks"

	"github.com/golang/mock/gomock"
)

func testHelperDefaultDepenedecyMap(ctrl *gomock.Controller) map[string]interface{} {
	mockHelper := mocks.NewMockIHelper(ctrl)
	mockLogger := mocks.NewMockILogger(ctrl)
	mockDAL := mocks.NewMockIDAL(ctrl)
	mockConfigObject := mocks.NewMockIConfig(ctrl)
	mockUI := mocks.NewMockIUIPlugin(ctrl)

	mockLogger.EXPECT().Tracef(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes()
	mockLogger.EXPECT().Trace(gomock.Any()).AnyTimes()

	deps := map[string]interface{}{
		"IHelper":       mockHelper,
		"IDAL":          mockDAL,
		"ILogger":       mockLogger,
		"IConfigObject": mockConfigObject,
		"IUIPlugin":     mockUI,
	}

	return deps
}

func TestHelloWorld(t *testing.T) {
	// Given

	// When

	// Then
}

// func TestSubAConfigYAML(t *testing.T) {
// 	// Given
// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()
// 	depsMap := testHelperDefaultDepenedecyMap(ctrl)
// 	mockUI := mocks.NewMockIUIPlugin(ctrl)
// 	mockUI.EXPECT().Println(gomock.Any())
// 	depsMap["IUIPlugin"] = mockUI
// 	depsMap["IConfigObject"] = config.ProvideConfigObject

// 	core := core.RegisterDependecies(depsMap)
// 	app := cli.NewApp()
// 	app.Metadata = make(map[string]interface{})
// 	app.Metadata["core"] = core

// 	// When
// 	c := cli.NewContext(app, nil, nil)

// 	// Then
// 	assert.NoError(t, SubAConfigYAML(c))
// }

// func TestSubAConfigJSON(t *testing.T) {
// 	// Given
// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()
// 	depsMap := testHelperDefaultDepenedecyMap(ctrl)
// 	mockUI := mocks.NewMockIUIPlugin(ctrl)
// 	depsMap["IUIPlugin"] = mockUI
// 	depsMap["IConfigObject"] = config.ProvideConfigObject
// 	mockUI.EXPECT().Println(gomock.Any())

// 	core := core.RegisterDependecies(depsMap)
// 	app := cli.NewApp()
// 	app.Metadata = make(map[string]interface{})
// 	app.Metadata["core"] = core

// 	// When
// 	c := cli.NewContext(app, nil, nil)

// 	// Then
// 	assert.NoError(t, SubAConfigJSON(c))
// }
