package method_test

import (
	"github.com/handball811/gosb/templates/method"
)

func dummyMethod() method.Method {
	return method.Method{
		NameVar:  "_f",
		Name:     "factory",
		FuncName: "start",
		Args: map[string]string{
			"duration": "time.Duration",
			"f":        "func() error",
		},
		Returns: []string{"int", "error"},
		Body:    &method.BodyDummy{},
	}
}

var dummyFields = map[string]method.Field{
	"key": {
		Name:           "key",
		Type:           "string",
		Pointer:        false,
		Optional:       false,
		DefaultFunc:    "",
		ValidationFunc: "",
	},
	"note": {
		Name:           "note",
		Type:           "*string",
		Pointer:        true,
		Optional:       true,
		DefaultFunc:    "noteDefault",
		ValidationFunc: "",
	},
	"age": {
		Name:           "age",
		Type:           "int",
		Pointer:        false,
		Optional:       false,
		DefaultFunc:    "defaultAge",
		ValidationFunc: "ageValidation",
	},
	"data": {
		Name:           "data",
		Type:           "*Data",
		Pointer:        true,
		Optional:       false,
		DefaultFunc:    "",
		ValidationFunc: "dataValidation",
	},
}
