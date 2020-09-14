package wrappers

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"

	"grgd/interfaces"
	"grgd/persistence"

	"github.com/gregod-com/grgdplugincontracts"
)

// InitNewProfile ...
func InitNewProfile(
	model *persistence.Profile,
	logger interfaces.ILogger,
	UI grgdplugincontracts.IUIPlugin,
	fm interfaces.IFileSystemManipulator) *ProfileWrapper {
	// defaults for new profile
	model.HomeDir = fm.HomeDir()
	model.PluginDir = fm.HomeDir(".grgd", "plugins")

	UI.ClearScreen(nil)
	UI.Printf("Hey %v, let's init your profile \n\n", model.Name)
	UI.Question("Base profile directory [`Enter` for default: "+model.HomeDir+"]: ", &model.HomeDir)

	for !fm.PathExists(model.HomeDir) {
		answer := model.HomeDir
		model.HomeDir = fm.HomeDir()
		UI.Question("The path `"+answer+"` does not exists. Try again or use default ["+model.HomeDir+"]: ", model.HomeDir)
	}

	UI.Question("Base plugin directory [`Enter` for default: "+model.PluginDir+"]: ", &model.PluginDir)

	for !fm.PathExists(model.PluginDir) {
		answer := model.PluginDir
		model.PluginDir = fm.HomeDir(".grgd", "plugins")
		UI.Question("The path `"+answer+"` does not exists. Try again or use default ["+model.PluginDir+"]: ", &model.PluginDir)
	}

	// fmt.Println(model)
	// if !UI.YesNoQuestion(nil, "Looking good?") {
	// 	InitNewProfile(model, logger, UI)
	// }

	first := "your first "
	for UI.YesNoQuestion(nil, "Should we setup "+first+"project now?") {
		newprofileWrapper := InitNewProject(&persistence.GRGDProject{}, logger, UI, fm)
		model.Projects = append(model.Projects, newprofileWrapper.model)
		first = "another "
	}

	// CurrentProjectID uint
	model.Initialized = true
	return CreateProfile(model)
}

// InitNewProject ...
func InitNewProject(
	model *persistence.GRGDProject,
	logger interfaces.ILogger,
	UI grgdplugincontracts.IUIPlugin,
	fm interfaces.IFileSystemManipulator) *ProjectWrapper {
	var basepath string
	var name string

	UI.Printf("Let's init your project\n")
	UI.Question("What is the name of this project? ", &name)

	// UI.Question("Where is the base path of your project? ", &basepath)
	basepath = "/Users/gregor/iam"

	for !fm.PathExists(basepath) {
		if UI.YesNoQuestion(nil, "The path "+basepath+" does not seem to exists. Should we create the path now?") {
			fm.CheckOrCreateFolder(basepath, os.FileMode(uint32(0760)))
			continue
		}
		UI.Question("Where is the base path of your project? ", &basepath)
	}

	model.Name = name
	model.Path = basepath

	if UI.YesNoQuestion(nil, "Try to AUTOSETUP services?") {
		autosetup := AutosetupServices(model, logger, UI)
		model.Services = append(model.Services, autosetup...)
	}

	model.Initialized = true

	return CreateProjectWrapper(model)
}

// AutosetupServices ...
func AutosetupServices(model *persistence.GRGDProject, logger interfaces.ILogger, UI grgdplugincontracts.IUIPlugin) []*persistence.Service {

	content, err := ioutil.ReadFile(path.Join(model.Path, ".grgdproject.yaml"))
	if err == nil {
		logger.Info("Found project metadata")
		logger.Info(string(content))
		// TODO
		// Unmarschal and check yaml + ask user if all looks good
		return nil
	}

	files, err := ioutil.ReadDir(model.Path)
	if err != nil {
		return nil
	}
	for _, file := range files {
		if !file.IsDir() {
			continue
		}
		if !UI.YesNoQuestion(nil, "Is "+file.Name()+" a folder containing a service you would like to add to the project?") {
			continue
		}
		filepath.Walk(path.Join(model.Path, file.Name()), walker)
		logger.Info("Adding service " + file.Name())
	}

	logger.Fatal("Autosetup in " + model.Path)
	return nil
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
