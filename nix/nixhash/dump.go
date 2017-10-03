package nixhash

import (
	"io"
	"io/ioutil"
	"os"
	"path"
	"sort"

	"github.com/orivej/e"
	"github.com/orivej/go-nix/internal"
)

func dump(p string, sink Sink) {
	fi, err := os.Lstat(p)
	e.Exit(err)
	sink.S("(")
	switch {
	case fi.Mode().IsRegular():
		sink.S("type", "regular")
		if fi.Mode()&0100 != 0 {
			sink.S("executable", "")
		}
		dumpContents(p, fi.Size(), sink)
	case fi.IsDir():
		sink.S("type", "directory")
		for _, name := range readDirectory(p) {
			sink.S("entry", "(", "name", name, "node")
			dump(path.Join(p, name), sink)
			sink.S(")")
		}
	case fi.Mode()&os.ModeSymlink != 0:
		target, err := os.Readlink(p)
		e.Exit(err)
		sink.S("type", "symlink", "target", target)
	default:
		internal.Panicf("illegal file: %q", path.Join(p, fi.Name()))
	}
	sink.S(")")
}

func dumpContents(p string, size int64, sink Sink) {
	sink.S("contents")
	f, err := os.Open(p)
	e.Exit(err)
	e.Exit(sink.Begin(size))
	_, err = io.Copy(sink, f)
	e.Exit(err)
	e.Exit(sink.End())
	e.Exit(f.Close())
}

func readDirectory(p string) []string {
	fis, err := ioutil.ReadDir(p)
	internal.Check(err)
	names := make([]string, 0, len(fis))
	for _, fi := range fis {
		name := fi.Name()
		if name != "" {
			names = append(names, name)
		}
	}
	sort.Strings(names)
	return names
}
