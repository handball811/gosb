package templates

import "github.com/handball811/gosb/templates/parse"

const ReturnsName = "returns"

var ReturnsTmpl = parse.GenerateTemplate(
	ReturnsName,
	`
{{- $size := len . -}}
{{- if eq $size 1 -}}
{{- index . 0 -}}
{{else if gt $size 1 -}}
({{- range $i, $v := . -}}{{if ne $i 0 -}}, {{end}}{{$v}}{{- end}})
{{- end -}}
`)
