package gormdal

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/gregod-com/grgd/interfaces"
	"github.com/gregod-com/grgd/interfaces/mocks"
	"github.com/tj/assert"
)

func TestLoadTESTService(t *testing.T) {
	// Given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	helper := mocks.NewMockIHelper(ctrl)
	logger := mocks.NewMockILogger(ctrl)

	helper.EXPECT().LoadBootConfig().Return(&interfaces.Bootconfig{DatabasePath: "testdatabase"}).AnyTimes()
	helper.EXPECT().CheckOrCreateFolder(gomock.Any(), gomock.Any()).AnyTimes()
	helper.EXPECT().CheckOrCreateParentFolder(gomock.Any(), gomock.Any()).AnyTimes()
	helper.EXPECT().HomeDir(gomock.Any(), gomock.Any()).AnyTimes()
	dal := setupDatabase(helper, logger)
	defer tearDownDatabase(helper)

	// When
	search := ServiceModel{Name: "TESTService"}
	err := dal.Read(&search)
	assert.Nil(t, err)

	// Then
	assert.Equal(t, "TESTService", search.Name)
	assert.Equal(t, "./test-me-dir/", search.Path)
}

func TestDeleteTESTService(t *testing.T) {
	// Given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	helper := mocks.NewMockIHelper(ctrl)
	logger := mocks.NewMockILogger(ctrl)

	helper.EXPECT().LoadBootConfig().Return(&interfaces.Bootconfig{DatabasePath: "testdatabase"}).AnyTimes()
	helper.EXPECT().CheckOrCreateFolder(gomock.Any(), gomock.Any()).AnyTimes()
	helper.EXPECT().CheckOrCreateParentFolder(gomock.Any(), gomock.Any()).AnyTimes()
	helper.EXPECT().HomeDir(gomock.Any(), gomock.Any()).AnyTimes()
	dal := setupDatabase(helper, logger)
	defer tearDownDatabase(helper)

	// When
	search := &ServiceModel{Name: "TESTService"}
	err := dal.Read(search)
	assert.Nil(t, err)
	dal.Delete(search)

	// Then
	err2 := dal.Read(search)
	assert.Error(t, err2)
}

func TestEditTESTService(t *testing.T) {
	// Given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	helper := mocks.NewMockIHelper(ctrl)
	logger := mocks.NewMockILogger(ctrl)

	helper.EXPECT().LoadBootConfig().Return(&interfaces.Bootconfig{DatabasePath: "testdatabase"}).AnyTimes()
	helper.EXPECT().CheckOrCreateFolder(gomock.Any(), gomock.Any()).AnyTimes()
	helper.EXPECT().CheckOrCreateParentFolder(gomock.Any(), gomock.Any()).AnyTimes()
	helper.EXPECT().HomeDir(gomock.Any(), gomock.Any()).AnyTimes()
	dal := setupDatabase(helper, logger)
	defer tearDownDatabase(helper)

	// When
	search := &ServiceModel{Name: "TESTService"}
	err := dal.Read(search)
	assert.Nil(t, err)
	search.Name = "edited-service"
	dal.Update(search)

	// Then

	err2 := dal.Read(search)
	assert.Nil(t, err2)

	assert.Equal(t, "edited-service", search.Name)
}

func TestAddServiceToProject(t *testing.T) {
	// Given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	helper := mocks.NewMockIHelper(ctrl)
	logger := mocks.NewMockILogger(ctrl)

	helper.EXPECT().LoadBootConfig().Return(&interfaces.Bootconfig{DatabasePath: "testdatabase"}).AnyTimes()
	helper.EXPECT().CheckOrCreateFolder(gomock.Any(), gomock.Any()).AnyTimes()
	helper.EXPECT().CheckOrCreateParentFolder(gomock.Any(), gomock.Any()).AnyTimes()
	helper.EXPECT().HomeDir(gomock.Any(), gomock.Any()).AnyTimes()
	dal := setupDatabase(helper, logger)
	defer tearDownDatabase(helper)

	// When
	searchProfile := &ProfileModel{Name: "TESTProfile"}
	err1 := dal.Read(searchProfile)
	assert.Nil(t, err1)

	searchProject := &ProjectModel{Name: "TESTProject"}
	err2 := dal.Read(searchProject)
	assert.Nil(t, err2)

	searchService := &ServiceModel{Name: "TESTService"}
	err3 := dal.Read(searchService)
	assert.Nil(t, err3)

	// searchProject.Services = append(searchProject.Services, searchService)
	searchProfile.Projects = append(searchProfile.Projects, searchProject)
	dal.Update(searchProfile)

	// Then
	err4 := dal.Read(searchProfile)
	assert.Nil(t, err4)

	assert.NotNil(t, searchProfile.Projects)
	assert.Equal(t, "TESTProject", searchProfile.Projects[0].Name)
	// assert.NotNil(t, searchProfile.Projects[0].Services)
	// assert.Equal(t, "TESTService", searchProfile.Projects[0].Services[0].Name)
}
