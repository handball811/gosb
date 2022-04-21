{{ define "new_struct" -}}
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
}
{{end}}