package rage

import (
	"fmt"
	"reflect"
	"testing"
)

type composeReasonTestCase struct {
	v         string
	m         string
	vExpected string
	mExpected string
}

var testCases = []composeReasonTestCase{
	{v: "bas", m: "ter", vExpected: "bas", mExpected: "ter"},
	{v: "No string", m: "The object is supposed to have a string field named Color", vExpected: "No string", mExpected: "The object is supposed to have a string field named Color"},
}

func TestComposeReason(t *testing.T) {
	for _, testCase := range testCases {
		result := ComposeReason(testCase.v, testCase.m)
		fmt.Println(reflect.TypeOf(result))
		if result.Value != testCase.vExpected {
			t.Errorf("expected result.Value to be %s but got %s", testCase.vExpected, result.Value)
		}
		if result.Message != testCase.mExpected {
			t.Errorf("expected result.Message to be %s but got %s", testCase.mExpected, result.Message)
		}
	}
}
