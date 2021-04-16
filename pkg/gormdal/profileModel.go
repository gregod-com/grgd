package gormdal

import (
	"github.com/gregod-com/grgd/interfaces"
	"github.com/gregod-com/grgd/pkg/project"
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
	out.SetCurrentProjectID(in.CurrentProjectID)
	for _, proj := range in.Projects {
		// todo make implementation agnostic?
		iproj := &project.Project{}
		projectModelToIProject(proj, iproj)
		out.AddProjectDirect(iproj)
	}

	out.SetInitialized(in.Initialized)
	return nil
}

// iprofileToProfileModel ...
func iprofileToProfileModel(in interfaces.IProfile, out *ProfileModel) error {
	out.ID = in.GetID()
	out.Name = in.GetName()
	out.HomeDir = in.GetMetaData("homeDir")
	out.HackDir = in.GetMetaData("hackDir")
	out.PluginDir = in.GetMetaData("pluginDir")
	out.UpdateURL = in.GetMetaData("updateURL")
	out.AWSRegion = in.GetMetaData("awsRegion")
	out.CurrentProjectID = in.GetCurrentProjectID()
	for _, proj := range in.GetProjects() {
		pmdl := &ProjectModel{}
		iprojectToProjectModel(proj, pmdl)
		out.Projects = append(out.Projects, pmdl)
	}
	out.Initialized = in.IsInitialized()
	return nil
}

// ProfileModel ...
type ProfileModel struct {
	gorm.Model
	Name             string `gorm:"unique;not null;default:null"`
	HomeDir          string
	HackDir          string
	PluginDir        string
	Projects         []*ProjectModel `gorm:"constraint:OnUpdate:UPDATE,OnDelete:SET NULL;;many2many:profiles_projects;"`
	CurrentProjectID uint
	Initialized      bool
	UpdateURL        string
	AWSRegion        string
}

// GetID ...
func (profile *ProfileModel) GetID() uint {
	return profile.ID
}

// GetName ...
func (profile *ProfileModel) GetName() string {
	return profile.Name
}
