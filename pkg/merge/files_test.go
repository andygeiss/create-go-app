package merge_test

import (
	"path/filepath"
	"testing"

	"github.com/andygeiss/create-go-app/pkg/assert"
	"github.com/andygeiss/create-go-app/pkg/merge"
	"github.com/andygeiss/create-go-app/pkg/path"
)

func TestMerge(t *testing.T) {
	merge.Files(
		filepath.Join("testdata", "c.txt"),
		filepath.Join("testdata", "a.txt"),
		filepath.Join("testdata", "b.txt"),
	)
	assert.That("file [testdata/c.txt] exists", t, path.HasFile(filepath.Join("testdata", "c.txt")), true)
}
