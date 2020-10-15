package project

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/gregod-com/grgd/core"
	"github.com/gregod-com/grgd/interfaces/mocks"
	"github.com/gregod-com/grgd/logger"
)

func TestProvider(t *testing.T) {
	// Given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// mockHelper := mocks.NewMockIHelper(ctrl)

	// l := logger.ProvideDefaultLogger(mockHelper)

	core := core.RegisterDependecies(
		[]interface{}{
			mocks.NewMockIHelper,
			logger.ProvideDefaultLogger,
		})
	core.GetLogger()

	// // time.Now(), nil, nil, nil, nil, nil, mockUIPlugin, mockDownloader)

	// // When
	// mockDownloader.EXPECT().Load(gomock.Eq("file_location"), gomock.Eq("repo_url")).Return(nil)
	// mockUIPlugin.EXPECT().Println(gomock.Eq(nil), gomock.Eq("Downloaded: ")).Return(nil)

	// // Then
	// updater.CheckUpdate(core)

	// var mockDownloader mocks.MockIDownloader
	// var mockUIPlugin mocks.MockIUIPlugin

	// core.Get(&mockDownloader)
	// core.Get(&mockUIPlugin)
	// mockDownloader.EXPECT().Load(gomock.Eq("file_location"), gomock.Eq("repo_url")).Return(nil)
	// mockUIPlugin.EXPECT().Println(gomock.Eq(nil), gomock.Eq("Downloaded: ")).Return(nil)

	// updater := helper.ProvideUpdater()
	// mockDownloader := mocks.NewMockIDownloader(ctrl)
	// mockUIPlugin := mocks.NewMockIUIPlugin(ctrl)
	// time.Now(), nil, nil, nil, nil, nil, mockUIPlugin, mockDownloader)

	// When

	// Then
	// updater.CheckUpdate(core)
	// p := ProvideProject()
	// p.GetName()

	// When

	// Then

}
