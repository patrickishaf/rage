package main

import (
	"reflect"

	"github.com/patrickishaf/rage/pkg/validate"
)

func main() {
	schema := validate.NewSchema().AddField("name", reflect.TypeOf("name")).SetIsRequired(true).AddField("age", reflect.TypeOf(1)).SetIsRequired(true)

	type person struct {
		name string
	}

	pete := person{
		name: "Pete",
	}

	schema.Validate(pete)
}
