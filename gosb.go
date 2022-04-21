package main

import "github.com/urfave/cli/v2"

var app = &cli.App{
	Action: generateFactory,
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:     "type",
			Usage:    "type name which you want to generate factory",
			Required: true,
		},
	},
}

func main() {
}

func generateFactory(ctx *cli.Context) error {
	// TODO
	// ページ内の型情報を取得する
	// 型情報を元にFactory, Builderを作成する
	return nil
}
