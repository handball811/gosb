package method

import (
	"text/template"

	"github.com/handball811/gosb/templates/test_helper"
)

type ExampleInlineStruct struct {
	Name           string
	InlineTemplate func(string, any) string
	Data           any
}

func ExampleInlineTemplate() {
	tmpl := template.Must(template.New("").Parse(`{{ call .InlineTemplate "returns" .Data }}`))
	test_helper.RunCase(tmpl, []ExampleInlineStruct{
		{
			Name:           "returns",
			InlineTemplate: InlineTemplate,
			Data: []string{
				"int",
				"string",
			},
		},
	})
	// Output:
	// (int, string)
}
