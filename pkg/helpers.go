package rage

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
	return (reflect.TypeOf(obj).Name() == Int || reflect.TypeOf(obj).Name() == Uint || reflect.TypeOf(obj).Name() == Int8 || reflect.TypeOf(obj).Name() == Uint8 || reflect.TypeOf(obj).Name() == Int16 || reflect.TypeOf(obj).Name() == Uint16 || reflect.TypeOf(obj).Name() == Int32 || reflect.TypeOf(obj).Name() == Uint32 || reflect.TypeOf(obj).Name() == Int64 || reflect.TypeOf(obj).Name() == Uint64 || reflect.TypeOf(obj).Name() == Float32 || reflect.TypeOf(obj).Name() == Float64)
}

func fieldNameIsNumericType(fieldName string) bool {
	return (fieldName == Int || fieldName == Uint || fieldName == Int8 || fieldName == Uint8 || fieldName == Int16 || fieldName == Uint16 || fieldName == Int32 ||
		fieldName == Uint32 || fieldName == Int64 || fieldName == Uint64 || fieldName == Float32 || fieldName == Float64)
}
