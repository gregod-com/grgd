package gormdal

import (
	"github.com/gregod-com/grgd/interfaces"
	"gorm.io/gorm"
)

func projectModelToIProject(in *ProjectModel, out interfaces.IProject) error {
	out.SetID(in.ID)
	out.SetName(in.Name)
	out.SetPath(in.Path)
	out.SetSettingsYamlPath(in.SettingsYamlPath)
	out.SetInitialized(in.Initialized)
	return nil
}

func iprojectToProjectModel(in interfaces.IProject, out *ProjectModel) error {
	out.ID = in.GetID()
	out.Name = in.GetName()
	out.Initialized = in.IsInitialized()
	out.Path = in.GetPath()
	out.SettingsYamlPath = in.GetSettingsYamlPath()
	return nil
}

// ProjectModel ...
type ProjectModel struct {
	gorm.Model
	Name             string `gorm:"unique;not null;default:null"`
	Path             string
	SettingsYamlPath string
	Initialized      bool
	Description      string
	Profiles         []*ProfileModel `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;;many2many:profiles_projects;"`
}

// ProjectTag ...
type ProjectTag struct {
	gorm.Model
	Name string
}

// GetID ...
func (p *ProjectModel) GetID() uint {
	return p.ID
}

// GetName ...
func (p *ProjectModel) GetName() string {
	return p.Name
}
