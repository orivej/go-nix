[![Build Status](https://travis-ci.com/orivej/go-nix.svg?branch=master)](https://travis-ci.com/orivej/go-nix)

# Overview

This repository contains:

- `nix/parser` - a Nix parser. Optimized for speed, it parses all Nixpkgs in 2 seconds. It preserves comments and source positions and can be used to implement Nix files formatting.
- `nix/nixhash` - Nix-compatible hasher for store paths.
- `nix/eval` - an incomplete Nix evaluator. It can't evaluate realistic Nix files, but it's a start.
- `cmd/gon` - an utility that exposes these libraries from the command line.

# Development

## How to build

```sh
$ nix-build
```

## Build cmd/gon

```sh
$ result-bin/bin/gon help
```

## Get a development environment

```sh
$ nix-shell
$ go build ./...
```

## Run tests

```sh
$ go test ./...
```

# Credits

- [ragel](https://www.colm.net/open-source/ragel/) generates the Nix lexer
- [goyacc](https://godoc.org/golang.org/x/tools/cmd/goyacc) generates the Nix parser
- [kingpin](https://github.com/alecthomas/kingpin) powers the CLI
