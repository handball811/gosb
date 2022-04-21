{{ define "body_new_struct" }}
{{range $i, $field := .Body.Fields}}
// {{$field.Name}}
{{- $tempVar := (printf "%s%s" $.Body.VarPrefix $field.Name)  }}
var {{$tempVar}} {{$field.Type}}
{{- if $field.DefaultFunc}}
if {{$field.Name}} == nil {
    {{$tempVar}} = {{$.NameVar}}.{{$field.DefaultFunc}}()
} else {
    {{$tempVar}} = {{if and $field.Optional (not $field.Pointer) -}}*{{- end -}}{{$field.Name}}
}
{{- else}}
{{$tempVar}} = {{$field.Name}}
{{- end }}
{{- if $field.ValidationFunc}}
if err := {{$.NameVar}}.{{$field.ValidationFunc}}({{$tempVar}}); err != nil {
    return nil, fmt.Errorf("`{{$field.Name}}` validation error: %v", err)
}
{{- end}}
{{end}}
return &{{.Name}} {
{{- range $i, $field := .Body.Fields}}
    {{- $tempVar := printf "%s%s" $.Body.VarPrefix $field.Name }}
    {{$field.Name}}: {{$tempVar}},
{{- end}}
}, nil
{{end}}
