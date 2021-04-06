package gormdal

import (
	"log"
	"os"

	"github.com/gregod-com/grgd/interfaces"
)

func setupDatabase(fsm interfaces.IFileSystemManipulator, logger interfaces.ILogger) interfaces.IDAL {
	os.Remove(fsm.LoadBootConfig().DatabasePath)

	dal := ProvideDAL(fsm, logger)
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

func tearDownDatabase(fsm interfaces.IFileSystemManipulator) {
	err := os.Remove(fsm.LoadBootConfig().DatabasePath)
	if err != nil {
		log.Fatal(err)
	}
	return
}
