package templates

import "github.com/handball811/gosb/templates/parse"

const MethodName = "method"

var MethodTmpl = parse.GenerateTemplate(
	MethodName,
	`
func({{.NameVar}} *{{.Name}}) {{.FuncName}}(
{{- range $name, $type := .Args }}
	{{$name}} {{$type}},
{{- end}}
) {{template "returns" .Returns }} {
{{- if .Body}}
{{- InlineTemplate .Body.Name .}}
{{- end}}
}
`)

type Config[T any] struct {
	Method
	Body T
}

type Method struct {
	NameVar  string
	Name     string
	FuncName string
	Args     map[string]string
	Returns  []string
	Body     Body
}
