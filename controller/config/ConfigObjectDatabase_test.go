package config

import (
	"grgd/interfaces"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/gregod-com/grgd/core"
	"github.com/gregod-com/grgd/interfaces/mocks"
	"github.com/tj/assert"
)

func testHelperDefaultDepenedecyMap(ctrl *gomock.Controller) map[string]interface{} {
	mockHelper := mocks.NewMockIHelper(ctrl)
	mockLogger := mocks.NewMockILogger(ctrl)
	mockDAL := mocks.NewMockIDAL(ctrl)
	mockConfigObject := mocks.NewMockIConfigObject(ctrl)

	mockLogger.EXPECT().Tracef(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes()
	mockLogger.EXPECT().Trace(gomock.Any()).AnyTimes()

	return map[string]interface{}{
		"IHelper":       mockHelper,
		"IDAL":          mockDAL,
		"ILogger":       mockLogger,
		"IConfigObject": mockConfigObject,
	}
}

func TestProvideConfigObject(t *testing.T) {
	// Given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	depsMap := testHelperDefaultDepenedecyMap(ctrl)
	depsMap["IConfigObject"] = ProvideConfigObject

	// When
	core := core.RegisterDependecies(depsMap)

	// Then
	var configInterface interfaces.IConfigObject
	assert.Implements(t, &configInterface, core.GetConfig())
}

func TestSave(t *testing.T) {
	// Given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	depsMap := testHelperDefaultDepenedecyMap(ctrl)
	depsMap["IConfigObject"] = ProvideConfigObject

	mockDAL := mocks.NewMockIDAL(ctrl)

	mockDAL.EXPECT().Update(gomock.Any())
	depsMap["IDAL"] = mockDAL

	// When
	core := core.RegisterDependecies(depsMap)
	core.GetConfig().Save()

	// Then

}
