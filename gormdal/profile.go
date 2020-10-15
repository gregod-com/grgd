package gormdal

import (
	"gorm.io/gorm"
)

// Profile ...
type Profile struct {
	gorm.Model
	Name             string
	HomeDir          string
	PluginDir        string
	Projects         []*GRGDProject `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CurrentProjectID uint
	Initialized      bool
}

// Save ...
func (profile *Profile) Save(db *gorm.DB, i ...interface{}) error {
	db.Save(profile)
	for k := range profile.Projects {
		profile.Projects[k].Save(db)
	}
	return nil
}

// Delete ...
func (profile *Profile) Delete(db *gorm.DB, i ...interface{}) error {
	db.Delete(profile)
	for k := range profile.Projects {
		profile.Projects[k].Delete(db)
	}
	return nil
}
