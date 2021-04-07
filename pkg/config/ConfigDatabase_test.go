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
	mockConfig := mocks.NewMockIConfig(ctrl)

	mockLogger.EXPECT().Tracef(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes()
	mockLogger.EXPECT().Trace(gomock.Any()).AnyTimes()
	mockHelper.EXPECT().HomeDir(".grgd", "plugins")

	deps := map[string]interface{}{
		"IHelper": mockHelper,
		"IDAL":    mockDAL,
		"ILogger": mockLogger,
		"IConfig": mockConfig,
	}
	return deps
}

func TestHelloWorld(t *testing.T) {
	// Given

	// When

	// Then
}

// func TestProvideConfig(t *testing.T) {
// 	// Given
// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()
// 	depsMap := testHelperDefaultDepenedecyMap(ctrl)
// 	depsMap["IConfig"] = ProvideConfig

// 	// When
// 	core := core.RegisterDependecies(depsMap)

// 	// Then
// 	var configInterface interfaces.IConfig
// 	assert.Implements(t, &configInterface, core.GetConfig())
// }

// func TestSave(t *testing.T) {
// 	// Given
// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()
// 	depsMap := testHelperDefaultDepenedecyMap(ctrl)
// 	depsMap["IConfig"] = ProvideConfig

// 	mockDAL := mocks.NewMockIDAL(ctrl)

// 	mockDAL.EXPECT().Update(gomock.Any())
// 	depsMap["IDAL"] = mockDAL

// 	// When
// 	core := core.RegisterDependecies(depsMap)
// 	core.GetConfig().Save()

// 	// Then

// }
