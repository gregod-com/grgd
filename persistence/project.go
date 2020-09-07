package persistence

import (
	"encoding/json"
	"fmt"

	"github.com/TylerBrock/colorjson"
	"github.com/gregod-com/grgd/helpers"
	"gorm.io/gorm"
)

// GRGDProject ...
type GRGDProject struct {
	gorm.Model  `json:"-"`
	Name        string
	Path        string
	ProfileID   uint
	Initialized bool
	Services    []Service
	Description string
}

// String  ...
func (proj GRGDProject) String() string {
	var obj map[string]interface{}
	// create json string from object
	str, err := json.MarshalIndent(proj, "", "  ")
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
func (proj *GRGDProject) Save(db *gorm.DB) error {
	db.Save(proj)
	for k := range proj.Services {
		proj.Services[k].Save(db)
	}
	return nil
}

// BeforeUpdate ...
func (proj *GRGDProject) BeforeUpdate(tx *gorm.DB) (err error) {
	fmt.Println("Updating proj")
	return
}

// ProjectTag ...
type ProjectTag struct {
	gorm.Model
	Name string
}
