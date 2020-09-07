package persistence

import (
	"gorm.io/gorm"
)

// Service ...
type Service struct {
	gorm.Model    `json:"-"`
	Name          string
	Active        bool
	Environment   string
	Path          string
	GRGDProjectID uint
}

// Save ...
func (service *Service) Save(db *gorm.DB) error {
	db.Save(service)
	return nil
}

// ServiceAlias ...
type ServiceAlias struct {
	gorm.Model
	ServiceID uint
}
