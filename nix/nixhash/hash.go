// Package nixhash implements nix derivation hashing.
package nixhash

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"io"
	"os"

	"github.com/orivej/e"
)

var narVersionMagic1 = []byte("nix-archive-1")

type Hash []byte

// String sha256.
func String(s string) Hash {
	b := sha256.Sum256([]byte(s))
	return Hash(b[:])
}

// File sha256.
func File(path string) Hash {
	f, err := os.Open(path)
	e.Exit(err)
	defer e.CloseOrExit(f)

	h := sha256.New()
	_, err = io.Copy(h, f)
	e.Exit(err)

	return Hash(h.Sum(nil))
}

// Path hash.
func Path(path string) Hash {
	h := sha256.New()
	hs := NewSink(h)
	_, err := hs.Write(narVersionMagic1)
	e.Exit(err)
	dump(path, hs)

	return Hash(h.Sum(nil))
}

// Compress the hash down to size bytes.
func (h Hash) Compress(size int) Hash {
	r := make([]byte, size)
	for i, b := range h {
		r[i%size] ^= b
	}
	return Hash(r)
}

// String encodes the hash to a string in a given base (16, 32 or 64).
func (h Hash) String(base int) string {
	switch base {
	case 16:
		return hex.EncodeToString(h)
	case 32:
		return string(base32Encode(h))
	case 64:
		return base64.StdEncoding.EncodeToString(h)
	default:
		panic("unsupported base")
	}
}

// TypeString prepends "sha256:" to String.
func (h Hash) TypeString(base int) string {
	return "sha256:" + h.String(base)
}
