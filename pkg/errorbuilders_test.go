package rage

import "testing"

func TestBuildFieldNotFoundError(t *testing.T) {
	rawErrorMessage := "the field 'age' was not found in the object"
	derivedErrorMessage := buildFieldNotFoundError("age")

	if derivedErrorMessage != rawErrorMessage {
		t.Errorf("failed to construct FieldNotFound error message. \nexpected %s\ngot %s", rawErrorMessage, derivedErrorMessage)
	}
}
