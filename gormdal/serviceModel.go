package gormdal

import (
	"gorm.io/gorm"
)

// ServiceModel ...
type ServiceModel struct {
	gorm.Model
	Name           string
	Active         bool
	Environment    string
	Skaffold       bool
	Path           string
	Initialized    bool
	ProjectModelID uint
}
