package validate

import (
	"fmt"
	"reflect"
)

func getValueOfFieldWithName(p interface{}, fieldName string) (interface{}, error) {
	reflectedObject := reflect.ValueOf(p)
	value := reflect.Indirect(reflectedObject).FieldByName(fieldName)

	if !value.IsValid() {
		return nil, fmt.Errorf("the field '%s' was not found", fieldName)
	}

	return value.Interface(), nil
}

func isNumericType(obj interface{}) bool {
	return (reflect.TypeOf(obj).Name() == "int" || reflect.TypeOf(obj).Name() == "uint" || reflect.TypeOf(obj).Name() == "int8" || reflect.TypeOf(obj).Name() == "uint8" || reflect.TypeOf(obj).Name() == "int16" || reflect.TypeOf(obj).Name() == "uint16" || reflect.TypeOf(obj).Name() == "int32" || reflect.TypeOf(obj).Name() == "uint32" || reflect.TypeOf(obj).Name() == "int64" || reflect.TypeOf(obj).Name() == "uint64" || reflect.TypeOf(obj).Name() == "float32" || reflect.TypeOf(obj).Name() == "float64" || reflect.TypeOf(obj).Name() == "complex64" || reflect.TypeOf(obj).Name() == "complex128")
}

func fieldNameIsNumericType(fieldName string) bool {
	return (fieldName == "int" || fieldName == "uint" || fieldName == "int8" || fieldName == "uint8" || fieldName == "int16" || fieldName == "uint16" || fieldName == "int32" ||
		fieldName == "uint32" || fieldName == "int64" || fieldName == "uint64" || fieldName == "float32" || fieldName == "float64" || fieldName == "complex64" || fieldName == "complex128")
}
