package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

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

	err := templates.OutputTemplate(filepath.Join(folderPath(), "./source"), output)
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
