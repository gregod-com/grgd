package helper_test

import (
	"grgd/interfaces/mocks"
	"testing"

	"grgd/controller/helper"
	h "grgd/controller/helper"

	"grgd/core"

	"github.com/golang/mock/gomock"
)

func TestCheckUpdate(t *testing.T) {
	// Given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	updater := &Updater{}
	mockDownloader := mocks.NewMockIDownloader(ctrl)
	mockUIPlugin := mocks.NewMockIUIPlugin(ctrl)
	core := controller.CreateCore(time.Now(), nil, nil, nil, nil, nil, mockUIPlugin, mockDownloader)

	// When
	mockDownloader.EXPECT().Load(gomock.Eq("file_location"), gomock.Eq("repo_url")).Return(nil)
	mockUIPlugin.EXPECT().Println(gomock.Eq(nil), gomock.Eq("Downloaded: ")).Return(nil)

	// Then
	updater.CheckUpdate(core)
}
