package helpers

import (
	"fmt"
	"reflect"

	"go_sales_api/internal/model"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var structMap = map[string]reflect.Type{
	"Cashiers":   reflect.TypeOf(model.Cashiers{}),
	"Categories": reflect.TypeOf(model.Categories{}),
	// Add more struct types as needed
}

func GetStructMap(tablenames string) (reflect.Type, bool) {
	val, ok := structMap[tablenames]
	return val, ok
}

func CreateStructInstance(structName string) (interface{}, error) {
	// Get the type of the struct based on its name
	structType, err := StructNameToType(structName)
	if err != nil {
		return nil, err
	}

	// Create a new instance of the struct
	instance := reflect.New(structType).Elem().Interface()
	return instance, nil
}

func StructNameToType(structName string) (reflect.Type, error) {
	structName = cases.Title(language.English, cases.NoLower).String(structName)

	// Get the type from structMap
	structType, ok := structMap[structName]
	if !ok {
		return nil, fmt.Errorf("struct type %s not found", structName)
	}
	return structType, nil
}
