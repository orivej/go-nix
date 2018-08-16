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
)

var parseMain = register("parse", func() {
	pr, err := parser.ParseString(*parseExprArg)
	e.Exit(err)
	fmt.Println(pr.LispResult())
})
