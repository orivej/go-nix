# Overview

This repository contains:

- `nix/parser` - a Nix parser. Optimized for speed, it parses all Nixpkgs in 2 seconds. It preserves comments and source positions and can be used to implement Nix files formatting.
- `nix/nixhash` - Nix-compatible hasher for store paths.
- `nix/eval` - an incomplete Nix evaluator. It can't evaluate realistic Nix files, but it's a start.
- `cmd/gon` - an utility that exposes these libraries from the command line.
