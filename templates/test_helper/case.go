package test_helper

import (
	"bytes"
	"log"
	"os"
	"testing"
	"text/template"
)

type Case struct {
	Name string
	Arg  any
	Exp  string
}

func RunTest(
	t *testing.T,
	tmpl *template.Template,
	cases []Case,
) {
	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			buf := bytes.NewBufferString("")
			err := tmpl.Execute(buf, c.Arg)
			if err != nil {
				t.Fatalf("template exec error: %v", err)
			}
			actual := buf.String()
			if actual != c.Exp {
				t.Fatalf("incorrect template result \nexpected:\n%v\nactual:\n%v\n", c.Exp, actual)
			}
		})
	}
}

func RunCase[T any](
	tmpl *template.Template,
	cases []T,
) {
	for _, c := range cases {
		err := tmpl.Execute(os.Stdout, c)
		if err != nil {
			log.Fatalf("executing template: %v", err)
		}
	}
}
