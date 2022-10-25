package project

import (
	"fmt"
	"io/fs"
	"os"
	"path"

	"gopkg.in/yaml.v3"

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

func defaultPermissions() fs.FileMode {
	return os.FileMode(uint32(0760))
}

const metadataFolder = ".grgd"

// AutosetupServices ...
func (p *Project) autoSetupServices(core interfaces.ICore) error {
	// log := core.GetLogger()
	ui := core.GetUI()
	h := core.GetHelper()
	p.settingsyamlpath = path.Join(p.path, metadataFolder, ".grgdproject.yaml")
	h.CheckOrCreateParentFolder(p.settingsyamlpath, defaultPermissions())
	projMeta, err := p.readSettings(h)

	shouldBeIgnored(metadataFolder, &projMeta.IgnoreFolders)
	files, err := os.ReadDir(p.path)
	if err != nil {
		return err
	}
	for _, serviceDir := range files {
		if !serviceDir.IsDir() {
			continue
		}
		if isIgnored(serviceDir.Name(), projMeta.IgnoreFolders) {
			continue
		}
		servMeta := &interfaces.ServiceMetadata{Name: serviceDir.Name()}

		servYamlPath := path.Join(p.path, serviceDir.Name(), ".grgdservice.yaml")
		dat, err := h.ReadFile(servYamlPath)
		if err != nil {
			dat, err = yaml.Marshal(servMeta)
			if err != nil {
				return err
			}
			if !ui.YesNoQuestion("Is " + serviceDir.Name() + " a folder containing a service you would like to add to the project?") {
				if ui.YesNoQuestion("Add " + serviceDir.Name() + " to the ignore list? (can be whitelisted again): ") {
					shouldBeIgnored(serviceDir.Name(), &projMeta.IgnoreFolders)
				}
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
		projMeta.Services[servMeta.Name] = interfaces.ServiceLocator{Path: servYamlPath, Active: true}
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

func shouldBeIgnored(val string, arr *[]string) {
	if !isIgnored(val, *arr) {
		*arr = append(*arr, val)
	}
}
func isIgnored(val string, arr []string) bool {
	for _, v := range arr {
		if val == v {
			return true
		}
	}
	return false
}

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
func (project *Project) GetID() uint {
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
func (project *Project) GetServices(i ...interface{}) map[string]interfaces.ServiceMetadata {
	services := make(map[string]interfaces.ServiceMetadata)
	c, ok := i[0].(interfaces.ICore)
	if !ok {
		return nil
	}
	log := c.GetLogger()
	h := c.GetHelper()
	projMeta, err := project.readSettings(h)
	if err != nil {
		return nil

	}
	for _, v := range projMeta.Services {
		byts, err := os.ReadFile(v.Path)
		if err != nil {
			log.Warnf("%s", err.Error())
			continue
		}
		srv := interfaces.ServiceMetadata{}
		err = yaml.Unmarshal(byts, &srv)
		if err != nil {
			log.Warnf("%s", err.Error())
			continue
		}
		services[srv.Name] = srv
	}
	return services
}

// GetServiceByName ...
func (project *Project) GetServiceByName(serviceName string, i ...interface{}) interfaces.ServiceMetadata {

	return interfaces.ServiceMetadata{}
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
