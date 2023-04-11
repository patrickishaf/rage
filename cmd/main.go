package main

import (
	"reflect"

	"github.com/patrickishaf/rage/pkg/validate"
)

func main() {
	schema := validate.NewSchema().AddField("name", reflect.TypeOf("name")).AddField("age", reflect.TypeOf(1))

	type person struct {
		name string
	}

	pete := person{
		name: "Pete",
	}

	schema.Validate(pete)
}
