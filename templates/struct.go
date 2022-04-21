package templates

import "github.com/handball811/gosb/templates/parse"

const StructName = "struct"

var StructTmpl = parse.GenerateTemplate(
	StructName,
	`
type {{.Name}} struct {
{{- range $name, $type := .Fields }}
	{{$name}} {{$type}}
{{- end}}
}`)

type Struct struct {
	Name   string
	Fields map[string]string
}
