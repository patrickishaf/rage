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
TODO: Implement this function to return a validation result depending on the status of the object
*/
func (s *Schema) Validate(obj interface{}) ValidationResult {
	fmt.Println("the validate function is running")
	result := ValidationResult{}
	isValid := true
	for _, field := range s.content.fields {
		/*
			TODO: If the field is not required, you still need to validate it to ensure that it has all the right properties
		*/
		if !field.isRequired {
			continue
		}
		/*
			If found == true, it means that the object has a field/property with the given name
		*/
		_, found := reflect.TypeOf(obj).FieldByName(field.name)
		if !found {
			isValid = false
			result.Reasons = append(result.Reasons, ValidationReason{
				Value:   field.name,
				Message: "the field '" + field.name + "' was not found in the object",
			})
			continue
		}
		// add a result reason if its value is below the min value
		// dataTypeOfSchema := s.content.fields[len(s.content.fields)-1].dataType
		/*
			You have the name of the field you are looking for. Check the value of the field of the object
			that has that name

			If the value of the error is not nil, the object you are validating has no field with the current field name
			If the value of [value] is not nil, the field exists. This goes without saying.
		*/
		value, err := getValueOfFieldWithName(obj, s.content.fields[len(s.content.fields)-1].name)
		if err != nil {
			isValid = false
			result.Reasons = append(result.Reasons, ValidationReason{
				Value:   field.name,
				Message: "the field '" + field.name + "' was not found in the object",
			})
			continue
		}
		// fmt.Println("dataTypeOfSchema => " + dataTypeOfSchema + "\nvalue of field with name " + field.name + " => ")
		fmt.Println(value)
		fmt.Println("")
		// If thid doesn't work I might have to convert all numeric types to float64 and store them like that
		if fieldNameIsNumericType(field.name) {
			fmt.Println("the value is a numeric type")
			// if floatValue, err := value.(float64); !err {
			// 	fmt.Println("the value is not a float type")
			// 	result.Reasons = append(result.Reasons, ValidationReason{
			// 		Value:   field.name,
			// 		Message: "the field ",
			// 	})
			// }
		} else {
			fmt.Println("the value is not a numeric type")
		}
	}
	result.IsValid = isValid
	return result
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

func (s *Schema) SetMinValue(value int) *Schema {
	if len(s.content.fields) == 0 {
		s.content.fields = append(s.content.fields, schemaField{})
	}
	currentField := &(s.content.fields[len(s.content.fields)-1])
	if fieldNameIsNumericType((*currentField).name) {
		(*currentField).min = value
	}
	return s
}

func (s *Schema) SetMaxValue(value int) *Schema {
	if len(s.content.fields) == 0 {
		s.content.fields = append(s.content.fields, schemaField{})
	}
	currentField := &(s.content.fields[len(s.content.fields)-1])
	if fieldNameIsNumericType((*currentField).name) {
		(*currentField).max = value
	}
	return s
}

func (s *Schema) SetMinLength(value int) *Schema {
	if len(s.content.fields) == 0 {
		s.content.fields = append(s.content.fields, schemaField{})
	}

	/*
		This method will not work if the field you want to give a minimum length is a numeric
		field. this is because numeric fields don't have length properties
	*/
	if !fieldNameIsNumericType(s.content.fields[len(s.content.fields)-1].dataType) {
		s.content.fields[len(s.content.fields)-1].min = value
	}
	return s
}

func (s *Schema) SetMaxLength(value int) *Schema {
	if len(s.content.fields) == 0 {
		s.content.fields = append(s.content.fields, schemaField{})
	}

	/*
		This method will not work if the field you want to give a minimum length is a numeric
		field. this is because numeric fields don't have length properties
	*/
	if !fieldNameIsNumericType(s.content.fields[len(s.content.fields)-1].dataType) {
		s.content.fields[len(s.content.fields)-1].max = value
	}
	return s
}
