package gormdal

import (
	"fmt"
	"os"
	"reflect"

	"github.com/gregod-com/grgd/interfaces"
	"github.com/gregod-com/grgd/pkg/profile"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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
	return dal.db.Session(&gorm.Session{FullSaveAssociations: true}).Save(dto).Error
}

// Delete ...
func (dal *GormDAL) Delete(i interface{}) error {
	dal.logger.Tracef("%T", i)
	dto, err := dal.toDTO(i)
	if err != nil {
		return err
	}
	// no soft delete, but relly remove from db
	return dal.db.Unscoped().Delete(dto).Error
}

// ReadAll ...
func (dal *GormDAL) ReadAll(dataType interface{}) (map[string]interface{}, error) {
	dal.logger.Tracef("%T", dataType)
	var result *gorm.DB
	retmap := make(map[string]interface{})

	switch typ := reflect.ValueOf(dataType).Interface().(type) {
	case []interfaces.IProfile:
		profileModels := []ProfileModel{}
		result = dal.db.Preload(clause.Associations).Find(&profileModels)
		if result.Error != nil {
			return nil, result.Error
		}
		for k := range profileModels {
			p := profile.ProvideProfile(dal.logger, nil)
			dal.toInterface(&profileModels[k], p)
			retmap[p.GetName()] = p
		}
	default:
		dal.logger.Fatalf("result: %v (%T)", typ, typ)

	}
	return retmap, result.Error
}

func (dal *GormDAL) connect() {
	var lvl logger.LogLevel
	switch dal.logger.GetLevel() {
	case "trace":
		lvl = logger.Info
	case "debug":
		lvl = logger.Info
	case "info":
		lvl = logger.Warn
	case "warn":
		lvl = logger.Warn
	case "error":
		lvl = logger.Error
	default:
		lvl = logger.Silent
	}

	dal.logger.Tracef("Connecting to %v", dal.databasePath)
	db, err := gorm.Open(sqlite.Open(dal.databasePath+"?cache=shared&mode=memory"), &gorm.Config{
		Logger: logger.Default.LogMode(lvl)})
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
		err := iprofileToProfileModel(v, pm)
		return pm, err
	case interfaces.IProject:
		pm := &ProjectModel{}
		dal.logger.Tracef("converting %T to %T", v, pm)
		err := iprojectToProjectModel(v, pm)
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
		dal.logger.Debug("here should be some projects loaded %v", v.Projects)
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
		dal.logger.Warnf("FOUND dto %T", v)
	}
	return nil
}
