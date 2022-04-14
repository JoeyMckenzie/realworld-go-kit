package utilities

import (
	"fmt"
	"reflect"
)

type postgresTypeMap map[reflect.Kind]string

var postgresTypeMapValues = postgresTypeMap{
	reflect.String: "varchar",
	reflect.Int:    "integer",
}

func GetValueFromType(value interface{}) string {
	reflectedKind := reflect.TypeOf(value).Kind()

	if typeMapValue, exists := postgresTypeMapValues[reflectedKind]; exists {
		return typeMapValue
	}

	return ""
}

func GetParameterizedValue(value interface{}, parameterizedInput []interface{}) string {
	reflectedKind := reflect.TypeOf(value).Kind()

	if typeMapValue, exists := postgresTypeMapValues[reflectedKind]; exists {
		return fmt.Sprintf("$%v::%s", len(parameterizedInput), typeMapValue)
	}

	return ""
}
