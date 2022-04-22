package parse

import (
	"bytes"
	"fmt"
	"strings"
	"text/template"
)

func SetUpTemplate(templates ...string) *template.Template {
	tpl := template.New("")

	funcMap := template.FuncMap{
		"InlineTemplate": GenerateInlineTemplate(tpl),
	}
	tpl = tpl.Funcs(funcMap)
	for _, t := range templates {
		tpl = template.Must(tpl.Parse(t))
	}
	return tpl
}

func GenerateInlineTemplate(tmpl *template.Template) func(string, any) string {
	return func(name string, param any) string {
		buf := bytes.NewBuffer([]byte{})
		tmpl := template.Must(tmpl.Parse(fmt.Sprintf(`{{template "%s" .}}`, name)))
		if err := tmpl.Execute(buf, param); err != nil {
			panic(err)
		}
		return buf.String()
	}
}

const TemplateBase = `
{{- define "%s" }}
%s
{{- end}}
`

func GenerateTemplate(
	name string,
	raw string,
) string {
	return strings.Trim(fmt.Sprintf(TemplateBase, name, strings.Trim(raw, "\n")), "\n")
}
