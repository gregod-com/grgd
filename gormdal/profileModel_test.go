package gormdal

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/gregod-com/grgd/interfaces/mocks"
	"github.com/tj/assert"
)

func TestLoadTESTProfile(t *testing.T) {
	// Given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	dal := setupDatabase(mocks.NewMockIFileSystemManipulator(ctrl))
	defer tearDownDatabase(dal)

	// When
	search := ProfileModel{Name: "TESTProfile"}
	err := dal.Read(&search)
	assert.Nil(t, err)

	// Then
	assert.Equal(t, "TESTProfile", search.Name)
	assert.Equal(t, "TESTProfile", search.GetName())
	assert.Equal(t, "./test-me-dir/", search.HomeDir)
	assert.Equal(t, "./test-me-dir/", search.GetBasePath())

}

func TestDeleteTESTProfile(t *testing.T) {
	// Given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	dal := setupDatabase(mocks.NewMockIFileSystemManipulator(ctrl))
	defer tearDownDatabase(dal)

	// When
	search := &ProfileModel{Name: "TESTProfile"}
	err := dal.Read(search)
	assert.Nil(t, err)
	dal.Delete(search)

	// Then
	err2 := dal.Read(search)
	assert.Error(t, err2)
}

func TestEditTESTProfile(t *testing.T) {
	// Given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	dal := setupDatabase(mocks.NewMockIFileSystemManipulator(ctrl))
	defer tearDownDatabase(dal)

	// When
	search := &ProfileModel{Name: "TESTProfile"}
	err := dal.Read(search)
	assert.Nil(t, err)
	search.Name = "edited-name"
	dal.Update(search)

	// Then

	err2 := dal.Read(search)
	assert.Nil(t, err2)

	assert.Equal(t, "edited-name", search.GetName())
}
