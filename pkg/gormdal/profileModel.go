package gormdal

import (
	"github.com/gregod-com/grgd/interfaces"
	"gorm.io/gorm"
)

// ProfileModelToIProfile ...
func profileModelToIProfile(in *ProfileModel, out interfaces.IProfile) error {
	out.SetID(in.ID)
	out.SetName(in.Name)

	out.SetMetaData("homeDir", in.HomeDir)
	out.SetMetaData("hackDir", in.HackDir)
	out.SetMetaData("pluginDir", in.PluginDir)
	out.SetMetaData("updateURL", in.UpdateURL)
	out.SetMetaData("awsRegion", in.AWSRegion)
	out.SetInitialized(in.Initialized)
	return nil
}

// profileToIProfileModel ...
func profileToIProfileModel(in interfaces.IProfile, out *ProfileModel) error {
	out.ID = in.GetID()
	out.Name = in.GetName()
	out.HomeDir = in.GetMetaData("homeDir")
	out.HackDir = in.GetMetaData("hackDir")
	out.PluginDir = in.GetMetaData("pluginDir")
	out.UpdateURL = in.GetMetaData("updateURL")
	out.AWSRegion = in.GetMetaData("awsRegion")
	out.Initialized = in.IsInitialized()
	return nil
}

// ProfileModel ...
type ProfileModel struct {
	gorm.Model
	Name             string `gorm:"unique"`
	HomeDir          string
	HackDir          string
	PluginDir        string
	Projects         []*ProjectModel `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CurrentProjectID uint
	Initialized      bool
	UpdateURL        string
	AWSRegion        string
}

// // IsInitialized ...
// func (profile *ProfileModel) IsInitialized() bool {
// 	return profile.Initialized
// }

// GetID ...
func (profile *ProfileModel) GetID() uint {
	return profile.ID
}

// GetName ...
func (profile *ProfileModel) GetName() string {
	return profile.Name
}

// // GetBasePath ...
// func (profile *ProfileModel) GetBasePath() string {
// 	return profile.HomeDir
// }

// // GetCurrentProjectID ...
// func (profile *ProfileModel) GetCurrentProjectID() uint {
// 	return profile.CurrentProjectID
// }

// // GetUpdateURL ...
// func (profile *ProfileModel) GetUpdateURL() string {
// 	return profile.UpdateURL
// }
