package main

import (
	"fmt"

	"github.com/alecthomas/kingpin"
	"github.com/orivej/e"
	"github.com/orivej/go-nix/nix/eval"
	"github.com/orivej/go-nix/nix/parser"
)

var (
	evalCmd     = kingpin.Command("eval", "Eval Nix expression.")
	evalExprArg = evalCmd.Arg("expr", "Expression.").Required().String()
)

var evalMain = register("eval", func() {
	pr, err := parser.ParseString(*evalExprArg)
	e.Exit(err)
	fmt.Println(eval.ValueString(eval.ParseResult(pr)))
})
