package persistence

import "gorm.io/gorm"

// DataObject ...
type DataObject interface {
	Save(db *gorm.DB, i ...interface{}) error
	Delete(db *gorm.DB, i ...interface{}) error
}
