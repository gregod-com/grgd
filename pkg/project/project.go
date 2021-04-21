package project

import (
	"fmt"
	"io/fs"
	"os"
	"path"

	"gopkg.in/yaml.v2"

	"github.com/gregod-com/grgd/interfaces"
)

// ProvideProject ...
func ProvideProject() interfaces.IProject {
	return &Project{}
}

// Project ...
type Project struct {
	id               uint
	name             string
	path             string
	initialized      bool
	description      string
	settingsyamlpath string
	// services    map[string]interfaces.IService
}

// Init ...
func (p *Project) Init(core interfaces.ICore) error {
	helper := core.GetHelper()
	ui := core.GetUI()
	var name, basepath string
	var err error
	if p.name != "" {
		name = p.name
	}
	if p.path != "" {
		basepath = p.path
	} else {
		basepath = helper.CurrentWorkdir()
	}

	ui.Printf("Let's init your project\n")
	ui.Questionf("What is the name of this project? %s: ", &name, name)

	ui.Questionf("Where is the base path of this project? %s: ", &basepath, basepath)
	for !helper.PathExists(basepath) {
		if ui.YesNoQuestion("The path " + basepath + " does not seem to exists. Should we create the path now?") {
			helper.CheckOrCreateFolder(basepath, os.FileMode(uint32(0760)))
			continue
		}
		ui.Questionf("Where is the base path of your project? %s: ", &basepath, basepath)
	}

	p.name = name
	p.path = basepath

	if ui.YesNoQuestion("Try to AUTOSETUP services?") {
		// TODO autosetup should be injected
		err = p.autoSetupServices(core)
	}

	p.initialized = true
	return err
}

type ServiceMetadata struct {
	Name    string
	Version string
	// define how service is deployed for dev/staging/live...
	RunTimes map[string]RunTime
	//optional
	Meta map[string]interface{}
}

type RunTime struct {
	// define technology type (i.e. kubernetes, vm, serverless....)
	Technology string
	// define technology for dev deployment (i.e. skaffold&helm, skaffold & manifests, kubectl, helm-only,...)
	Helper map[string]interface{}
}

func defaultPermissions() fs.FileMode {
	return os.FileMode(uint32(0760))
}

// AutosetupServices ...
func (p *Project) autoSetupServices(core interfaces.ICore) error {
	// log := core.GetLogger()
	ui := core.GetUI()
	h := core.GetHelper()
	p.settingsyamlpath = path.Join(p.path, ".grgd", ".grgdproject.yaml")
	h.CheckOrCreateParentFolder(p.settingsyamlpath, defaultPermissions())
	projMeta, err := p.readSettings(h)

	files, err := os.ReadDir(p.path)
	if err != nil {
		return err
	}
	for _, serviceDir := range files {
		if !serviceDir.IsDir() {
			continue
		}
		if serviceDir.Name() == ".grgd" {
			continue
		}
		servMeta := &ServiceMetadata{Name: serviceDir.Name()}

		servYamlPath := path.Join(p.path, serviceDir.Name(), ".grgdservice.yaml")
		dat, err := h.ReadFile(servYamlPath)
		if err != nil {
			dat, err = yaml.Marshal(servMeta)
			if err != nil {
				return err
			}
			if !ui.YesNoQuestion("Is " + serviceDir.Name() + " a folder containing a service you would like to add to the project?") {
				continue
			}
			err = h.UpdateOrWriteFile(servYamlPath, dat, defaultPermissions())
			if err != nil {
				return err
			}
		}
		err = yaml.Unmarshal(dat, servMeta)
		if err != nil {
			return err
		}
		projMeta.Services[servMeta.Name] = map[string]interface{}{"path": servYamlPath, "active": true}
	}

	p.writeSettings(projMeta, h)

	return nil
}

func (p *Project) readSettings(h interfaces.IHelper) (*interfaces.ProjectMetadata, error) {
	projMeta := &interfaces.ProjectMetadata{Name: p.name}

	dat, err := h.ReadFile(p.settingsyamlpath)
	if err != nil {
		dat, err = yaml.Marshal(projMeta)
		if err != nil {
			return projMeta, err
		}
		err = h.UpdateOrWriteFile(p.settingsyamlpath, dat, defaultPermissions())
		if err != nil {
			return projMeta, err
		}
	}
	return projMeta, yaml.Unmarshal(dat, projMeta)
}

func (p *Project) writeSettings(projMeta *interfaces.ProjectMetadata, h interfaces.IHelper) error {
	dat, err := yaml.Marshal(projMeta)
	if err != nil {
		return err
	}

	err = h.UpdateOrWriteFile(p.settingsyamlpath, dat, defaultPermissions())
	if err != nil {
		return err
	}
	return nil
}

// func walker(path string, info os.FileInfo, err error) error {
// 	if !info.IsDir() && strings.Contains(info.Name(), "docker-compose") {
// 		fmt.Println(path)
// 	}

// 	if !info.IsDir() && strings.Contains(info.Name(), "skaffold.yaml") {
// 		fmt.Println(path)
// 	}

// 	if info.IsDir() && strings.Contains(info.Name(), "Chart") {
// 		fmt.Println(path)
// 	}

// 	return nil
// }

// GetName ...
func (project *Project) GetName() string {
	return project.name
}

// SetName ...
func (project *Project) SetName(name string) error {
	project.name = name
	return nil
}

// SetID ...
func (project *Project) SetID(id uint) error {
	project.id = id
	return nil

}

// GetID ...
func (project *Project) GetID(i ...interface{}) uint {
	return project.id
}

// SetInitialized ...
func (project *Project) SetInitialized(init bool) error {
	project.initialized = init
	return nil
}

// IsInitialized ...
func (project *Project) IsInitialized() bool {
	return project.initialized
}

// GetPath ...
func (project *Project) GetPath() string {
	return project.path
}

// SetPath ...
func (project *Project) SetPath(path string, i ...interface{}) error {
	project.path = path
	return nil
}

// GetServices ...
func (project *Project) GetServices(i ...interface{}) map[string]interfaces.IService {
	return nil
}

// GetServiceByName ...
func (project *Project) GetServiceByName(serviceName string, i ...interface{}) interfaces.IService {
	return nil
}

// GetValues ...
func (project *Project) GetValues(i ...interface{}) []string {
	return []string{project.name, project.path, project.description}
}

// SetSettingsYamlPath ...
func (project *Project) SetSettingsYamlPath(path string, i ...interface{}) error {
	project.settingsyamlpath = path
	return nil
}

// GetSettingsYamlPath ...
func (project *Project) GetSettingsYamlPath(i ...interface{}) string {
	return project.settingsyamlpath
}

// SetSettingsObject ...
func (project *Project) WriteSettingsObject(h interfaces.IHelper, i ...interface{}) error {
	ps, ok := i[0].(*interfaces.ProjectMetadata)
	if !ok {
		return fmt.Errorf("unsupported settings object")
	}
	project.writeSettings(ps, h)
	return nil
}

// GetSettingsObject ...
func (project *Project) ReadSettingsObject(h interfaces.IHelper, i ...interface{}) (*interfaces.ProjectMetadata, error) {
	return project.readSettings(h)
}
