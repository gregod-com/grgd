package initialise

import (
	"testing"

	"github.com/gregod-com/grgd/controller/config"
	"github.com/gregod-com/grgd/core"
	"github.com/gregod-com/grgd/interfaces/mocks"

	"github.com/golang/mock/gomock"
	"github.com/tj/assert"
	"github.com/urfave/cli/v2"
)

func testHelperDefaultDepenedecyMap(ctrl *gomock.Controller) map[string]interface{} {
	mockHelper := mocks.NewMockIHelper(ctrl)
	mockLogger := mocks.NewMockILogger(ctrl)
	mockDAL := mocks.NewMockIDAL(ctrl)
	mockConfig := mocks.NewMockIConfig(ctrl)
	mockUI := mocks.NewMockIUIPlugin(ctrl)
	mockPlLoader := mocks.NewMockIPluginLoader(ctrl)
	mockFSM := mocks.NewMockIFileSystemManipulator(ctrl)

	mockLogger.EXPECT().Tracef(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes()
	mockLogger.EXPECT().Trace(gomock.Any()).AnyTimes()
	mockPlLoader.EXPECT().LoadPlugins(gomock.Any())
	mockFSM.EXPECT().HomeDir(".grgd", "plugins")
	mockDAL.EXPECT().GetProfile().AnyTimes()

	deps := map[string]interface{}{
		"IHelper":                mockHelper,
		"IDAL":                   mockDAL,
		"ILogger":                mockLogger,
		"IConfig":                mockConfig,
		"IUIPlugin":              mockUI,
		"IPluginLoader":          mockPlLoader,
		"IFileSystemManipulator": mockFSM,
	}

	return deps
}

func TestInit(t *testing.T) {
	// Given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	depsMap := testHelperDefaultDepenedecyMap(ctrl)
	mockUI := mocks.NewMockIUIPlugin(ctrl)
	mockUI.EXPECT().Println(gomock.Any())
	depsMap["IUIPlugin"] = mockUI
	depsMap["IConfig"] = config.ProvideConfig

	core := core.RegisterDependecies(depsMap)
	app := cli.NewApp()
	app.Metadata = make(map[string]interface{})
	app.Metadata["core"] = core

	// When
	c := cli.NewContext(app, nil, nil)

	// Then
	assert.Nil(t, AInit(c))
}
