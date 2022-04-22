package templates

import "github.com/handball811/gosb/templates/parse"

const ImportsName = "imports"

var ImportsTmpl = parse.GenerateTemplate(
	ImportsName,
	`
{{- $size := len . -}}
{{ if eq $size 1 -}}
import "{{ index . 0 -}}"
{{else if gt $size 1 -}}
import (
{{- range $i, $v := .}}
    "{{$v}}"
{{- end}}
)
{{- end -}}
`)

type Imports []string
