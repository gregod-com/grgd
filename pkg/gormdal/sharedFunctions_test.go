package gormdal

import (
	"log"
	"os"

	"github.com/gregod-com/grgd/interfaces"
)

func setupDatabase(helper interfaces.IHelper, logger interfaces.ILogger) interfaces.IDAL {
	os.Remove(helper.LoadBootConfig().DatabasePath)

	dal := ProvideDAL(helper, logger)
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

func tearDownDatabase(helper interfaces.IHelper) {
	err := os.Remove(helper.LoadBootConfig().DatabasePath)
	if err != nil {
		log.Fatal(err)
	}
	return
}
