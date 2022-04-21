package method_test

import (
	"text/template"

	"github.com/handball811/gosb/templates/method"
	"github.com/handball811/gosb/templates/test_helper"
)

func ExampleBodyNewStruct() {
	tmpl := template.Must(method.METHOD_TEMPLATE.Parse(`{{template "body_new_struct" .}}`))

	m := dummyMethod()
	m.Body = &method.BodyNewStruct{
		VarPrefix: "xxxReal_",
		Fields: []method.Field{
			dummyFields["key"],
			dummyFields["note"],
			dummyFields["age"],
			dummyFields["data"],
		},
	}
	test_helper.RunCase(tmpl, []method.Method{m})

	// Output:
	// // key
	// var xxxReal_key string
	// xxxReal_key = key
	//
	// // note
	// var xxxReal_note *string
	// if note == nil {
	//     xxxReal_note = _f.noteDefault()
	// } else {
	//     xxxReal_note = note
	// }
	//
	// // age
	// var xxxReal_age int
	// if age == nil {
	//     xxxReal_age = _f.defaultAge()
	// } else {
	//     xxxReal_age = age
	// }
	// if err := _f.ageValidation(xxxReal_age); err != nil {
	//     return nil, fmt.Errorf("`age` validation error: %v", err)
	// }
	//
	// // data
	// var xxxReal_data *Data
	// xxxReal_data = data
	// if err := _f.dataValidation(xxxReal_data); err != nil {
	//     return nil, fmt.Errorf("`data` validation error: %v", err)
	// }
	//
	// return &factory {
	//     key: xxxReal_key,
	//     note: xxxReal_note,
	//     age: xxxReal_age,
	//     data: xxxReal_data,
	// }, nil
}
