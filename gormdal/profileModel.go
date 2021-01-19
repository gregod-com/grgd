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
	UpdateURL        string
	AWSRegion        string
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

// GetUpdateURL ...
func (profile *ProfileModel) GetUpdateURL() string {
	return profile.UpdateURL
}

// GetMetaMap ...
func (profile *ProfileModel) GetMetaMap() map[string]string {
	return map[string]string{
		"Name":      profile.Name,
		"UpdateURL": profile.UpdateURL,
		"HomeDir":   profile.HomeDir,
		"AWSRegion": profile.AWSRegion,
		"PluginDir": profile.PluginDir,
	}
}
