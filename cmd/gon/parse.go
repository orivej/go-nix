package main

import (
	"fmt"

	"github.com/orivej/e"
	"github.com/orivej/go-nix/nix/parser"
)

var (
	parseExprArg = parseCmd.Arg("expr", "Expression.").Required().String()
)

var parseCmd, parseMain = register("parse", "Parse Nix expression.", func() {
	pr, err := parser.ParseString(*parseExprArg)
	e.Exit(err)
	fmt.Println(pr.LispResult())
})
