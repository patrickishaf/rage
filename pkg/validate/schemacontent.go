package validate

import (
	"reflect"
)

type SchemaContent struct {
	fields []schemaField
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
