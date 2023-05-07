package main

import (
	"reflect"

	rage "github.com/patrickishaf/rage/pkg"
)

func main() {
	schema := rage.NewSchema().AddField("name", reflect.TypeOf("name")).SetIsRequired(true).AddField("age", reflect.TypeOf(1)).SetIsRequired(true)

	type person struct {
		name string
	}

	pete := person{
		name: "Pete",
	}

	schema.Validate(pete)
}
