package parser

import (
	"os"
	"path/filepath"
)

func walkNix(dir string, f func(string) error) error {
	return filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() || filepath.Ext(info.Name()) != ".nix" {
			return nil
		}
		return f(path)
	})
}
