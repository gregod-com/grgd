package gormdal

import (
	"fmt"
	"os"

	"github.com/gregod-com/grgd/interfaces"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// ProvideDAL ...
func ProvideDAL(helper interfaces.IHelper, logger interfaces.ILogger) interfaces.IDAL {
	dal := new(GormDAL)
	dal.logger = logger
	dal.logger.Tracef("provide %T", dal)
	dal.databasePath = helper.LoadBootConfig().DatabasePath
	helper.CheckOrCreateParentFolder(dal.databasePath, os.FileMode(uint32(0760)))
	dal.connect()
	dal.db.AutoMigrate(&ProfileModel{})
	dal.db.AutoMigrate(&ProjectModel{})
	dal.db.AutoMigrate(&ServiceModel{})
	return dal
}

// GormDAL ...
type GormDAL struct {
	databasePath string
	db           *gorm.DB
	logger       interfaces.ILogger
}

type dto interface {
	GetName() string
	GetID() uint
}

// Create ...
func (dal *GormDAL) Create(i interface{}) error {
	dal.logger.Tracef("%T", i)
	dto, err := dal.toDTO(i)
	if err != nil {
		return err
	}
	return dal.db.Create(dto).Error
}

// Read ...
func (dal *GormDAL) Read(i interface{}) error {
	dal.logger.Tracef("%v", i)
	dto, err := dal.toDTO(i)
	if err != nil {
		return err
	}
	err = dal.db.First(dto, "name = ?", dto.GetName()).Error
	if err != nil {
		return err
	}
	return dal.toInterface(dto, i)
}

// Update ...
func (dal *GormDAL) Update(i interface{}) error {
	dal.logger.Tracef("%T", i)
	dto, err := dal.toDTO(i)
	if err != nil {
		return err
	}
	_, ok := i.(interfaces.IProfile)
	if !ok {
		dal.logger.Fatal("hmmm konisch")
	}
	// dal.logger.Fatalf("%v ", p.GetID())

	return dal.db.Save(dto).Error
}

// Delete ...
func (dal *GormDAL) Delete(i interface{}) error {
	dal.logger.Tracef("%T", i)
	dto, err := dal.toDTO(i)
	if err != nil {
		return err
	}
	return dal.db.Delete(dto).Error
}

// ReadAll ...
func (dal *GormDAL) ReadAll(dataType interface{}) (map[string]interface{}, error) {
	dal.logger.Tracef("%T", dataType)
	dto, err := dal.toDTO(dataType)
	if err != nil {
		return nil, err
	}
	result := dal.db.Find(dto)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, fmt.Errorf("No entries for %T", dataType)
	}
	rows, err := result.Rows()
	defer rows.Close()
	if err != nil {
		return nil, err
	}

	retmap := make(map[string]interface{})

	for rows.Next() {
		tempDTO, _ := dal.toDTO(dataType)
		// ScanRows is a method of `gorm.DB`, it can be used to scan a row into a struct
		dal.db.ScanRows(rows, tempDTO)
		dal.toInterface(tempDTO, dataType)
		retmap[tempDTO.GetName()] = dataType
	}
	return retmap, result.Error
}

func (dal *GormDAL) connect() {
	dal.logger.Tracef("Connecting to %v", dal.databasePath)
	db, err := gorm.Open(sqlite.Open(dal.databasePath+"?cache=shared&mode=memory"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent)})
	if err != nil {
		dal.logger.Fatal("failed to connect database")
	}
	dal.db = db
}

func (dal *GormDAL) toDTO(i interface{}) (dto, error) {
	switch v := i.(type) {
	case interfaces.IProfile:
		pm := &ProfileModel{}
		dal.logger.Tracef("converting %T to %T", v, pm)
		err := profileToIProfileModel(v, pm)
		return pm, err
	case *ProfileModel:
		dal.logger.Tracef("No Conversion needed (%T is aleady dto)", v)
		return v, nil
	case *ProjectModel:
		dal.logger.Tracef("No Conversion needed (%T is aleady dto)", v)
		return v, nil
	case *ServiceModel:
		dal.logger.Tracef("No Conversion needed (%T is aleady dto)", v)
		return v, nil
	default:
		dal.logger.Fatalf("IIIIIIIS %T!!!!!!", v)
	}
	return nil, nil
}

func (dal *GormDAL) toInterface(dto interface{}, i interface{}) error {
	switch v := dto.(type) {
	case *ProfileModel:
		p, ok := i.(interfaces.IProfile)
		if !ok {
			return fmt.Errorf("missmatch when trying to convert %T to %T", dto, i)
		}
		dal.logger.Tracef("converting %T to %T", v, p)
		return profileModelToIProfile(v, p)
	case ProjectModel:
		_, ok := i.(interfaces.IProject)
		if !ok {
			return fmt.Errorf("missmatch when trying to convert %T to %T", dto, i)
		}
	case ServiceModel:
		_, ok := i.(interfaces.IProfile)
		if !ok {
			return fmt.Errorf("missmatch when trying to convert %T to %T", dto, i)
		}

	default:
		dal.logger.Tracef("FOUND dto %T", v)
	}
	dal.logger.Tracef("converted %T to %T", dto, i)
	return nil
}
