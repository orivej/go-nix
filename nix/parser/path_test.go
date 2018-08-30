package parser

import (
	"testing"

	"github.com/alecthomas/assert"
)

var allPackages = ResolvePath("nixpkgs/pkgs/top-level/all-packages.nix")
var nixpkgs = ResolvePath("nixpkgs")
var attrsets = ResolvePath("nixpkgs/lib/attrsets.nix")

func skipWithoutNixPath(t *testing.T) {
	if len(NixPath) == 0 {
		t.Skip("NIX_PATH is not set")
	}
}

func TestSplitNixPath(t *testing.T) {
	tests := map[string][][2]string{
		"":          nil,
		"a":         {{"", "a"}},
		"a:b:c":     {{"", "a"}, {"", "b"}, {"", "c"}},
		"p=b":       {{"p", "b"}},
		"a:p=b:c":   {{"", "a"}, {"p", "b"}, {"", "c"}},
		"a:p=b:q=c": {{"", "a"}, {"p", "b"}, {"q", "c"}},
	}
	for q, a := range tests {
		assert.Equal(t, a, splitNixPath(q), q)
	}
}

func TestResolvePath(t *testing.T) {
	no := func(string) bool { return false }
	yes := func(string) bool { return true }
	assert.Equal(t, "", resolvePath("f", [][2]string{{"", "/a"}}, no))
	assert.Equal(t, "/a/f", resolvePath("f", [][2]string{{"", "/a"}}, yes))
	assert.Equal(t, "/b/f", resolvePath("p/f", [][2]string{{"p", "/b"}}, yes))
	assert.Equal(t, "", resolvePath("pa/f", [][2]string{{"p", "/b"}}, yes))
}
