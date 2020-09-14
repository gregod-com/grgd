package wrappers

import (
	"grgd/persistence"
)

// ServiceWrapper ...
type ServiceWrapper struct {
	model *persistence.Service
}

// CreateServiceWrapper ...
func CreateServiceWrapper(mService *persistence.Service) *ServiceWrapper {
	return &ServiceWrapper{model: mService}
}

// GetName ...
func (service *ServiceWrapper) GetName() string {
	return service.model.Name
}

// GetPath ...
func (service *ServiceWrapper) GetPath(i ...interface{}) string {
	return service.model.Path
}

// SetPath ...
func (service *ServiceWrapper) SetPath(path string, i ...interface{}) error {
	service.model.Path = path
	return nil
}

// GetActive ...
func (service *ServiceWrapper) GetActive(i ...interface{}) bool {
	return service.model.Active
}

// SetActive ...
func (service *ServiceWrapper) SetActive(active bool, i ...interface{}) error {
	service.model.Active = active
	return nil
}

// ToggleActive ...
func (service *ServiceWrapper) ToggleActive(i ...interface{}) error {
	service.model.Active = !service.model.Active
	return nil
}

// GetValues ...
func (service *ServiceWrapper) GetValues(i ...interface{}) []string {
	return []string{service.model.Environment, service.model.Name, service.model.Path}
}

// // Edit ...
// func (service *Service) Edit(db *gorm.DB, i ...interface{}) error {
// 	UI.Println(nil, "Here in edit profile")
// 	return nil
// }

// // Init ...
// func (service *Service) Init(db *gorm.DB, i ...interface{}) error {
// 	if !UI.YesNoQuestion(nil, "Init service "+service.Name+" now?") {
// 		return nil
// 	}

// 	UI.Question("Service name ["+service.Name+"]: ", &service.Name)
// 	UI.Question("Service path (absolute)["+service.Path+"]: ", &service.Path)
// 	service.Initialized = true
// 	return nil
// }
