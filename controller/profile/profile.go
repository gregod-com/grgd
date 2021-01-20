package profile

import (
	"log"
	"net/http"
	"sort"

	"github.com/gregod-com/grgd/gormdal"
	"github.com/gregod-com/grgd/interfaces"
)

// ProvideProfile ...
func ProvideProfile(id uint, projects map[string]interfaces.IProject) interfaces.IProfile {
	return &Profile{id: id, projects: projects}
}

// Profile ...
type Profile struct {
	id       uint
	model    interfaces.IProfileModel
	projects map[string]interfaces.IProject
}

var updateurl = "https://s3.iamstudent.dev/public/grgd/index.yaml"

// InitNewProfile ...
func InitNewProfile(
	name string,
	logger interfaces.ILogger,
	UI interfaces.IUIPlugin,
	fm interfaces.IFileSystemManipulator) *Profile {

	var profileModel gormdal.ProfileModel
	// defaults for new profile
	profileModel.Name = name
	profileModel.HomeDir = fm.HomeDir(".grgd")
	profileModel.PluginDir = fm.HomeDir(".grgd", "hack")
	profileModel.UpdateURL = updateurl
	profileModel.AWSRegion = "eu-central-1"

	UI.ClearScreen()
	UI.Printf("Hey %v, let's init your profile \n\n", profileModel.Name)
	UI.Questionf("Base grgd directory [%s]: ", profileModel.HomeDir, profileModel.HomeDir)

	for !fm.PathExists(profileModel.HomeDir) {
		answer := profileModel.HomeDir
		profileModel.HomeDir = fm.HomeDir()
		UI.Questionf(
			"The path `%s` does not exists. Try again or use default [%s]: ",
			answer,
			profileModel.HomeDir,
			profileModel.HomeDir)
	}

	UI.Questionf("Base scripts directory [%s]: ", &profileModel.PluginDir, profileModel.PluginDir)

	for !fm.PathExists(profileModel.PluginDir) {
		answer := profileModel.PluginDir
		profileModel.PluginDir = fm.HomeDir(".grgd", "hack")
		UI.Questionf("The path `%s` does not exists. Try again or use default [%s]: ", answer, &profileModel.PluginDir, profileModel.PluginDir)
	}

	UI.Questionf("URL to fetch updates from: [%s]: ", &profileModel.UpdateURL, profileModel.UpdateURL)

	for !ping(profileModel.UpdateURL) {
		answer := profileModel.UpdateURL
		profileModel.UpdateURL = updateurl
		UI.Questionf("The url `%s` it not reachable. Use anyways or use default [%s]: ", answer, profileModel.UpdateURL, profileModel.UpdateURL)
	}

	profileModel.Initialized = true
	return CreateProfile(&profileModel)
}

func ping(url string) bool {
	log.Println("checking " + url)
	resp, err := http.Get(url)
	log.Println("got " + resp.Status)
	if err != nil || resp.StatusCode != 200 {
		return false
	}
	return true
}

// CreateProfile ...
func CreateProfile(profileModel interfaces.IProfileModel) *Profile {
	// pros := make(map[string]interfaces.IProject)
	// for k, mProj := range mProfile.Projects {
	// persistence.GetAll(&mProj)
	// pros[mProj.Name] = CreateProjectWrapper(mProfile.Projects[k])
	// }
	return &Profile{model: profileModel}
}

// GetMetaMap ...
func (p *Profile) GetMetaMap() map[string]string {
	return p.model.GetMetaMap()
}

// Model ...
func (p *Profile) Model() interfaces.IProfileModel {
	return p.model
}

// GetUpdateURL ...
func (p *Profile) GetUpdateURL() string {
	return p.model.GetUpdateURL()
}

// IsInitialized ...
func (p *Profile) IsInitialized() bool {
	return p.model.IsInitialized()
}

// GetName ...
func (p *Profile) GetName() string {
	return p.model.GetName()
}

// GetBasePath ...
func (p *Profile) GetBasePath() string {
	return p.model.GetBasePath()
}

// GetProjects ....
func (p *Profile) GetProjects() map[string]interfaces.IProject {
	return p.projects
}

// GetProjectsTable ....
func (p *Profile) GetProjectsTable() [][]string {
	rows := [][]string{}
	keys := []string{}
	for k := range p.projects {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	rows = append(rows, []string{"Current", "Name", "Path", "Description"})
	for _, key := range keys {
		row := []string{}
		currentFlag := ""
		if p.model.GetCurrentProjectID() == p.projects[key].GetID() {
			currentFlag = "*"
		}
		row = append(row, currentFlag)
		row = append(row, p.projects[key].GetValues()...)
		rows = append(rows, row)
	}

	return rows
}

// AddProject ...
func (p *Profile) AddProject(proj string) error {
	// newProj := &persistence.GRGDProject{Name: proj}
	// p.projects[proj] = CreateProjectWrapper(newProj)
	// p.model.Projects = append(p.model.Projects, newProj)
	return nil
}

// RemoveProject ...
func (p *Profile) RemoveProject(proj interfaces.IProject) error {
	delete(p.projects, proj.GetName())
	return nil
}

// RemoveProjectByName ...
func (p *Profile) RemoveProjectByName(proj string) error {
	delete(p.projects, proj)
	return nil
}

// GetCurrentProject ...
func (p *Profile) GetCurrentProject() interfaces.IProject {
	// for _, v := range p.projects {
	// 	if v.GetID() == p.model.CurrentProjectID {
	// 		return p.projects[v.GetName()]
	// 	}
	// }
	return nil
}

// SetCurrentProject ...
func (p *Profile) SetCurrentProject(newProject interfaces.IProject) error {
	// p.model.GetCurrentProjectID() = newProject.GetID()
	return nil
}

// GetValues ...
func (p *Profile) GetValues(i ...interface{}) []string {
	return nil
	// return []string{p.name, p.homeDir}
}

// // String  ...

// func (profile Profile) String() string {
// 	var obj map[string]interface{}
// 	// create json string from object
// 	str, err := json.MarshalIndent(profile, "", "  ")

// 	// create simplified object from json string
// 	json.Unmarshal([]byte(str), &obj)

// 	f := colorjson.NewFormatter()
// 	f.Indent = 4

// 	// create colored json string from simplified object
// 	data, err := f.Marshal(obj)

// 	return string(data)
// }
