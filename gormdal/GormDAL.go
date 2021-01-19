package gormdal

import (
	"log"
	"os"

	"github.com/gregod-com/grgd/interfaces"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/logger"
)

// ProvideDAL ...
func ProvideDAL(dbPath string, fsmanipulator interfaces.IFileSystemManipulator) interfaces.IDAL {
	dal := new(GormDAL)
	dal.databasePath = dbPath

	dbFolder := fsmanipulator.HomeDir(".grgd", "gorm")
	fsmanipulator.CheckOrCreateFolder(dbFolder, os.FileMode(uint32(0760)))

	if dbPath == "" {
		dal.databasePath = fsmanipulator.HomeDir(".grgd", "gorm", "data.db")
	}

	dal.connect()
	dal.db.AutoMigrate(&ProfileModel{})
	dal.db.AutoMigrate(&ProjectModel{})
	dal.db.AutoMigrate(&ServiceModel{})
	return dal
}

// ProvideDefaultDBPath ...
func ProvideDefaultDBPath(fsmanipulator interfaces.IFileSystemManipulator) string {
	return fsmanipulator.HomeDir(".grgd", "gorm", "data.db")
}

// ProvideTESTDBPath ...
func ProvideTESTDBPath(fsmanipulator interfaces.IFileSystemManipulator) string {
	return fsmanipulator.HomeDir(".grgd", "gorm", "testdata.db")
}

// GormDAL ...
type GormDAL struct {
	databasePath string
	db           *gorm.DB
}

// Create ...
func (dal *GormDAL) Create(i interface{}) error {
	return dal.db.Create(i).Error
}

// Read ...
func (dal *GormDAL) Read(i interface{}) error {
	return dal.db.First(i).Error
}

// Update ...
func (dal *GormDAL) Update(i interface{}) error {
	return dal.db.Save(i).Error
}

// Delete ...
func (dal *GormDAL) Delete(i interface{}) error {
	return dal.db.Delete(i).Error
}

func (dal *GormDAL) GetProfile() (interfaces.IProfileModel, error) {
	profile := &ProfileModel{}
	err := dal.Read(profile)
	return profile, err
}

// GetAll ...
func (dal *GormDAL) GetAll(array []interface{}) error {
	result := dal.db.Find(array)
	dal.db.Preload(clause.Associations).Find(array)
	return result.Error
}

func (dal *GormDAL) connect() {
	log.Printf("Connecting to %v", dal.databasePath)
	db, err := gorm.Open(sqlite.Open(dal.databasePath+"?cache=shared&mode=memory"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent)})
	if err != nil {
		log.Fatal("failed to connect database")
	}
	dal.db = db
}
