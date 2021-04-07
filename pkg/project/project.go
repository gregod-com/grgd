package project

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"

	GI "github.com/gregod-com/grgd/interfaces"
)

// ProvideProject ...
func ProvideProject(
	logger GI.ILogger,
	ui GI.IUIPlugin,
	helper GI.IHelper,
	services map[string]GI.IService) GI.IProject {
	return &Project{logger: logger, ui: ui, helper: helper, services: services}
}

// Project ...
type Project struct {
	id          uint
	name        string
	path        string
	initialized bool
	description string
	services    map[string]GI.IService
	logger      GI.ILogger
	ui          GI.IUIPlugin
	helper      GI.IHelper
}

// Init ...
func (p *Project) Init() error {
	var name, basepath string
	p.ui.Printf("Let's init your project\n")
	p.ui.Question("What is the name of this project? ", &name)

	// TODO: replace static with actual question
	p.ui.Question("Where is the base path of your project? ", &basepath)

	for !p.helper.PathExists(basepath) {
		if p.ui.YesNoQuestion("The path " + basepath + " does not seem to exists. Should we create the path now?") {
			p.helper.CheckOrCreateFolder(basepath, os.FileMode(uint32(0760)))
			continue
		}
		p.ui.Question("Where is the base path of your project? ", &basepath)
	}

	p.name = name
	p.path = basepath

	if p.ui.YesNoQuestion("Try to AUTOSETUP services?") {
		p.autoSetupServices()
	}

	p.initialized = true
	return nil
}

// AutosetupServices ...
func (p *Project) autoSetupServices() {
	content, err := ioutil.ReadFile(path.Join(p.path, ".grgdproject.yaml"))
	if err == nil {
		p.logger.Info("Found project metadata")
		p.logger.Info(string(content))
		// TODO
		// Unmarschal and check yaml + ask user if all looks good
	}

	files, err := ioutil.ReadDir(p.path)
	if err != nil {
		return
	}
	for _, file := range files {
		if !file.IsDir() {
			continue
		}
		if !p.ui.YesNoQuestion("Is " + file.Name() + " a folder containing a service you would like to add to the project?") {
			continue
		}
		filepath.Walk(path.Join(p.path, file.Name()), walker)
		p.logger.Info("Adding service " + file.Name())
	}

	p.logger.Fatal("Autosetup in " + p.path)
	return
}

func walker(path string, info os.FileInfo, err error) error {
	if !info.IsDir() && strings.Contains(info.Name(), "docker-compose") {
		fmt.Println(path)
	}

	if !info.IsDir() && strings.Contains(info.Name(), "skaffold.yaml") {
		fmt.Println(path)
	}

	if info.IsDir() && strings.Contains(info.Name(), "Chart") {
		fmt.Println(path)
	}

	return nil
}

// GetName ...
func (project *Project) GetName() string {
	return project.name
}

// GetID ...
func (project *Project) GetID(i ...interface{}) uint {
	return project.id
}

// GetPath ...
func (project *Project) GetPath(i ...interface{}) string {
	return project.path
}

// SetPath ...
func (project *Project) SetPath(path string, i ...interface{}) error {
	project.path = path
	return nil
}

// GetServices ...
func (project *Project) GetServices(i ...interface{}) map[string]GI.IService {
	return project.services
}

// GetServiceByName ...
func (project *Project) GetServiceByName(serviceName string, i ...interface{}) GI.IService {
	return project.services[serviceName]
}

// GetValues ...
func (project *Project) GetValues(i ...interface{}) []string {
	return []string{project.name, project.path, project.description}
}

// // Edit ...
// func (proj *GRGDProject) Edit(db *gorm.DB, i ...interface{}) error {
// 	if !p.ui.YesNoQuestion("Edit project `"+proj.Name+"` now?") {
// 		return nil
// 	}

// 	p.ui.Question("Project name ["+proj.Name+"]: ", &proj.Name)
// 	p.ui.Question("Project path (absolute)["+proj.Path+"]: ", &proj.Path)
// 	p.ui.Question("Project description ["+proj.Description+"]:", &proj.Description)

// 	for k := range proj.Services {
// 		Edit(&proj.Services[k])
// 	}

// 	p.ui.Println(proj)
// 	if !p.ui.YesNoQuestion("Looking good?") {
// 		proj.Edit(db)
// 	}

// 	return nil
// }

// // String  ...
// func (proj GRGDProject) String() string {
// 	var obj map[string]interface{}
// 	// create json string from object
// 	str, err := json.MarshalIndent(proj, "", "  ")

// 	// create simplified object from json string
// 	json.Unmarshal([]byte(str), &obj)

// 	f := colorjson.NewFormatter()
// 	f.Indent = 4

// 	// create colored json string from simplified object
// 	data, err := f.Marshal(obj)

// 	return string(data)
// }

// // IsInitialized ...
// func (proj *GRGDProject) IsInitialized(i ...interface{}) bool {
// 	return proj.Initialized
// }
