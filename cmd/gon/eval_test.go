package main

import (
	"testing"

	"github.com/alecthomas/assert"
	"github.com/orivej/go-nix/nix/eval"
	"github.com/orivej/go-nix/nix/parser"
)

func TestEval(t *testing.T) {
	for _, test := range [][2]string{
		{`{ "abc" = 1; }`, `{ abc = ...; }`},
		{`{ abc = 1; }`, `{ abc = ...; }`},
		{`{ a = 1; }.a`, `1`},
		{`{ a."b".${"c"} = 1; }."${"a"}".${"b"}.c`, `1`},
	} {
		pr, err := parser.ParseString(test[0])
		assert.NoError(t, err)
		s := eval.ValueString(eval.ParseResult(pr))
		assert.Equal(t, test[1], s)
	}
}
