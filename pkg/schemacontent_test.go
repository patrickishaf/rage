package rage

import (
	"testing"
)

// Make sure it does not add multiple fields with the same name
func TestAddFieldAddsFieldWithNameAndType(t *testing.T) {
	schema := NewSchema().AddField("hi", String)

	found := false

	for _, field := range schema.content.fields {
		if field.name == "hi" && field.dataType == "string" {
			found = true
		}
	}

	if found == false {
		t.Error("AddField does not add an appropriate field to the schema")
	}
}

func TestAddFieldRejectsDuplicateFields(t *testing.T) {
	newSchema := NewSchema().AddField("hi", String).AddField("hi", String)

	counter := 0

	for _, field := range newSchema.content.fields {
		if field.name == "hi" {
			counter++
		}
	}

	if counter > 1 {
		t.Error("AddField can add duplicate fields to the schema")
	}
}
