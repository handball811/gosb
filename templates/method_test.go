package templates_test

import (
	"text/template"

	"github.com/handball811/gosb/templates"
	"github.com/handball811/gosb/templates/parse"
	"github.com/handball811/gosb/templates/test_helper"
)

func ExampleMethod() {
	tmpl := template.Must(
		parse.SetUpTemplate(
			templates.MethodTmpl,
			templates.ReturnsTmpl,
			templates.BodyDummyTmpl,
		).Parse(`{{template "method" .}}`))
	m := dummyMethod()
	test_helper.RunCase(tmpl, []*templates.Method{
		&m,
	})

	// Output:
	// func(_f *factory) start(
	// 	duration time.Duration,
	// 	f func() error,
	// ) (int, error) {
	// return nil
	// }
}
