package rage

type SchemaContent struct {
	fields []schemaField
}

// I can make this to return an object that contains a pointer to the field and another pointer
// to the schema. that object will contain the AddField method and the other methods for manipulating
// the field
func (s *SchemaContent) addField(fieldName string, dataType FieldType) *schemaField {
	for _, field := range s.fields {
		if field.name == fieldName {
			return &s.fields[len(s.fields)-1]
		}
	}
	field := &schemaField{
		name:     fieldName,
		dataType: dataType,
	}
	s.fields = append(s.fields, *field)
	return field
}

// func (s *SchemaContent) setName(name string) *schemaField {}
