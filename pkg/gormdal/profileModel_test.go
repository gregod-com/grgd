package gormdal

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/gregod-com/grgd/interfaces"
	"github.com/gregod-com/grgd/interfaces/mocks"
	"github.com/gregod-com/grgd/pkg/profile"
	"github.com/stretchr/testify/assert"
)

func TestLoadTESTProfile(t *testing.T) {
	// Given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	// testBootConfig :=
	helper := mocks.NewMockIHelper(ctrl)
	logger := mocks.NewMockILogger(ctrl)

	helper.EXPECT().LoadBootConfig().Return(&interfaces.Bootconfig{DatabasePath: "testdatabase"}).AnyTimes()
	helper.EXPECT().CheckOrCreateFolder(gomock.Any(), gomock.Any()).AnyTimes()
	helper.EXPECT().CheckOrCreateParentFolder(gomock.Any(), gomock.Any()).AnyTimes()
	helper.EXPECT().HomeDir(gomock.Any(), gomock.Any()).AnyTimes()
	helper.EXPECT().HomeDir(gomock.Any(), gomock.Any()).AnyTimes()
	logger.EXPECT().Tracef(gomock.Any(), gomock.Any()).AnyTimes()
	logger.EXPECT().GetLevel().AnyTimes()

	dal := setupDatabase(helper, logger)
	defer tearDownDatabase(helper)

	// When
	search := &profile.Profile{}
	search.SetName("TESTProfile")
	err := dal.Read(search)
	assert.Nil(t, err)

	// Then
	assert.Equal(t, "TESTProfile", search.GetName())
	assert.Equal(t, "./test-me-dir/", search.GetBasePath())
}

func TestDeleteTESTProfile(t *testing.T) {
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
	dal := setupDatabase(helper, logger)
	defer tearDownDatabase(helper)

	// When
	// When
	search := &profile.Profile{}
	search.SetName("TESTProfile")
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
	helper := mocks.NewMockIHelper(ctrl)
	logger := mocks.NewMockILogger(ctrl)

	helper.EXPECT().LoadBootConfig().Return(&interfaces.Bootconfig{DatabasePath: "testdatabase"}).AnyTimes()
	helper.EXPECT().CheckOrCreateFolder(gomock.Any(), gomock.Any()).AnyTimes()
	helper.EXPECT().CheckOrCreateParentFolder(gomock.Any(), gomock.Any()).AnyTimes()
	helper.EXPECT().HomeDir(gomock.Any(), gomock.Any()).AnyTimes()
	logger.EXPECT().Tracef(gomock.Any(), gomock.Any()).AnyTimes()
	logger.EXPECT().GetLevel().AnyTimes()
	dal := setupDatabase(helper, logger)
	defer tearDownDatabase(helper)

	// When
	// When
	search := &profile.Profile{}
	search.SetName("TESTProfile")
	err := dal.Read(search)
	assert.Nil(t, err)
	search.SetName("edited-name")
	dal.Update(search)

	// Then

	err2 := dal.Read(search)
	assert.Nil(t, err2)

	assert.Equal(t, "edited-name", search.GetName())
}
