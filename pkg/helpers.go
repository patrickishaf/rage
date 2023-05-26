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
	return (FieldType(reflect.TypeOf(obj).String()) == Int ||
		FieldType(reflect.TypeOf(obj).String()) == Uint ||
		FieldType(reflect.TypeOf(obj).String()) == Int8 ||
		FieldType(reflect.TypeOf(obj).String()) == Uint8 ||
		FieldType(reflect.TypeOf(obj).String()) == Int16 ||
		FieldType(reflect.TypeOf(obj).String()) == Uint16 ||
		FieldType(reflect.TypeOf(obj).String()) == Int32 ||
		FieldType(reflect.TypeOf(obj).String()) == Uint32 ||
		FieldType(reflect.TypeOf(obj).String()) == Int64 ||
		FieldType(reflect.TypeOf(obj).String()) == Uint64 ||
		FieldType(reflect.TypeOf(obj).String()) == Float32 ||
		FieldType(reflect.TypeOf(obj).String()) == Float64)
}

func fieldNameIsNumericType(fieldName FieldType) bool {
	return (fieldName == Int || fieldName == Uint || fieldName == Int8 ||
		fieldName == Uint8 || fieldName == Int16 || fieldName == Uint16 ||
		fieldName == Int32 || fieldName == Uint32 || fieldName == Int64 ||
		fieldName == Uint64 || fieldName == Float32 || fieldName == Float64)
}
