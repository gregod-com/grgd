package persistence

import (
	"github.com/gregod-com/grgd/helpers"
	"github.com/gregod-com/grgdplugincontracts"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// DataObject ...
type DataObject interface {
	Save(db *gorm.DB) error
}

var datbasePath string
var ui grgdplugincontracts.IUIPlugin

// InitDatabase ...
func InitDatabase(dbPath string, uiIn grgdplugincontracts.IUIPlugin) {
	datbasePath = dbPath
	ui = uiIn
	db := connect()

	// Migrate the schema
	db.AutoMigrate(&Profile{})
	db.AutoMigrate(&GRGDProject{})
	db.AutoMigrate(&Service{})

}

// GetProfile ...
func GetProfile(profilename string) *Profile {
	var profile Profile

	// defaults for new profile
	defaultHomedir := helpers.HomeDir()
	defaultPlugindir := defaultHomedir + "/.grgd/plugins"

	db := connect()

	db.FirstOrCreate(&profile,
		Profile{
			Name: profilename,
		})

	if !profile.Initialized {
		profile.HomeDir = defaultHomedir
		profile.PluginDir = defaultPlugindir
		profile.initProfile()
		db.Save(&profile)
	}
	return &profile
}

// GetProfileEager ...
func GetProfileEager(profilename string) *Profile {
	db := connect()
	profile := GetProfile(profilename)

	db.Preload(clause.Associations).Find(profile)
	db.Preload(clause.Associations).Find(&profile.Projects)

	return profile
}

// Save ...
func Save(obj DataObject) error {
	db := connect()
	obj.Save(db)
	return nil
}

func connect() *gorm.DB {
	db, err := gorm.Open(sqlite.Open(datbasePath+"?cache=shared&mode=memory"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return db
}
