package validate

import (
	"fmt"
	"reflect"
)

type SchemaContent struct {
	fields []schemaField
}

// This will return a ValidationResult eventually. Just chill for now
func (s *SchemaContent) Validate(object interface{}) (bool, error) {
	for i, n := range s.fields {
		fmt.Println(i, n)
	}
	return true, nil
}

// I can make this to return an object that contains a pointer to the field and another pointer
// to the schema. that object will contain the AddField method and the other methods for manipulating
// the field
func (s *SchemaContent) addField(fieldName string, dataType reflect.Type) *schemaField {
	field := &schemaField{
		name:     fieldName,
		dataType: dataType.Name(),
	}
	s.fields = append(s.fields, *field)
	return field
}
