package main

import (
	"fmt"

	"github.com/alecthomas/kingpin"
	"github.com/orivej/e"
	"github.com/orivej/go-nix/nix/parser"
)

var (
	parseCmd     = kingpin.Command("parse", "Parse Nix expression.")
	parseExprArg = parseCmd.Arg("expr", "Expression.").Required().String()
	parseFile    = parseCmd.Flag("file", "Parse file.").Short('f').Bool()
)

var parseMain = register("parse", func() {
	f := parser.ParseString
	if *parseFile {
		f = parser.ParseFile
	}
	pr, err := f(*parseExprArg)
	e.Exit(err)
	fmt.Println(pr.LispResult())
})
