package helpers

import (
	"errors"
	"fmt"
	"log"
	"os"
	"reflect"
)

// HomeDir ...
func HomeDir() string {
	dir, errHomeDir := os.UserHomeDir()
	if errHomeDir != nil {
		log.Fatal(errHomeDir)
	}
	return dir
}

// ExtractMetadata ...
func ExtractMetadata(m map[string]interface{}, key string, container interface{}) error {
	if m[key] == nil {
		return fmt.Errorf("Could not find key `%v` in passed map ", key)
	}

	// fmt.Println(reflect.Indirect(reflect.ValueOf(container)).Type())
	// fmt.Println(reflect.TypeOf(m[key]))

	val := reflect.ValueOf(container)
	if val.Kind() != reflect.Ptr {
		return errors.New("Passed container is not a pointer")
	}
	containerType := reflect.Indirect(reflect.ValueOf(container)).Type()
	elementType := reflect.TypeOf(m[key])

	if !elementType.AssignableTo(containerType) {
		return fmt.Errorf("Value at key `%v` (type %v) in passed map is not assignable to pointer (type `%v`)", key, reflect.TypeOf(m[key]), reflect.Indirect(reflect.ValueOf(container)).Type())
	}

	val.Elem().Set(reflect.ValueOf(m[key]))
	return nil
}

// ExtractMetadataFatal calls ExtractMetadata but fails fataly before returning to caller if extraction has error
// this allows for less lines in calling code for essential extractions that would need to interrupt the application anyways
func ExtractMetadataFatal(m map[string]interface{}, key string, container interface{}) {
	if err := ExtractMetadata(m, key, container); err != nil {
		log.Fatal(err.Error())
	}
}
