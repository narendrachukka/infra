package server

import (
	"net/http"
	"testing"

	"github.com/psanford/memfs"
	"gotest.tools/v3/assert"
)

func TestStaticFileSystem(t *testing.T) {
	fs := memfs.New()
	err := fs.MkdirAll("ui", 0755)
	assert.NilError(t, err)
	err = fs.WriteFile("ui/dashboard.html", nil, 0644)
	assert.NilError(t, err)

	sfs := &StaticFileSystem{
		base: http.FS(fs),
	}

	t.Run("open file with .html suffix", func(t *testing.T) {
		f, err := sfs.Open("dashboard.html")
		assert.NilError(t, err)

		stat, err := f.Stat()
		assert.NilError(t, err)
		assert.Equal(t, stat.Name(), "dashboard.html")
	})

	t.Run("open file without suffix", func(t *testing.T) {
		f, err := sfs.Open("dashboard")
		assert.NilError(t, err)

		stat, err := f.Stat()
		assert.NilError(t, err)
		assert.Equal(t, stat.Name(), "dashboard.html")
	})

	t.Run("file exists with .html suffix", func(t *testing.T) {
		assert.Assert(t, sfs.Exists("/", "/dashboard.html"))
	})

	t.Run("file exists without suffix", func(t *testing.T) {
		assert.Assert(t, sfs.Exists("/", "/dashboard"))
	})
}
