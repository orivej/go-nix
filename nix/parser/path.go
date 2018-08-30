package parser

import (
	"os"
	gopath "path"
	"strings"
)

var NixPath = splitNixPath(os.Getenv("NIX_PATH"))

func ResolvePath(s string) string {
	return resolvePath(s, NixPath, fsPathExists)
}

func splitNixPath(s string) [][2]string {
	if s == "" {
		return nil
	}
	parts := strings.Split(s, ":")
	entries := make([][2]string, len(parts))
	for i, part := range parts {
		entries[i][1] = part
		if strings.Contains(part, "=") {
			copy(entries[i][:], strings.SplitN(part, "=", 2))
		}
	}
	return entries
}

func fsPathExists(s string) bool {
	_, err := os.Stat(s)
	return err == nil
}

func resolvePath(s string, nixpath [][2]string, exists func(string) bool) string {
	for _, pair := range nixpath {
		k, v := pair[0], pair[1]
		var r string
		switch {
		case k == "":
			r = gopath.Join(v, s)
		case s == k || strings.HasPrefix(s, k+"/"):
			r = v + s[len(k):]
		default:
			continue
		}
		if exists(r) {
			return r
		}
	}
	return ""
}
