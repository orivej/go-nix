package parser

import (
	"io/ioutil"
	"sort"
	"testing"

	"github.com/alecthomas/assert"
	"github.com/orivej/e"
	"github.com/orivej/go-nix/nix/util"
)

func TestLexOne(t *testing.T) {
	skipWithoutNixPath(t)
	path := attrsets
	data, err := ioutil.ReadFile(path)
	e.Exit(err)
	r, err := lex(data, path)
	t.Log(r.tokens)
	t.Log(r.comments)
	t.Log(r.Last())
	assert.NoError(t, err)
}

func TestLexAll(t *testing.T) {
	if testing.Short() {
		t.SkipNow()
	}
	skipWithoutNixPath(t)
	hist := map[int]int{}
	err := util.WalkNix(nixpkgs, func(path string) error {
		r, err := lexFile(path)
		assert.NoError(t, err, path)
		for _, tok := range r.tokens {
			hist[tok.sym]++
		}
		for _, tok := range r.comments {
			hist[tok.sym]++
		}
		return nil
	})
	assert.NoError(t, err)
	printSymHist(t, hist)
}

func BenchmarkLex(b *testing.B) {
	path := allPackages
	data, err := ioutil.ReadFile(path)
	e.Exit(err)

	for i := 0; i < b.N; i++ {
		_, err = lex(data, path)
		assert.NoError(b, err)
	}
}

func printSymHist(t *testing.T, hist map[int]int) {
	t.Helper()
	keys := make([]int, 0, len(hist))
	for k := range hist {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool { return hist[keys[i]] < hist[keys[j]] })
	for _, k := range keys {
		t.Log(hist[k], symString(k))
	}
}
