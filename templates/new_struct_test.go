package templates_test

import (
	"text/template"

	"github.com/handball811/gosb/templates"
	"github.com/handball811/gosb/templates/test_helper"
)

func ExampleNewStruct() {
	tmpl := template.Must(template.ParseGlob("new_struct.gtpl"))
	tmpl, _ = tmpl.Parse(`{{template "new_struct" .}}`)
	test_helper.RunCase(tmpl, []templates.NewStruct{
		{
			FuncName: "NewStructFactory",
			Struct: templates.Struct{
				Name: "structFactory",
				Fields: map[string]string{
					"name":   "string",
					"number": "*int",
				},
			},
		},
		{
			FuncName: "NewNoneFactory",
			Struct: templates.Struct{
				Name:   "noneFactory",
				Fields: map[string]string{},
			},
		},
	})

	// Output:
	// func NewStructFactory(
	//     name string,
	//     number *int,
	// ) *structFactory {
	//     return &structFactory{
	//         name: name,
	//         number: number,
	//     }
	// }
	// func NewNoneFactory(
	// ) *noneFactory {
	//     return &noneFactory{
	//     }
	// }
}
