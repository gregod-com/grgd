package service

import I "github.com/gregod-com/grgd/interfaces"

// ProvideService ...
func ProvideService() I.IService {
	return &Service{}
}

// Service ...
type Service struct {
	id          uint
	name        string
	path        string
	active      bool
	environment string
	executer    interface{}
	initialized bool
}

// Init ...
func (service *Service) Init(
	id uint,
	name string,
	path string,
	active bool,
	environment string,
	executor interface{}) error {
	service.id = id
	service.name = name
	service.path = path
	service.active = active
	service.environment = environment
	service.executer = executor
	service.initialized = true
	return nil
}

// GetName ...
func (service *Service) GetName() string {
	return service.name
}

// GetPath ...
func (service *Service) GetPath(i ...interface{}) string {
	return service.path
}

// SetPath ...
func (service *Service) SetPath(path string, i ...interface{}) error {
	service.path = path
	return nil
}

// GetActive ...
func (service *Service) GetActive(i ...interface{}) bool {
	return service.active
}

// SetActive ...
func (service *Service) SetActive(active bool, i ...interface{}) error {
	service.active = active
	return nil
}

// ToggleActive ...
func (service *Service) ToggleActive(i ...interface{}) error {
	service.active = !service.active
	return nil
}

// GetValues ...
func (service *Service) GetValues(i ...interface{}) []string {
	return []string{service.environment, service.name, service.path}
}

// // Edit ...
// func (service *Service) Edit(db *gorm.DB, i ...interface{}) error {
// 	UI.Println("Here in edit profile")
// 	return nil
// }

// // Init ...
// func (service *Service) Init(db *gorm.DB, i ...interface{}) error {
// 	if !UI.YesNoQuestion("Init service "+service.Name+" now?") {
// 		return nil
// 	}

// 	UI.Question("Service name ["+service.Name+"]: ", &service.Name)
// 	UI.Question("Service path (absolute)["+service.Path+"]: ", &service.Path)
// 	service.Initialized = true
// 	return nil
// }
