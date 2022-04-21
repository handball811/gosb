{{ define "struct" -}}
type {{.Name}} struct {
{{- range $name, $type := .Fields }}
    {{$name}} {{$type}}
{{- end}}
}
{{end}}
