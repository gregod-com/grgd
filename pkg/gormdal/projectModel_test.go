package gormdal

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/gregod-com/grgd/interfaces"
	"github.com/gregod-com/grgd/interfaces/mocks"
	"github.com/gregod-com/grgd/pkg/profile"
	"github.com/gregod-com/grgd/pkg/project"
	"github.com/tj/assert"
)

func TestLoadTESTProject(t *testing.T) {
	// Given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	helper := mocks.NewMockIHelper(ctrl)
	logger := mocks.NewMockILogger(ctrl)

	helper.EXPECT().LoadBootConfig().Return(&interfaces.Bootconfig{DatabasePath: "testdatabase"}).AnyTimes()
	helper.EXPECT().CheckOrCreateFolder(gomock.Any(), gomock.Any()).AnyTimes()
	helper.EXPECT().CheckOrCreateParentFolder(gomock.Any(), gomock.Any()).AnyTimes()
	helper.EXPECT().HomeDir(gomock.Any(), gomock.Any()).AnyTimes()
	logger.EXPECT().Tracef(gomock.Any(), gomock.Any()).AnyTimes()
	logger.EXPECT().GetLevel().AnyTimes()
	logger.EXPECT().Warnf(gomock.Any(), gomock.Any()).AnyTimes()
	dal := setupDatabase(helper, logger)
	defer tearDownDatabase(helper)

	// When
	search := ProjectModel{Name: "TESTProject"}
	err := dal.Read(&search)
	assert.Nil(t, err)

	// Then
	assert.Equal(t, "TESTProject", search.Name)
	assert.Equal(t, "./test-me-dir/", search.Path)
}

func TestDeleteTESTProject(t *testing.T) {
	// Given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	helper := mocks.NewMockIHelper(ctrl)
	logger := mocks.NewMockILogger(ctrl)

	helper.EXPECT().LoadBootConfig().Return(&interfaces.Bootconfig{DatabasePath: "testdatabase"}).AnyTimes()
	helper.EXPECT().CheckOrCreateFolder(gomock.Any(), gomock.Any()).AnyTimes()
	helper.EXPECT().CheckOrCreateParentFolder(gomock.Any(), gomock.Any()).AnyTimes()
	helper.EXPECT().HomeDir(gomock.Any(), gomock.Any()).AnyTimes()

	logger.EXPECT().Tracef(gomock.Any(), gomock.Any()).AnyTimes()
	logger.EXPECT().GetLevel().AnyTimes()
	logger.EXPECT().Warnf(gomock.Any(), gomock.Any()).AnyTimes()

	dal := setupDatabase(helper, logger)
	defer tearDownDatabase(helper)

	// When
	search := &ProjectModel{Name: "TESTProject"}
	err := dal.Read(search)
	assert.Nil(t, err)
	dal.Delete(search)

	// Then
	err2 := dal.Read(search)
	assert.Error(t, err2)
}

func TestEditTESTProject(t *testing.T) {
	// Given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	helper := mocks.NewMockIHelper(ctrl)
	logger := mocks.NewMockILogger(ctrl)

	helper.EXPECT().LoadBootConfig().Return(&interfaces.Bootconfig{DatabasePath: "testdatabase"}).AnyTimes()
	helper.EXPECT().CheckOrCreateFolder(gomock.Any(), gomock.Any()).AnyTimes()
	helper.EXPECT().CheckOrCreateParentFolder(gomock.Any(), gomock.Any()).AnyTimes()
	helper.EXPECT().HomeDir(gomock.Any(), gomock.Any()).AnyTimes()

	logger.EXPECT().Tracef(gomock.Any(), gomock.Any()).AnyTimes()
	logger.EXPECT().GetLevel().AnyTimes()
	logger.EXPECT().Warnf(gomock.Any(), gomock.Any()).AnyTimes()

	dal := setupDatabase(helper, logger)
	defer tearDownDatabase(helper)

	// When
	search := &project.Project{}
	search.SetName("TESTProject")
	err := dal.Read(search)
	assert.Nil(t, err)
	search.SetName("edited-project")
	dal.Update(search)

	// Then

	err2 := dal.Read(search)
	assert.Nil(t, err2)

	assert.Equal(t, "edited-project", search.GetName())
}

func TestAddProjectToProfile(t *testing.T) {
	// Given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	helper := mocks.NewMockIHelper(ctrl)
	logger := mocks.NewMockILogger(ctrl)

	helper.EXPECT().LoadBootConfig().Return(&interfaces.Bootconfig{DatabasePath: "testdatabase"}).AnyTimes()
	helper.EXPECT().CheckOrCreateFolder(gomock.Any(), gomock.Any()).AnyTimes()
	helper.EXPECT().CheckOrCreateParentFolder(gomock.Any(), gomock.Any()).AnyTimes()
	helper.EXPECT().HomeDir(gomock.Any(), gomock.Any()).AnyTimes()

	logger.EXPECT().Tracef(gomock.Any(), gomock.Any()).AnyTimes()
	logger.EXPECT().GetLevel().AnyTimes()
	logger.EXPECT().Warnf(gomock.Any(), gomock.Any()).AnyTimes()

	dal := setupDatabase(helper, logger)
	defer tearDownDatabase(helper)

	// When
	searchProfile := &profile.Profile{}
	searchProfile.SetName("TESTProfile")
	err1 := dal.Read(searchProfile)
	assert.Nil(t, err1)
	searchProject := &project.Project{}
	searchProject.SetName("TESTProject")
	err2 := dal.Read(searchProject)
	assert.Nil(t, err2)

	searchProfile.AddProjectDirect(searchProject)
	dal.Update(searchProfile)

	// Then
	err3 := dal.Read(searchProfile)
	assert.Nil(t, err3)

	assert.NotNil(t, searchProfile.GetProjects())
	assert.Equal(t, "TESTProject", searchProfile.GetProjects()["TESTProject"].GetName())
}
