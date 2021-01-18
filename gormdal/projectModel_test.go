package gormdal

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/gregod-com/grgd/interfaces/mocks"
	"github.com/tj/assert"
)

func TestLoadTESTProject(t *testing.T) {
	// Given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	dal := setupDatabase(mocks.NewMockIFileSystemManipulator(ctrl))
	defer tearDownDatabase(dal)

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
	dal := setupDatabase(mocks.NewMockIFileSystemManipulator(ctrl))
	defer tearDownDatabase(dal)

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
	dal := setupDatabase(mocks.NewMockIFileSystemManipulator(ctrl))
	defer tearDownDatabase(dal)

	// When
	search := &ProfileModel{Name: "TESTProject"}
	err := dal.Read(search)
	assert.Nil(t, err)
	search.Name = "edited-project"
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
	dal := setupDatabase(mocks.NewMockIFileSystemManipulator(ctrl))
	defer tearDownDatabase(dal)

	// When
	searchProfile := &ProfileModel{Name: "TESTProfile"}
	err1 := dal.Read(searchProfile)
	assert.Nil(t, err1)
	searchProject := &ProjectModel{Name: "TESTProject"}
	err2 := dal.Read(searchProject)
	assert.Nil(t, err2)

	searchProfile.Projects = append(searchProfile.Projects, searchProject)
	dal.Update(searchProfile)

	// Then
	err3 := dal.Read(searchProfile)
	assert.Nil(t, err3)

	assert.NotNil(t, searchProfile.Projects)
	assert.Equal(t, "TESTProject", searchProfile.Projects[0].Name)
}