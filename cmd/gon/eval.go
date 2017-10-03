package main

import (
	"fmt"

	"github.com/orivej/e"
	"github.com/orivej/go-nix/nix/eval"
	"github.com/orivej/go-nix/nix/parser"
)

var (
	evalExprArg = evalCmd.Arg("expr", "Expression.").Required().String()
)

var evalCmd, evalMain = register("eval", "Eval Nix expression.", func() {
	pr, err := parser.ParseString(*evalExprArg)
	e.Exit(err)
	fmt.Println(eval.ValueString(eval.ParseResult(pr)))
})
