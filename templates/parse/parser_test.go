package parse_test

import (
	"text/template"

	"github.com/handball811/gosb/templates/parse"
	"github.com/handball811/gosb/templates/test_helper"
)

type ExampleInlineStruct struct {
	Name           string
	InlineTemplate func(string, any) string
	Data           any
}

func ExampleGenerateInlineTemplate() {
	tmpl := template.Must(template.New("").Parse(`{{ call .InlineTemplate "returns" .Data }}`))
	test_helper.RunCase(tmpl, []ExampleInlineStruct{
		{
			Name:           "returns",
			InlineTemplate: parse.GenerateInlineTemplate(parse.ParseTemplates("../../source")),
			Data: []string{
				"int",
				"string",
			},
		},
	})
	// Output:
	// (int, string)
}
