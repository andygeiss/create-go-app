package generate_test

import (
	"path/filepath"
	"testing"

	"github.com/andygeiss/create-go-app/pkg/assert"
	"github.com/andygeiss/create-go-app/pkg/generate"
	"github.com/andygeiss/create-go-app/pkg/path"
)

func TestFilesOneFile(t *testing.T) {
	files := map[string]string{
		filepath.Join("testdata", "my.go.tmpl"): filepath.Join("testdata", "my.go"),
	}
	generate.FilesByData(files, nil)
	assert.That("file [my.go] exists", t, path.HasFile(filepath.Join("testdata", "my.go")), true)
}
