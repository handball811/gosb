package templates_test

import (
	"text/template"

	"github.com/handball811/gosb/templates"
	"github.com/handball811/gosb/templates/test_helper"
)

func ExampleStruct() {
	tmpl := template.Must(template.ParseGlob("struct.gtpl"))
	tmpl, _ = tmpl.Parse(`{{template "struct" .}}`)
	test_helper.RunCase(tmpl, []templates.Struct{
		{
			Name: "structFactory",
			Fields: map[string]string{
				"name":   "string",
				"number": "*int",
			},
		},
		{
			Name:   "noneFactory",
			Fields: map[string]string{},
		},
	})

	// Output:
	// type structFactory struct {
	//     name string
	//     number *int
	// }
	// type noneFactory struct {
	// }
}
