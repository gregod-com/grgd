package gormdal

import (
	"gorm.io/gorm"
)

// Service ...
type Service struct {
	gorm.Model
	Name          string
	Active        bool
	Environment   string
	Skaffold      bool
	Path          string
	Initialized   bool
	GRGDProjectID uint
}

// Save ...
func (service *Service) Save(db *gorm.DB, i ...interface{}) error {
	db.Save(service)
	return nil
}

// Delete ...
func (service *Service) Delete(db *gorm.DB, i ...interface{}) error {
	db.Delete(service)
	return nil
}

// ServiceAlias ...
type ServiceAlias struct {
	gorm.Model
	ServiceID uint
}
