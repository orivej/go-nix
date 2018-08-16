package main

import (
	"fmt"
	"os"

	"github.com/alecthomas/kingpin"
	"github.com/orivej/go-nix/internal"
	"github.com/pkg/profile"
)

var (
	actions = map[string]func(){}
	prof    = kingpin.Flag("profile", "Profile performance.").Bool()
)

func register(name string, f func()) func() {
	actions[name] = f
	return f
}

func main() {
	kingpin.HelpFlag.Short('h')
	action := kingpin.Parse()
	if *prof {
		defer profile.Start(profile.ProfilePath(".")).Stop()
	}
	defer handlePanic()
	actions[action]()
}

func handlePanic() {
	if v := recover(); v != nil {
		if e, ok := v.(internal.Error); ok {
			fmt.Fprintln(os.Stderr, "error:", e)
		} else {
			panic(v)
		}
	}
}
