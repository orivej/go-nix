package parser

import (
	"testing"

	"github.com/alecthomas/assert"
	"github.com/orivej/e"
)

func TestParseOne(t *testing.T) {
	p, err := ParseString(oneText)
	assert.NoError(t, err)
	t.Log(p.LispResult())
}

func TestParseAll(t *testing.T) {
	if testing.Short() {
		t.SkipNow()
	}
	err := walkNix(nixpkgs, func(path string) error {
		_, err := parseFile(path)
		assert.NoError(t, err, path)
		return nil
	})
	assert.NoError(t, err)
}

func BenchmarkParse(b *testing.B) {
	path := allPackages
	lr, err := lexFile(path)
	e.Exit(err)

	for i := 0; i < b.N; i++ {
		_, err = parse(lr)
		assert.NoError(b, err)
	}
}

var oneText = `
# Function comment.
{ lib }:
with lib.trivial;
rec {
  inherit lib;
  inherit (builtins) head tail;
  s.a = "a${"b"}c";
}
`
