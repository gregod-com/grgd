package wrappers

import (
	"sort"

	"grgd/interfaces"
	"grgd/persistence"
)

// ProfileWrapper ...
type ProfileWrapper struct {
	model    *persistence.Profile
	projects map[string]interfaces.IProject
}

// CreateProfile ...
func CreateProfile(mProfile *persistence.Profile) *ProfileWrapper {
	pros := make(map[string]interfaces.IProject)
	for k, mProj := range mProfile.Projects {
		// persistence.GetAll(&mProj)
		pros[mProj.Name] = CreateProjectWrapper(mProfile.Projects[k])
	}

	return &ProfileWrapper{model: mProfile, projects: pros}
}

// IsInitialized ...
func (p *ProfileWrapper) Model() interface{} {
	return &p.model
}

// IsInitialized ...
func (p *ProfileWrapper) IsInitialized() bool {
	return p.model.Initialized
}

// GetName ...
func (p *ProfileWrapper) GetName() string {
	return p.model.Name
}

// GetBasePath ...
func (p *ProfileWrapper) GetBasePath() string {
	return p.model.HomeDir
}

// GetProjects ....
func (p *ProfileWrapper) GetProjects() map[string]interfaces.IProject {
	return p.projects
}

// GetProjectsTable ....
func (p *ProfileWrapper) GetProjectsTable() [][]string {
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
		if p.model.CurrentProjectID == p.projects[key].GetID() {
			currentFlag = "*"
		}
		row = append(row, currentFlag)
		row = append(row, p.projects[key].GetValues()...)
		rows = append(rows, row)
	}

	return rows
}

// AddProject ...
func (p *ProfileWrapper) AddProject(proj string) error {
	newProj := &persistence.GRGDProject{Name: proj}
	p.projects[proj] = CreateProjectWrapper(newProj)
	p.model.Projects = append(p.model.Projects, newProj)
	return nil
}

// RemoveProject ...
func (p *ProfileWrapper) RemoveProject(proj interfaces.IProject) error {
	delete(p.projects, proj.GetName())
	return nil
}

// RemoveProjectByName ...
func (p *ProfileWrapper) RemoveProjectByName(proj string) error {
	delete(p.projects, proj)
	return nil
}

// GetCurrentProject ...
func (p *ProfileWrapper) GetCurrentProject() interfaces.IProject {
	for _, v := range p.projects {
		if v.GetID() == p.model.CurrentProjectID {
			return p.projects[v.GetName()]
		}
	}
	return nil
}

// SetCurrentProject ...
func (p *ProfileWrapper) SetCurrentProject(newProject interfaces.IProject) error {
	p.model.CurrentProjectID = newProject.GetID()
	return nil
}

// GetValues ...
func (p *ProfileWrapper) GetValues(i ...interface{}) []string {
	return []string{p.model.Name, p.model.HomeDir}
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
