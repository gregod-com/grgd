package gormdal

import (
	"grgd/interfaces"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// ProvideDAL ...
func ProvideDAL(fsmanipulator interfaces.IFileSystemManipulator) interfaces.IDAL {
	dbName := "data.db"
	dbFolder := fsmanipulator.HomeDir(".grgd", "gorm")
	dbPath := fsmanipulator.HomeDir(".grgd", "gorm", dbName)
	fsmanipulator.CheckOrCreateFolder(dbFolder, os.FileMode(uint32(0760)))

	dal := new(GormDAL)
	dal.datbasePath = dbPath

	dal.connect()
	// dal.db.AutoMigrate(&Profile{})
	// dal.db.AutoMigrate(&GRGDProject{})
	// dal.db.AutoMigrate(&Service{})
	return dal
}

// GormDAL ...
type GormDAL struct {
	datbasePath string
	db          *gorm.DB
}

// Create ...
func (dal *GormDAL) Create(i ...interface{}) error {
	dal.db.Save(i)
	return nil
}

// Read ...
func (dal *GormDAL) Read(i ...interface{}) (interface{}, error) {
	dal.db.First(i)
	return i, nil
}

// Update ...
func (dal *GormDAL) Update(i ...interface{}) error {
	dal.db.Save(i)
	return nil
}

// Delete ...
func (dal *GormDAL) Delete(i ...interface{}) error {
	dal.db.Delete(i)
	return nil
}

// GetAll ...
func (dal *GormDAL) GetAll(array []interface{}) error {
	result := dal.db.Find(array)
	dal.db.Preload(clause.Associations).Find(array)
	return result.Error
}

// Get ...
func (dal *GormDAL) Get(obj DataObject) error {
	// db := connect()
	// db.First(obj, obj)
	// db.Preload(clause.Associations).Find(obj)
	return nil
}

// GetOrCreate ...
func (dal *GormDAL) GetOrCreate(i interface{}) error {
	// db := connect()
	// db.FirstOrCreate(obj, obj)

	// db.Preload(clause.Associations).Find(obj)
	return nil
}

// Save ...
func (dal *GormDAL) Save(i ...interface{}) error {
	// db := connect()
	// obj.Save(db)
	return nil
}

// Remove ...
func (dal *GormDAL) Remove(i ...interface{}) error {
	// db := connect()
	// obj.Delete(db)
	return nil
}

func (dal *GormDAL) connect() {
	db, err := gorm.Open(sqlite.Open(dal.datbasePath+"?cache=shared&mode=memory"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	dal.db = db
}
