{{- define "method"}}
func({{.NameVar}} *{{.Name}}) {{.FuncName}}(
{{- range $name, $type := .Args }}
    {{$name}} {{$type}},
{{- end}}
) {{template "returns" .Returns }} {
{{- call .InlineTemplate .Body.Name .}}
}
{{end}}