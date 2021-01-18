package logger

import (
	"testing"

	"github.com/gregod-com/grgd/interfaces"

	"github.com/golang/mock/gomock"
	"github.com/gregod-com/grgd/core"
	"github.com/gregod-com/grgd/interfaces/mocks"
	"github.com/tj/assert"
)

func TestProvideLogrusLogger(t *testing.T) {
	// Given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockHelper := mocks.NewMockIHelper(ctrl)

	// When
	mockHelper.EXPECT().CheckFlag("debug")
	mockHelper.EXPECT().CheckFlag("d")
	mockHelper.EXPECT().CheckFlagArg(gomock.Eq("log-level"))

	core := core.RegisterDependecies(
		map[string]interface{}{
			"IHelper": mockHelper,
			"ILogger": ProvideLogrusLogger,
		})

	// Then
	var loggerInterface interfaces.ILogger
	assert.Implements(t, &loggerInterface, core.GetLogger())
}
