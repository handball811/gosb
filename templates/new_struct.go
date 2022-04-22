package templates

import "github.com/handball811/gosb/templates/parse"

const NewStructName = "new_struct"

var NewStructTmpl = parse.GenerateTemplate(
	NewStructName,
	`
func {{.FuncName}}(
{{- range $name, $type := .Fields }}
	{{$name}} {{$type}},
{{- end}}
) *{{.Name}} {
	return &{{.Name}}{
{{- range $name, $type := .Fields }}
		{{$name}}: {{$name}},
{{- end}}
	}
}`)

type NewStruct struct {
	Struct
	FuncName string
}
