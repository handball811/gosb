package main

import (
	"fmt"
	"go/ast"
	"go/constant"
	"go/token"
	"go/types"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/handball811/gosb/analysis"
	"github.com/handball811/gosb/templates"
	"github.com/urfave/cli/v2"
)

var app = &cli.App{
	Action: generateFactory,
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:     "type",
			Usage:    "type name which you want to generate factory",
			Required: true,
		},
		&cli.StringFlag{
			Name:  "output",
			Usage: "set output filename",
		},
	},
}

// Value represents a declared constant.
type Value struct {
	originalName string // The name of the constant.
	name         string // The name with trimmed prefix.
	// The value is stored as a bit pattern alone. The boolean tells us
	// whether to interpret it as an int64 or a uint64; the only place
	// this matters is when sorting.
	// Much of the time the str field is all we need; it is printed
	// by Value.String.
	value  uint64 // Will be converted to int64 when needed.
	signed bool   // Whether the constant is a signed type.
	str    string // The string representation given by the "go/constant" package.
}

// File holds a single parsed file and associated data.
type File struct {
	pkg  *Package  // Package to which this file belongs.
	file *ast.File // Parsed AST.
	// These fields are reset for each type being generated.
	typeName string  // Name of the constant type.
	values   []Value // Accumulator for constant values of that type.

	trimPrefix  string
	lineComment bool
}

// genDecl processes one declaration clause.
func (f *File) genDecl(node ast.Node) bool {
	decl, ok := node.(*ast.GenDecl)
	if !ok || decl.Tok != token.CONST {
		// We only care about const declarations.
		return true
	}
	// The name of the type of the constants we are declaring.
	// Can change if this is a multi-element declaration.
	typ := ""
	// Loop over the elements of the declaration. Each element is a ValueSpec:
	// a list of names possibly followed by a type, possibly followed by values.
	// If the type and value are both missing, we carry down the type (and value,
	// but the "go/types" package takes care of that).
	for _, spec := range decl.Specs {
		vspec := spec.(*ast.ValueSpec) // Guaranteed to succeed as this is CONST.
		if vspec.Type == nil && len(vspec.Values) > 0 {
			// "X = 1". With no type but a value. If the constant is untyped,
			// skip this vspec and reset the remembered type.
			typ = ""

			// If this is a simple type conversion, remember the type.
			// We don't mind if this is actually a call; a qualified call won't
			// be matched (that will be SelectorExpr, not Ident), and only unusual
			// situations will result in a function call that appears to be
			// a type conversion.
			ce, ok := vspec.Values[0].(*ast.CallExpr)
			if !ok {
				continue
			}
			id, ok := ce.Fun.(*ast.Ident)
			if !ok {
				continue
			}
			typ = id.Name
		}
		if vspec.Type != nil {
			// "X T". We have a type. Remember it.
			ident, ok := vspec.Type.(*ast.Ident)
			if !ok {
				continue
			}
			typ = ident.Name
		}
		if typ != f.typeName {
			// This is not the type we're looking for.
			continue
		}
		// We now have a list of names (from one line of source code) all being
		// declared with the desired type.
		// Grab their names and actual values and store them in f.values.
		for _, name := range vspec.Names {
			if name.Name == "_" {
				continue
			}
			// This dance lets the type checker find the values for us. It's a
			// bit tricky: look up the object declared by the name, find its
			// types.Const, and extract its value.
			obj, ok := f.pkg.defs[name]
			if !ok {
				log.Fatalf("no value for constant %s", name)
			}
			info := obj.Type().Underlying().(*types.Basic).Info()
			if info&types.IsInteger == 0 {
				log.Fatalf("can't handle non-integer constant type %s", typ)
			}
			value := obj.(*types.Const).Val() // Guaranteed to succeed as this is CONST.
			if value.Kind() != constant.Int {
				log.Fatalf("can't happen: constant is not an integer %s", name)
			}
			i64, isInt := constant.Int64Val(value)
			u64, isUint := constant.Uint64Val(value)
			if !isInt && !isUint {
				log.Fatalf("internal error: value of %s is not an integer: %s", name, value.String())
			}
			if !isInt {
				u64 = uint64(i64)
			}
			v := Value{
				originalName: name.Name,
				value:        u64,
				signed:       info&types.IsUnsigned == 0,
				str:          value.String(),
			}
			if c := vspec.Comment; f.lineComment && c != nil && len(c.List) == 1 {
				v.name = strings.TrimSpace(c.Text())
			} else {
				v.name = strings.TrimPrefix(v.originalName, f.trimPrefix)
			}
			f.values = append(f.values, v)
		}
	}
	return false
}

type Package struct {
	name  string
	defs  map[*ast.Ident]types.Object
	files []*File
}

func (p *Package) getStruct(tp string) (types.Object, *types.Struct) {
	var obj types.Object
	for ident, def := range p.defs {
		if ident.Name == tp {
			obj = def
		}
	}
	s, ok := obj.Type().Underlying().(*types.Struct)
	if !ok {
		log.Fatal("not struct")
	}
	return obj, s
}

func main() {
	app.Run(os.Args)
}

// generateFactory generates the struct factory for the struct you specified
func generateFactory(ctx *cli.Context) error {
	// TODO
	// ページ内の型情報を取得する
	// 型情報を元にFactory, Builderを作成する
	// setup
	args := ctx.Args().Slice()
	if len(args) == 0 {
		args = append(args, ".")
	}
	types := strings.Split(ctx.String("type"), ",")
	output := ctx.String("output")

	var dir string
	if len(args) == 1 && isDirectory(args[0]) {
		dir = args[0]
	} else {
		dir = filepath.Dir(args[0])
	}

	if output == "" {
		baseName := fmt.Sprintf("%s_factory.go", types[0])
		output = filepath.Join(dir, strings.ToLower(baseName))
	}

	pkg := analysis.GeneratePackage(args)

	err := templates.OutputTemplate(filepath.Join(folderPath(), "./source"), output, getTemplate(pkg, types[0]))
	if err != nil {
		panic(err)
	}

	return nil
}

// folderPath can retrieve file path of calling text
func folderPath() string {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)
	return exPath
}

// isDirectory reports whether the named file is a directory.
func isDirectory(name string) bool {
	info, err := os.Stat(name)
	if err != nil {
		log.Fatal(err)
	}
	return info.IsDir()
}

func getTemplate(pkg *analysis.Package, tp string) *templates.Template {

	s := pkg.FindStruct(tp)
	if s == nil {
		log.Fatal("Struct name not found")
	}

	astr := generateFactoryStruct(s)
	srct := getStruct(astr)

	return &templates.Template{
		PackageName: pkg.Name,
		Imports:     getImports(s),
		Factory: templates.TemplateFactory{
			Struct:    srct.Struct,
			NewStruct: srct,
		},
		Methods: getMethod(s, astr),
	}
}

func getStruct(s *analysis.Struct) templates.NewStruct {
	fields := make(map[string]string)
	for _, field := range s.Fields {
		fields[field.Name] = field.DeclareType
	}
	return templates.NewStruct{
		FuncName: "New" + s.Name,
		Struct: templates.Struct{
			Name:   s.Name,
			Fields: fields,
		},
	}
}

const (
	ValidationSuffix = "Vld"
	DefaultSuffix    = "Def"
)

func generateFactoryStruct(base *analysis.Struct) *analysis.Struct {
	s := &analysis.Struct{
		Name: base.Name + "Factory",
	}
	fields := make([]analysis.Field, 0)
	for _, field := range base.Fields {
		if Contains(field.Options, analysis.OP_Default) {
			fields = append(fields, analysis.Field{
				Name:        CamelCase(field.Name + DefaultSuffix),
				DeclareType: fmt.Sprintf("func() %s", field.DeclareType),
				Struct:      s,
			})
		}
		if Contains(field.Options, analysis.OP_Validation) {
			fields = append(fields, analysis.Field{
				Name:        CamelCase(field.Name + ValidationSuffix),
				DeclareType: fmt.Sprintf("func(%s) error", field.DeclareType),
				Struct:      s,
			})
		}
	}
	s.Fields = fields
	return s
}

func Contains[T comparable](s []T, t T) bool {
	for _, ss := range s {
		if ss == t {
			return true
		}
	}
	return false
}

func CamelCase(s string) string {
	if s[0] >= 'A' && s[0] <= 'Z' {
		s = fmt.Sprintf("%c%s", (s[0] - 'A' + 'a'), s[1:])
	}
	return s
}

func getMethod(
	base *analysis.Struct,
	factory *analysis.Struct,
) templates.TemplateMethods {
	args := make([]templates.Arg, 0, len(base.Fields))
	fields := make([]templates.Field, 0)

	for _, field := range base.Fields {
		name := field.Name
		tp := field.DeclareType
		if !strings.HasPrefix(tp, "*") && (Contains(field.Options, analysis.OP_Optional)) {
			tp = "*" + tp
		}
		args = append(args, templates.Arg{
			Name: CamelCase(name),
			Type: tp,
			Comment: strings.Join(analysis.Map(field.Options, func(t analysis.Option) (string, bool) {
				return t.String(), true
			}), ","),
		})

		defFunc := ""
		if Contains(field.Options, analysis.OP_Default) {
			defFunc = CamelCase(field.Name + DefaultSuffix)
		}

		valFunc := ""
		if Contains(field.Options, analysis.OP_Validation) {
			valFunc = CamelCase(field.Name + ValidationSuffix)
		}

		fields = append(fields, templates.Field{
			Name:           name,
			VarName:        CamelCase(name),
			Type:           field.DeclareType,
			Pointer:        strings.HasPrefix(field.DeclareType, "*"),
			Optional:       Contains(field.Options, analysis.OP_Optional),
			DefaultFunc:    defFunc,
			ValidationFunc: valFunc,
		})
	}

	return templates.TemplateMethods{
		NewStruct: templates.Method{
			NameVar:  "_f",
			Name:     factory.Name,
			FuncName: "New" + base.Name,
			Args:     args,
			Returns: []string{
				"*" + base.Name,
				"error",
			},
			Body: &templates.BodyNewStruct{
				Struct:    base.Name,
				VarPrefix: "xxx",
				Fields:    fields,
			},
		},
	}
}

func getImports(s *analysis.Struct) []string {
	base := make([]string, 0)
	base = append(base, s.File.Imports...)
	for _, f := range s.Fields {
		if Contains(f.Options, analysis.OP_Validation) {
			if !Contains(base, "fmt") {
				base = append(base, "fmt")
			}
		}
	}
	return base
}
