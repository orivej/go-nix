package main

import (
	"fmt"

	"github.com/orivej/go-nix/nix/nixhash"
)

var (
	hashPath = hashCmd.Arg("path", "path").Required().String()
)

var hashCmd, hashMain = register("hash", "Hash a path.", func() {
	fmt.Println(nixhash.StorePath(*hashPath))
})
