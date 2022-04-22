package templates

import "github.com/handball811/gosb/templates/parse"

const MethodName = "method"

var MethodTmpl = parse.GenerateTemplate(
	MethodName,
	`
func({{.NameVar}} *{{.Name}}) {{.FuncName}}(
{{- range $i, $arg := .Args }}
	{{ $arg.Name }} {{ $arg.Type }},{{if $arg.Comment}} // {{ $arg.Comment }}{{end}}
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
	Args     []Arg
	Returns  []string
	Body     Body
}
