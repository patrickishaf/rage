package validate

import (
	"fmt"
	"reflect"
)

type Schema struct {
	content *SchemaContent
}

type schemaField struct {
	name       string
	dataType   string
	isRequired bool
	min        int
	max        int
}

func NewSchema() *Schema {
	schemaContent := &SchemaContent{
		fields: []schemaField{},
	}
	return &Schema{
		content: schemaContent,
	}
}

/*
TODO: Implement this function to return a vaidation result depending on the status of the object
*/
func (s *Schema) Validate(obj interface{}) ValidationResult {
	for _, field := range s.content.fields {
		_, found := reflect.TypeOf(obj).FieldByName(field.name)
		fmt.Println(found)
	}
	return ValidationResult{}
}

func (s *Schema) AddField(fieldName string, dataType reflect.Type) *Schema {
	s.content.addField(fieldName, dataType)
	return s
}

func (s *Schema) SetName(name string) *Schema {
	if len(s.content.fields) == 0 {
		s.content.fields = append(s.content.fields, schemaField{})
	}
	s.content.fields[len(s.content.fields)-1].name = name
	return s
}

func (s *Schema) SetDataType(dataType reflect.Type) *Schema {
	if len(s.content.fields) == 0 {
		s.content.fields = append(s.content.fields, schemaField{})
	}
	s.content.fields[len(s.content.fields)-1].dataType = dataType.Name()
	return s
}

func (s *Schema) SetIsRequired(isRequired bool) *Schema {
	if len(s.content.fields) == 0 {
		s.content.fields = append(s.content.fields, schemaField{})
	}
	s.content.fields[len(s.content.fields)-1].isRequired = isRequired
	return s
}

func (s *Schema) SetMin(value int) *Schema {
	if len(s.content.fields) == 0 {
		s.content.fields = append(s.content.fields, schemaField{})
	}
	s.content.fields[len(s.content.fields)-1].min = value
	return s
}

func (s *Schema) SetMax(value int) *Schema {
	if len(s.content.fields) == 0 {
		s.content.fields = append(s.content.fields, schemaField{})
	}
	s.content.fields[len(s.content.fields)-1].max = value
	return s
}
