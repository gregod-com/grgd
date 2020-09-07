package persistence

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/TylerBrock/colorjson"
	"github.com/gregod-com/grgd/helpers"
	"gorm.io/gorm"
)

// Profile ...
type Profile struct {
	gorm.Model  `json:"-"`
	Name        string
	HomeDir     string
	PluginDir   string
	Projects    []GRGDProject
	Initialized bool
}

// String  ...
func (profile Profile) String() string {
	var obj map[string]interface{}
	// create json string from object
	str, err := json.MarshalIndent(profile, "", "  ")
	helpers.CheckErr(err)

	// create simplified object from json string
	json.Unmarshal([]byte(str), &obj)

	f := colorjson.NewFormatter()
	f.Indent = 4

	// create colored json string from simplified object
	data, err := f.Marshal(obj)
	helpers.CheckErr(err)

	return string(data)
}

// Save ...
func (profile *Profile) Save(db *gorm.DB) error {
	db.Save(profile)
	for k := range profile.Projects {
		profile.Projects[k].Save(db)
	}
	return nil
}

func (profile *Profile) initProfile() {

	log.Println("das")
	// fmt.ClearScreen(nil)
	var answer string

	answer = profile.HomeDir
	fmt.Printf("Hey %v, let's init your profile \n\n", profile.Name)
	fmt.Printf("First things first:\n\n")

	fmt.Printf("Base profile directory [`Enter` for default: %v]: ", profile.HomeDir)
	fmt.Scanln(&answer)

	for !helpers.PathExists(answer) {
		fmt.Printf("The path %v does not exists. Try again or use default [%v]: ", answer, profile.HomeDir)
		answer = profile.HomeDir
		fmt.Scanln(&answer)
	}
	profile.HomeDir = answer

	answer = profile.PluginDir
	fmt.Printf("Base plugin directory [`Enter` for default: %v]: ", profile.PluginDir)
	fmt.Scanln(&answer)

	for !helpers.PathExists(answer) {
		fmt.Printf("The path %v does not exists. Try again or use default [%v]: ", answer, profile.PluginDir)
		answer = profile.PluginDir
		fmt.Scanln(&answer)
	}

	// // TODO: move all those configs into the config file
	// c.App.Metadata["pluginIndex"] = homedir + "/.grgd/plugins/index.yaml"
	// c.App.Metadata["remoteIndex"] = "https://s3.gregod.com/public/plugins/index.yaml"
	// c.App.Metadata["AWS-REGION"] = "eu-central-1"
	// c.App.Metadata["updatecheckinterval"] = time.Millisecond * 50

	profile.Initialized = true

	fmt.Println(profile)
	if !helpers.FallbackUI.YesNoQuestion(helpers.FallbackUI{}, nil, "Looking good? [y/n]") {
		profile.initProfile()
	}
}
