package analysis

import (
	"fmt"
	"go/ast"
	"go/token"
	"go/types"
	"log"
	"reflect"
	"strconv"
	"strings"

	"golang.org/x/tools/go/packages"
)

//go:generate enumer -type=Option -linecomment -json
type Option int

const (
	OP_Unknown    Option = iota // unknown
	OP_Optional                 // optional
	OP_Nillable                 // nillable
	OP_Validation               // validation
	OP_Default                  // default
)

type Field struct {
	Name        string // フィールドの名称
	DeclareType string // 宣言時の型
	Options     []Option
	Struct      *Struct
}

type Struct struct {
	Name   string
	Fields []Field
	File   *File
}

type File struct {
	Imports []string
	Structs []Struct
	Package *Package

	AstFile *ast.File
}

type Package struct {
	Name  string
	Files []File
}

func GeneratePackage(
	patterns []string,
) *Package {
	cfg := &packages.Config{
		Mode: packages.NeedTypes | packages.NeedTypesInfo | packages.NeedSyntax | packages.NeedImports | packages.NeedName,
	}

	pkgs, err := packages.Load(cfg, patterns...)
	if err != nil {
		log.Fatal(err)
	}

	if len(pkgs) != 1 {
		log.Fatalf("error: %d packages found", len(pkgs))
	}

	p := pkgs[0]
	// ast.Print(p.Fset, p)

	// lcfg := loader.Config{
	// 	Fset: p.Fset,
	// }
	// prog, _ :=lcfg.Load()
	// prog.AllPackages[p.Types].Files[0].

	files := make([]File, 0, len(p.Syntax))
	for _, file := range p.Syntax {
		f := File{
			Imports: Map(file.Imports, func(t *ast.ImportSpec) (string, bool) {
				if t == nil || t.Path == nil {
					return "", false
				}
				return strings.Trim(t.Path.Value, `"`), true
			}),
			Structs: Map(file.Decls, toStruct),
			AstFile: file,
		}
		for i := range f.Structs {
			f.Structs[i].File = &f
		}
		files = append(files, f)
	}

	return &Package{
		Name:  p.Name,
		Files: files,
	}
}

func generateStruct(
	s *types.Struct,
) Struct {
	return Struct{}
}

func toStruct(t ast.Decl) (Struct, bool) {
	var s Struct
	decl, ok := t.(*ast.GenDecl)
	if !ok {
		return s, false
	}
	if decl.Tok != token.TYPE {
		return s, false
	}

	spec := decl.Specs[0].(*ast.TypeSpec)
	s.Name = spec.Name.Name
	st := spec.Type.(*ast.StructType)
	s.Fields = Map(st.Fields.List, func(t *ast.Field) (Field, bool) {
		var field Field
		field.Name = t.Names[0].Name
		field.DeclareType = getFieldType(t.Type)
		if t.Tag != nil {
			field.Options = Map(
				strings.Split(reflect.StructTag(unquote(t.Tag.Value)).Get("sc"), ","),
				func(t string) (Option, bool) {
					op, err := OptionString(strings.Trim(t, " "))
					if err != nil {
						return OP_Unknown, false
					}
					return op, true
				})
		}
		return field, true
	})
	for i := range s.Fields {
		s.Fields[i].Struct = &s
	}
	return s, true
}

func getFieldType(x ast.Expr) string {
	switch tp := x.(type) {
	case *ast.Ident:
		return tp.Name
	case *ast.MapType:
		return fmt.Sprintf("map[%s]%s", getFieldType(tp.Key), getFieldType(tp.Value))
	case *ast.StarExpr:
		return "*" + getFieldType(tp.X)
	case *ast.SelectorExpr:
		return fmt.Sprintf("%s.%s", getFieldType(tp.X), getFieldType(tp.Sel))
	case *ast.ArrayType:
		return fmt.Sprintf("[]%s", getFieldType(tp.Elt))
	}
	return ""
}

func unquote(t string) string {
	s, err := strconv.Unquote(t)
	if err != nil {
		panic(err) // 不正なタグはParseFileでエラーになる
	}
	return s
}

func (p *Package) FindStruct(name string) *Struct {
	for _, file := range p.Files {
		for _, s := range file.Structs {
			if s.Name == name {
				return &s
			}
		}
	}
	return nil
}
