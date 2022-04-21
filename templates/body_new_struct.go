package templates

import "github.com/handball811/gosb/templates/parse"

const BodyNewStructName = "body_new_struct"

var BodyNewStructTmpl = parse.GenerateTemplate(
	BodyNewStructName,
	`
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
	return nil, fmt.Errorf("'{{$field.Name}}' validation error: %v", err)
}
{{- end}}
{{end}}
return &{{.Name}} {
{{- range $i, $field := .Body.Fields}}
	{{- $tempVar := printf "%s%s" $.Body.VarPrefix $field.Name }}
	{{$field.Name}}: {{$tempVar}},
{{- end}}
}, nil
`)

var _ Body = new(BodyNewStruct)

type BodyNewStruct struct {
	VarPrefix string
	Fields    []Field
}

func (b *BodyNewStruct) Name() string {
	return "body_new_struct"
}
