package templates_test

import (
	"github.com/handball811/gosb/templates"
)

func dummyMethod() templates.Method {
	return templates.Method{
		NameVar:  "_f",
		Name:     "factory",
		FuncName: "start",
		Args: []templates.Arg{
			{
				Name: "duration",
				Type: "time.Duration",
			},
			{
				Name: "f",
				Type: "func() error",
			},
		},
		Returns: []string{"int", "error"},
		Body:    &templates.BodyDummy{},
	}
}

var dummyFields = map[string]templates.Field{
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
