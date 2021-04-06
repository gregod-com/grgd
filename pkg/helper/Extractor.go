package helper

import (
	"errors"
	"fmt"
	"log"
	"reflect"

	"github.com/gregod-com/grgd/interfaces"

	"github.com/urfave/cli/v2"
)

// GetExtractor ...
func GetExtractor() interfaces.IExtractor {
	return &Extractor{}
}

// Extractor ...
type Extractor struct {
	// do not put fields here, since the extractor will never be initialized this way
}

// GetMetadata ...
func (h *Extractor) GetMetadata(m map[string]interface{}, key string, container interface{}) error {
	if m[key] == nil {
		return fmt.Errorf("Could not find key `%v` in passed map", key)
	}

	val := reflect.ValueOf(container)
	if val.Kind() != reflect.Ptr {
		return errors.New("Passed container is not a pointer")
	}
	containerType := reflect.Indirect(reflect.ValueOf(container)).Type()
	elementType := reflect.TypeOf(m[key])

	if !elementType.AssignableTo(containerType) {
		return fmt.Errorf("Value at key `%v` (type %v) in passed map is not assignable to pointer (type `%v`)",
			key, reflect.TypeOf(m[key]), reflect.Indirect(reflect.ValueOf(container)).Type())
	}

	val.Elem().Set(reflect.ValueOf(m[key]))
	return nil
}

// ExtractMetadataFatal calls ExtractMetadata but fails fataly before returning to caller if extraction has error
// this allows for less lines in calling code for essential extractions that would need to interrupt the application anyways
func (h *Extractor) GetMetadataFatal(m map[string]interface{}, key string, container interface{}) {
	if err := h.GetMetadata(m, key, container); err != nil {
		log.Fatal(err.Error())
	}
}

// ExtractProfile ...
func (h *Extractor) GetCore(i interface{}) interfaces.ICore {
	var core interfaces.ICore
	switch c := i.(type) {
	case *cli.Context:
		h.GetMetadataFatal(c.App.Metadata, "core", &core)
	default:
		log.Fatalf("Passed context is of type %T\n", c)
	}
	return core
}
