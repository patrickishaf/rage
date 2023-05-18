package rage

import (
	"fmt"
	"reflect"
	"testing"
)

func TestValidate(t *testing.T) {
	schema := NewSchema().AddField("name", reflect.TypeOf(12.56)).SetMinLength(45).AddField("age", reflect.TypeOf(1)).SetMinValue(0).SetMaxValue(100)

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

	fmt.Println(peteResult)
	fmt.Println(goodnessResult)

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
	// personSchema := NewSchema().AddField("age", reflect.TypeOf(23)).SetIsRequired(true).SetMaxValue(200)

	// person := struct{ age int }{age: 201}

	// vResult := personSchema.Validate(person)
	// fmt.Println(vResult)

	// if vResult.IsValid {
	// 	t.Error("the max value validation does not work")
	// }
}

func TestSetMinLength(t *testing.T) {}

func TestSetMaxValue(t *testing.T) {}

func TestSetMinValue(t *testing.T) {}
