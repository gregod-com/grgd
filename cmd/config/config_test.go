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
