package gormdal

import (
	"log"
	"os"

	"github.com/gregod-com/grgd/interfaces"
)

var databaseFileName string

func setupDatabase(fsm interfaces.IFileSystemManipulator) interfaces.IDAL {
	var err error
	os.Remove(databaseFileName)

	databaseFileName, err = os.UserHomeDir()
	if err != nil {
		log.Fatal("Error getting HOMEDIR")
	}
	databaseFileName += "/.grgd/gorm/" + "test-database.db"
	dal := ProvideDAL(databaseFileName, fsm)
	dal.Create(&ProfileModel{
		Name:    "TESTProfile",
		HomeDir: "./test-me-dir/"})
	dal.Create(&ProjectModel{
		Name: "TESTProject",
		Path: "./test-me-dir/"})
	dal.Create(&ServiceModel{
		Name: "TESTService",
		Path: "./test-me-dir/"})

	return dal
}

func tearDownDatabase(dal interfaces.IDAL) {
	// dal.Delete(&ProfileModel{Name: "TESTProfile"})
	err := os.Remove(databaseFileName)
	if err != nil {
		log.Fatal(err)
	}
	return
}
