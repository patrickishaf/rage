package rage

import (
	"fmt"
	"reflect"
)

type Schema struct {
	content *SchemaContent
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
	result := ValidationResult{}
	isValid := true
	for _, field := range s.content.fields {
		/*
			TODO: If the field is not required, you still need to validate it to ensure that it has all the right properties
			Right now, this code means that if the current field is not required, you just jump to validating the next field

			I haven't decided if this is good behaviour or not but in the meantime, I'll leave this behaviour like this
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
				Message: buildFieldNotFoundError(field.name),
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
		_, err := getValueOfFieldWithName(obj, s.content.fields[len(s.content.fields)-1].name)
		if err != nil {
			isValid = false
			result.Reasons = append(result.Reasons, ValidationReason{
				Value:   field.name,
				Message: buildFieldNotFoundError(field.name),
			})
			continue
		}
		// fmt.Println("dataTypeOfSchema => " + dataTypeOfSchema + "\nvalue of field with name " + field.name + " => ")
		// fmt.Println(value)
		// fmt.Println("")
		// If thid doesn't work I might have to convert all numeric types to float64 and store them like that
		if fieldNameIsNumericType(field.dataType) {
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
	// Validate the name of the field
	// validate the data type of the field
	// validate the data type of the field
	// validate the min value of the field
	// Validate the max value of the field
	// Validate the min length of the field
	// Validate the max length of the field
	result.IsValid = isValid
	return result
}

func (s *Schema) AddField(fieldName string, dataType FieldType) *Schema {
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

func (s *Schema) SetDataType(dataType FieldType) *Schema {
	if len(s.content.fields) == 0 {
		s.content.fields = append(s.content.fields, schemaField{})
	}
	s.content.fields[len(s.content.fields)-1].dataType = dataType
	return s
}

func (s *Schema) SetIsRequired(isRequired bool) *Schema {
	if len(s.content.fields) == 0 {
		s.content.fields = append(s.content.fields, schemaField{})
	}
	s.content.fields[len(s.content.fields)-1].isRequired = isRequired
	return s
}

func (s *Schema) SetMinValue(value float64) *Schema {
	if len(s.content.fields) == 0 {
		s.content.fields = append(s.content.fields, schemaField{})
	}
	currentField := &(s.content.fields[len(s.content.fields)-1])
	if fieldNameIsNumericType((*currentField).dataType) {
		(*currentField).min = value
	}
	return s
}

func (s *Schema) SetMaxValue(value float64) *Schema {
	if len(s.content.fields) == 0 {
		s.content.fields = append(s.content.fields, schemaField{})
	}
	currentField := &(s.content.fields[len(s.content.fields)-1])
	if fieldNameIsNumericType((*currentField).dataType) {
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
		s.content.fields[len(s.content.fields)-1].minLength = value
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
		s.content.fields[len(s.content.fields)-1].maxLength = value
	}
	return s
}
