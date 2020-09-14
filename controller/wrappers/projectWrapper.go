package wrappers

import (
	"grgd/persistence"

	"grgd/interfaces"
)

// ProjectWrapper ...
type ProjectWrapper struct {
	model    *persistence.GRGDProject
	services map[string]interfaces.IService
}

// CreateProjectWrapper ...
func CreateProjectWrapper(model *persistence.GRGDProject) *ProjectWrapper {
	svc := make(map[string]interfaces.IService)
	for _, s := range model.Services {
		// persistence.GetAll(&s)
		svc[s.Name] = CreateServiceWrapper(s)
	}
	return &ProjectWrapper{model: model, services: svc}
}

// GetName ...
func (project *ProjectWrapper) GetName() string {
	return project.model.Name
}

// GetID ...
func (project *ProjectWrapper) GetID(i ...interface{}) uint {
	return project.model.ID
}

// GetPath ...
func (project *ProjectWrapper) GetPath(i ...interface{}) string {
	return project.model.Path
}

// SetPath ...
func (project *ProjectWrapper) SetPath(path string, i ...interface{}) error {
	project.model.Path = path
	return nil
}

// GetServices ...
func (project *ProjectWrapper) GetServices(i ...interface{}) map[string]interfaces.IService {
	return project.services
}

// GetServiceByName ...
func (project *ProjectWrapper) GetServiceByName(serviceName string, i ...interface{}) interfaces.IService {
	return project.services[serviceName]
}

// GetValues ...
func (project *ProjectWrapper) GetValues(i ...interface{}) []string {
	return []string{project.model.Name, project.model.Path, project.model.Description}
}

// // Init ...
// func (proj *GRGDProject) Init(db *gorm.DB, i ...interface{}) error {
// 	if !UI.YesNoQuestion(nil, "Init project "+proj.Name+" now?") {
// 		return nil
// 	}

// 	UI.Question("Project name ["+proj.Name+"]: ", &proj.Name)
// 	UI.Question("Project path (absolute)["+proj.Path+"]: ", &proj.Path)
// 	UI.Question("Project description ["+proj.Description+"]:", &proj.Description)

// 	for UI.YesNoQuestion(nil, "Do you want to add (another) service to the project?") {
// 		var service Service
// 		if service.Init(db) != nil {
// 			log.Error("Something went wrong when adding service")
// 		}
// 		proj.Services = append(proj.Services, service)

// 	}

// 	proj.Initialized = true
// 	return nil
// }

// // Edit ...
// func (proj *GRGDProject) Edit(db *gorm.DB, i ...interface{}) error {
// 	if !UI.YesNoQuestion(nil, "Edit project `"+proj.Name+"` now?") {
// 		return nil
// 	}

// 	UI.Question("Project name ["+proj.Name+"]: ", &proj.Name)
// 	UI.Question("Project path (absolute)["+proj.Path+"]: ", &proj.Path)
// 	UI.Question("Project description ["+proj.Description+"]:", &proj.Description)

// 	for k := range proj.Services {
// 		Edit(&proj.Services[k])
// 	}

// 	UI.Println(nil, proj)
// 	if !UI.YesNoQuestion(nil, "Looking good?") {
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
