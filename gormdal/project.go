package gormdal

import (
	log "github.com/sirupsen/logrus"

	"gorm.io/gorm"
)

// GRGDProject ...
type GRGDProject struct {
	gorm.Model
	Name        string
	Path        string
	ProfileID   uint
	Initialized bool
	Services    []*Service `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Description string
}

// Save ...
func (proj *GRGDProject) Save(db *gorm.DB, i ...interface{}) error {
	db.Save(proj)
	for k := range proj.Services {
		proj.Services[k].Save(db)
	}
	return nil
}

// Delete ...
func (proj *GRGDProject) Delete(db *gorm.DB, i ...interface{}) error {
	db.Delete(proj)
	for k := range proj.Services {
		proj.Services[k].Delete(db)
	}
	return nil
}

// BeforeUpdate ...
func (proj *GRGDProject) BeforeUpdate(tx *gorm.DB) (err error) {
	log.Trace("Updating proj")
	return
}

// BeforeDelete ...
func (proj *GRGDProject) BeforeDelete(tx *gorm.DB) (err error) {
	log.Trace("Deleteoig proj")
	return
}

// ProjectTag ...
type ProjectTag struct {
	gorm.Model
	Name string
}
