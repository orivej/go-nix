package parser

import (
	"testing"

	"github.com/alecthomas/assert"
	"github.com/orivej/e"
	"github.com/orivej/go-nix/nix/util"
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
	skipWithoutNixPath(t)
	cnt := 0
	err := util.WalkNix(nixpkgs, func(path string) error {
		_, err := ParseFile(path)
		assert.NoError(t, err, path)
		cnt++
		return nil
	})
	t.Logf("parsed %d files", cnt)
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
