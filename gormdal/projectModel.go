package gormdal

import (
	"gorm.io/gorm"
)

// ProjectModel ...
type ProjectModel struct {
	gorm.Model
	Name           string
	Path           string
	ProfileModelID uint
	Initialized    bool
	Services       []*ServiceModel `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Description    string
}

// ProjectTag ...
type ProjectTag struct {
	gorm.Model
	Name string
}
