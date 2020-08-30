package craft_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/andygeiss/create-go-app/craft"
	"github.com/andygeiss/create-go-app/pkg/assert"
	"github.com/andygeiss/create-go-app/pkg/path"
)

func TestLibCraft(t *testing.T) {
	os.RemoveAll("foo")
	prj := craft.NewProject("b", "g", "foo", "v")
	err := prj.Craft()
	assert.That("err should be nil", t, err, nil)
	assert.That("folder [foo] exists", t, path.HasFolder(filepath.Join("foo")), true)
	assert.That("folder [foo/.git] exists", t, path.HasFolder(filepath.Join("foo", ".git")), true)
	assert.That("file   [foo/go.mod] exists", t, path.HasFile(filepath.Join("foo", "go.mod")), true)
	assert.That("file   [foo/Makefile] exists", t, path.HasFile(filepath.Join("foo", "Makefile")), true)
	assert.That("file   [foo/pkg/assert/assert.go] exists", t, path.HasFile(filepath.Join("foo", "pkg", "assert", "assert.go")), true)
	os.RemoveAll("foo")
}
