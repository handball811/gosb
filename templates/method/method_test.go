package method_test

import (
	"text/template"

	"github.com/handball811/gosb/templates/method"
	"github.com/handball811/gosb/templates/test_helper"
)

func ExampleMethod() {
	tmpl := template.Must(method.METHOD_TEMPLATE.Parse(`{{template "method" .}}`))
	m := dummyMethod()
	test_helper.RunCase(tmpl, []*method.Method{
		&m,
	})

	// Output:
	// func(_f *factory) start(
	//     duration time.Duration,
	//     f func() error,
	// ) (int, error) {
	// return nil
	// }
}
