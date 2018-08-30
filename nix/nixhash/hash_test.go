package nixhash

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/alecthomas/assert"
	"github.com/orivej/e"
)

func TestHash(t *testing.T) {
	dir, err := ioutil.TempDir("", "TestHash")
	e.Panic(err)
	defer func() { e.Panic(os.RemoveAll(dir)) }()

	file := filepath.Join(dir, "name")
	err = ioutil.WriteFile(file, []byte("text\n"), 0644)
	e.Panic(err)

	assert.Equal(t, "/nix/store/a68bcimav7hwazsfk1iiabv7fxyr3dh4-name",
		StorePath(file, ""))
}
