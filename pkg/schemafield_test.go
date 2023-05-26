package rage

import (
	"reflect"
	"testing"
)

func TestSetName(t *testing.T) {
	field := schemaField{}
	field.setName("test")

	if field.name != "test" {
		t.Errorf("Does not set name accurately. Expected 'test' but got '%s'", field.name)
	}
}
func TestSetDataType(t *testing.T) {
	field := schemaField{}
	field.setDataType(reflect.TypeOf("Myron").String())

	if field.dataType != String {
		t.Errorf("Does not set data type accurately. Expected 'string' but got '%s'", field.dataType)
	}
}
func TestSetIsRequired(t *testing.T) {
	field := schemaField{}
	field.setIsRequired(true)

	if !field.isRequired {
		t.Errorf("Does not set isRequired accurately. Expected 'true' but got '%t'", field.isRequired)
	}
}
func TestSetMin(t *testing.T) {
	field := schemaField{}
	field.setMin(34.56)

	if field.min != 34.56 {
		t.Errorf("Does not set min accurately. Expected '34.56' but got '%f'", field.min)
	}
}
func TestSetMax(t *testing.T) {
	field := schemaField{}
	field.setMax(123.456)

	if field.max != 123.456 {
		t.Errorf("Does not set max accurately. Expected '123.456' but got '%f'", field.max)
	}
}
func TestSetMinLen(t *testing.T) {
	field := schemaField{}
	field.setMinLen(10)

	if field.minLength != 10 {
		t.Errorf("Does not set name accurately. Expected 'myron' but got '%d'", field.minLength)
	}
}
func TestSetMaxLen(t *testing.T) {
	field := schemaField{}
	field.setMaxLen(100)

	if field.maxLength != 100 {
		t.Errorf("Does not set name accurately. Expected 'myron' but got '%d'", field.maxLength)
	}
}
