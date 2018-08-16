package main

import (
	"fmt"

	"github.com/alecthomas/kingpin"
	"github.com/orivej/go-nix/nix/nixhash"
)

var (
	hashCmd  = kingpin.Command("hash", "Hash a path.")
	hashPath = hashCmd.Arg("path", "path").Required().String()
)

var hashMain = register("hash", func() {
	fmt.Println(nixhash.StorePath(*hashPath))
})
