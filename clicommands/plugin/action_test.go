package plugin

import (
	"grgd/controller/config"
	"grgd/core"
	"grgd/interfaces/mocks"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/urfave/cli/v2"
)

func testHelperDefaultDepenedecyMap(ctrl *gomock.Controller) map[string]interface{} {
	mockHelper := mocks.NewMockIHelper(ctrl)
	mockLogger := mocks.NewMockILogger(ctrl)
	mockDAL := mocks.NewMockIDAL(ctrl)
	mockConfigObject := mocks.NewMockIConfigObject(ctrl)
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

func TestAPluginList(t *testing.T) {
	// Given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	depsMap := testHelperDefaultDepenedecyMap(ctrl)
	mockUI := mocks.NewMockIUIPlugin(ctrl)
	mockUI.EXPECT().Println(gomock.Any())
	depsMap["IUIPlugin"] = mockUI
	depsMap["IConfigObject"] = config.ProvideConfigObject

	core := core.RegisterDependecies(depsMap)
	app := cli.NewApp()
	app.Metadata = make(map[string]interface{})
	app.Metadata["core"] = core

	// When
	c := cli.NewContext(app, nil, nil)

	// Then
	assert.NoError(t, APluginList(c))
}

func TestAPluginActivate(t *testing.T) {
	// Given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	depsMap := testHelperDefaultDepenedecyMap(ctrl)
	mockUI := mocks.NewMockIUIPlugin(ctrl)
	mockPluginIndex := mocks.NewMockIPluginIndex(ctrl)

	mockUI.EXPECT().Println(gomock.Any())
	mockPluginIndex.EXPECT().GetPluginList()

	depsMap["IUIPlugin"] = mockUI
	depsMap["IConfigObject"] = config.ProvideConfigObject
	depsMap["IPluginIndex"] = mockPluginIndex

	core := core.RegisterDependecies(depsMap)
	app := cli.NewApp()
	app.Metadata = make(map[string]interface{})
	app.Metadata["core"] = core

	// When
	c := cli.NewContext(app, nil, nil)

	// Then
	assert.NoError(t, APluginActivate(c))
}
