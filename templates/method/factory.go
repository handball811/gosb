package method

import (
	"bytes"
	"fmt"
	"text/template"
)

var METHOD_TEMPLATE = template.Must(template.ParseGlob("*.gtpl"))

func InlineTemplate(name string, param any) string {
	buf := bytes.NewBuffer([]byte{})
	tmpl := template.Must(
		METHOD_TEMPLATE.Parse(fmt.Sprintf(`{{template "%s" .}}`, name)))
	if err := tmpl.Execute(buf, param); err != nil {
		panic(err)
	}
	return buf.String()
}
