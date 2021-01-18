package gormdal

import (
	"gorm.io/gorm"
)

// ProfileModel ...
type ProfileModel struct {
	gorm.Model
	Name             string `gorm:"unique"`
	HomeDir          string
	PluginDir        string
	Projects         []*ProjectModel `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CurrentProjectID uint
	Initialized      bool
}

// IsInitialized ...
func (profile *ProfileModel) IsInitialized() bool {
	return profile.Initialized
}

// GetName ...
func (profile *ProfileModel) GetName() string {
	return profile.Name
}

// GetBasePath ...
func (profile *ProfileModel) GetBasePath() string {
	return profile.HomeDir
}

// GetCurrentProjectID ...
func (profile *ProfileModel) GetCurrentProjectID() uint {
	return profile.CurrentProjectID
}
