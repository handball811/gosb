{{ define "imports" -}}
{{ $size := len . -}}
{{ if eq $size 1 -}}
import "{{ index . 0 -}}"
{{else if gt $size 1 -}}
import (
{{- range $i, $v := .}}
    "{{$v}}"
{{- end}}
)
{{- end -}}
{{end}}