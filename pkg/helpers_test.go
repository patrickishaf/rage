package rage

import "testing"

func TestFieldNameIsNumericType(t *testing.T) {
	isNum := fieldNameIsNumericType("string")

	if isNum == true {
		t.Errorf("expected %t; got %t", false, true)
	}
}
