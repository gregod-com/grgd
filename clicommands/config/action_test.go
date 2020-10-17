package config

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/gregod-com/grgd/core"
	"github.com/gregod-com/grgd/interfaces/mocks"
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

func TestSubAConfigYAML(t *testing.T) {
	// Given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	depsMap := testHelperDefaultDepenedecyMap(ctrl)

	// When
	core := core.RegisterDependecies(depsMap)

	// Then
	SubAConfigYAML()
}
