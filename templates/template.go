package templates

import (
	"bytes"
	"go/format"
	"os"
	"strings"

	"github.com/handball811/gosb/templates/parse"
)

var TemplateTmpl = strings.Trim(`
// Code generated by gosb.go; DO NOT EDIT.

package {{ .PackageName }}

{{ template "imports" .Imports }}

{{ template "struct" .Factory.Struct}}

{{ template "new_struct" .Factory.NewStruct}}

{{ template "method" .Methods.NewStruct}}
`, "\n")

var AllTemplates = []string{
	TemplateTmpl,
	BodyNewStructTmpl,
	ImportsTmpl,
	MethodTmpl,
	NewStructTmpl,
	ReturnsTmpl,
	StructTmpl,
}

type Template struct {
	PackageName string
	Imports     []string
	Factory     TemplateFactory
	Methods     TemplateMethods
}

type TemplateFactory struct {
	Struct    Struct
	NewStruct NewStruct
}

type TemplateMethods struct {
	NewStruct Method
}

func OutputTemplate(
	templatePath string,
	output string,
	template *Template,
) error {
	file, err := os.Create(output)
	if err != nil {
		return err
	}
	defer file.Close()

	var buf bytes.Buffer

	tmpl := parse.SetUpTemplate(AllTemplates...)
	if err := tmpl.Execute(&buf, template); err != nil {
		return err
	}
	p, err := format.Source(buf.Bytes())
	if err != nil {
		return err
	}
	file.Write(p)
	return nil
}
