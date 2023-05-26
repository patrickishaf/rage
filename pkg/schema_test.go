package rage

import (
	"reflect"
	"testing"
)

func TestNewSchema(t *testing.T) {
	e := NewSchema()
	a := &Schema{
		content: &SchemaContent{
			fields: []schemaField{},
		},
	}
	if reflect.TypeOf(e) != reflect.TypeOf(a) {
		t.Error("New Schema does not create an object of type &Schema")
	}
}

func TestAddField(t *testing.T) {
	testSchema := NewSchema().AddField("namawa", Float32)

	if testSchema.content.fields[0].name != "namawa" && testSchema.content.fields[0].dataType != "float64" {
		t.Errorf("failed to add field properly. \nexpected name: 'namawa' ----- actual name: %s \nexpected dataType: 'float64' ----- actual dataType: %s", testSchema.content.fields[0].name, testSchema.content.fields[0].dataType)
	}
}

func TestValidate(t *testing.T) {
	schema := NewSchema().AddField("name", Float32).SetMinLength(45).AddField("age", Float32).SetMinValue(0).SetMaxValue(100)

	type person struct {
		name string
		age  int
	}

	pete := person{
		name: "Pete",
	}

	goodness := person{
		name: "Goodness",
		age:  24,
	}

	peteResult := schema.Validate(pete)
	goodnessResult := schema.Validate(goodness)

	if peteResult.IsValid == false {
		t.Error("the object is not valid")
	}

	if !goodnessResult.IsValid {
		t.Error("the object is not valid")
	}

	// TEST THAT SETTING THE MAX AND MIN VALUE ON A NON-NUMERIC TYPE WILL NOT WORK
	// TEST THAT SETTING THE MAX AND MIN VALUE ON A NUMERIC TYPE WILL NOT WORK
	// TEST THAT SETTING THE MAX AND MIN LENGTH ON A NUMERIC TYPE WILL NOT WORK
	// TEST THAT SETTING THE MAX AND MIN LENGTH ON A NUMERIC TYPE WILL WORK
}

func TestSetMaxLength(t *testing.T) {
	s := NewSchema().AddField("name", String).SetMaxLength(100)

	if s.content.fields[0].maxLength != 100 {
		t.Errorf("failed to set max length. expected %d but got %d", 100, s.content.fields[0].maxLength)
	}
}

func TestSetMinLength(t *testing.T) {
	s := NewSchema().AddField("name", String).SetMinLength(10)

	if s.content.fields[0].minLength != 10 {
		t.Errorf("failed to set min length. expected %d but got %d", 100, s.content.fields[0].minLength)
	}
}

func TestSetMaxValue(t *testing.T) {
	maxVal := 100.74
	s := NewSchema().AddField("balance", Float32).SetMaxValue(maxVal)

	if s.content.fields[0].max != maxVal {
		t.Errorf("failed to set max length. expected %f but got %f", maxVal, s.content.fields[0].max)
	}
}

func TestSetMinValue(t *testing.T) {
	minVal := 10.75
	s := NewSchema().AddField("balance", Float32).SetMinValue(minVal)

	if s.content.fields[0].min != minVal {
		t.Errorf("failed to set max length. expected %f but got %f", minVal, s.content.fields[0].max)
	}
}
