package templates_test

import (
	"text/template"

	"github.com/handball811/gosb/templates"
	"github.com/handball811/gosb/templates/parse"
	"github.com/handball811/gosb/templates/test_helper"
)

func ExampleImports() {
	tmpl := parse.SetUpTemplate(templates.ImportsTmpl)
	tmpl = template.Must(tmpl.Parse(`{{template "imports" .}}`))

	test_helper.RunCase(tmpl, []templates.Imports{
		templates.Imports([]string{}),
		templates.Imports([]string{
			"a/b/c",
		}),
		templates.Imports([]string{
			"a/b/c",
			"d/e/f",
		}),
	})

	// Output:
	// import "a/b/c"
	// import (
	//     "a/b/c"
	//     "d/e/f"
	// )
}
