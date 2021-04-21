package profile

import (
	"fmt"
	"log"
	"net/http"
	"path"
	"sort"

	"github.com/gregod-com/grgd/interfaces"
	"github.com/gregod-com/grgd/pkg/project"
)

// ProvideProfile ...
func ProvideProfile(logger interfaces.ILogger, ui interfaces.IUIPlugin) interfaces.IProfile {
	p := &Profile{}
	p.logger = logger
	p.ui = ui
	logger.Tracef("provide %T", p)
	return p
}

// Profile ...
type Profile struct {
	logger           interfaces.ILogger
	ui               interfaces.IUIPlugin
	id               uint
	name             string
	initialized      bool
	projects         map[string]interfaces.IProject
	currentProjectID uint
	metadata         map[string]string
}

// InitNewProfile ...
func InitNewProfile(name string, ui interfaces.IUIPlugin, log interfaces.ILogger, helper interfaces.IHelper, i ...interface{}) interfaces.IProfile {

	var profile Profile
	// defaults for new profile
	profile.name = name
	profile.metadata = make(map[string]string)
	profile.metadata["homeDir"] = helper.HomeDir(".grgd")
	profile.metadata["hackDir"] = path.Join(profile.metadata["homeDir"], "hack")
	profile.metadata["pluginDir"] = path.Join(profile.metadata["homeDir"], "pluginsv2")
	profile.metadata["updateURL"] = updateurl
	profile.metadata["awsRegion"] = awsregion

	ui.ClearScreen()

	ui.Printf("Hey %v, let's init your profile\n\n", profile.name)

	// Scripts
	ui.Questionf("Base scripts directory [%s]: ", profile.metadata["pluginDir"], profile.metadata["pluginDir"])
	for !helper.PathExists(profile.metadata["pluginDir"]) {
		answer := profile.metadata["pluginDir"]
		profile.metadata["pluginDir"] = helper.HomeDir(".grgd", "hack")
		ui.Questionf("The path `%s` does not exists. Try again or use default [%s]: ", answer, profile.metadata["pluginDir"], profile.metadata["pluginDir"])
	}

	ui.Questionf("URL to fetch updates from: [%s]: ", profile.metadata["updateURL"], profile.metadata["updateURL"])

	for !ping(profile.metadata["updateURL"]) {
		answer := profile.metadata["updateURL"]
		profile.metadata["updateURL"] = updateurl
		ui.Questionf("The url `%s` it not reachable. Use anyways or use default [%s]: ", answer, profile.metadata["updateURL"], profile.metadata["updateURL"])
	}

	profile.projects = make(map[string]interfaces.IProject)

	profile.initialized = true
	return &profile
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

// GetMetaData ...
func (p *Profile) GetMetaData(key string) string {
	switch key {
	case "Name":
		return p.name
	default:
		if val, ok := p.metadata[key]; ok {
			return val
		}
		return ""
	}
}

// SetMetaData ...
func (p *Profile) SetMetaData(key, value string) {
	if p.metadata == nil {
		p.metadata = make(map[string]string)
	}
	switch key {
	case "Name":
		p.name = value
	default:
		p.metadata[key] = value
	}
}

// GetUpdateURL ...
func (p *Profile) GetUpdateURL() string {
	if url, ok := p.metadata["updateURL"]; ok {
		return url
	}
	return ""
}

// IsInitialized ...
func (p *Profile) IsInitialized() bool {
	return p.initialized
}

// SetInitialized ...
func (p *Profile) SetInitialized(init bool) error {
	p.initialized = init
	return nil
}

// GetID ...
func (p *Profile) GetID() uint {
	return p.id
}

// SetID ...
func (p *Profile) SetID(id uint) error {
	p.id = id
	return nil
}

// GetName ...
func (p *Profile) GetName() string {
	return p.name
}

// SetName ...
func (p *Profile) SetName(n string) error {
	p.name = n
	return nil
}

// GetBasePath ...
func (p *Profile) GetBasePath() string {
	return p.metadata["homeDir"]
}

// GetPluginsDir ...
func (p *Profile) GetPluginsDir() string {
	return p.metadata["pluginDir"]
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
		if p.currentProjectID == p.projects[key].GetID() {
			currentFlag = "*"
		}
		row = append(row, currentFlag)
		row = append(row, p.projects[key].GetValues()...)
		rows = append(rows, row)
	}

	return rows
}

func (p *Profile) AddProject() error {
	return nil
}

func (p *Profile) AddProjectByName(name string) error {
	if p.projects == nil {
		p.projects = make(map[string]interfaces.IProject)
	}
	if _, ok := p.projects[name]; !ok {
		newP := &project.Project{}
		newP.SetName(name)
		p.projects[name] = newP
	}
	return nil
}

// AddProject ...
func (p *Profile) AddProjectDirect(proj interfaces.IProject) error {
	if p.projects == nil {
		p.projects = make(map[string]interfaces.IProject)
	}
	p.projects[proj.GetName()] = proj
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
func (p *Profile) GetCurrentProjectID() uint {
	return p.currentProjectID
}

// GetCurrentProject ...
func (p *Profile) SetCurrentProjectID(id uint) error {
	p.currentProjectID = id
	return nil

}

// GetCurrentProject ...
func (p *Profile) GetCurrentProject() interfaces.IProject {
	for _, v := range p.projects {
		if v.GetID() == p.currentProjectID {
			return p.projects[v.GetName()]
		}
	}
	return nil
}

// SetCurrentProject ...
func (p *Profile) SetCurrentProject(newProject interfaces.IProject) error {
	p.currentProjectID = newProject.GetID()
	return nil
}

// GetValues ...
func (p *Profile) GetValues(i ...interface{}) []string {
	retSlice := []string{}
	for k, v := range p.metadata {
		retSlice = append(retSlice, fmt.Sprintf("%s: %s", k, v))
	}
	return retSlice
}

func (p *Profile) GetValuesAsMap(i ...interface{}) map[string]string {
	return p.metadata
}
