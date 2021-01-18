package gormdal

import (
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
	dal.datbasePath = dbPath
	if dbPath == "" {
		dbFolder := fsmanipulator.HomeDir(".grgd", "gorm")
		fsmanipulator.CheckOrCreateFolder(dbFolder, os.FileMode(uint32(0760)))
		dal.datbasePath = fsmanipulator.HomeDir(".grgd", "gorm", "data.db")
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
	datbasePath string
	db          *gorm.DB
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
	if dal.datbasePath == "" {
		dal.datbasePath = "./database.db"
	}
	db, err := gorm.Open(sqlite.Open(dal.datbasePath+"?cache=shared&mode=memory"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent)})
	if err != nil {
		panic("failed to connect database")
	}
	dal.db = db
}
