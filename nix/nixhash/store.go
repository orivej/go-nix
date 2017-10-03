package nixhash

import (
	"path"
	"regexp"

	"github.com/orivej/go-nix/internal"
)

const storeDir = "/nix/store"

var rxStoreName = regexp.MustCompile(`^[A-Za-z0-9+._?=-]+$`)

// StorePath returns nix store path of the pathname.
func StorePath(pathname string) string {
	name := path.Base(pathname)
	checkStoreName(name)
	h := Path(pathname)
	return fixedOutputPath(true, h, name)
}

func fixedOutputPath(recursive bool, contentHash Hash, name string) string {
	if recursive {
		return makeStorePath("source", contentHash, name)
	}
	h := String("fixed:out" + contentHash.TypeString(16) + ":")
	return makeStorePath("output:out", h, name)
}

func makeStorePath(pathType string, h Hash, name string) string {
	s := pathType + ":" + h.TypeString(16) + ":" + storeDir + ":" + name
	return storeDir + "/" + String(s).Compress(20).String(32) + "-" + name
}

func checkStoreName(name string) {
	if !rxStoreName.MatchString(name) || name[0] == '.' {
		internal.Panicf("illegal name: %q", name)
	}
}
