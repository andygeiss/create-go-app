package generate_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/andygeiss/create-go-app/pkg/assert"
	"github.com/andygeiss/create-go-app/pkg/generate"
	"github.com/andygeiss/create-go-app/pkg/path"
)

func TestFoldersOneFolder(t *testing.T) {
	os.RemoveAll("foo")
	err := generate.Folders("foo")
	assert.That("err should be nil", t, err, nil)
	assert.That("project folder exists", t, path.HasFolder("foo"), true)
}

func TestFoldersTwoFolders(t *testing.T) {
	err := generate.Folders("foo", filepath.Join("foo", "bar"))
	assert.That("err should be nil", t, err, nil)
	assert.That("folder [foo] exists", t, path.HasFolder("foo"), true)
	assert.That("folder [foo/bar] exists", t, path.HasFolder(filepath.Join("foo", "bar")), true)
	os.RemoveAll("foo")
}

func TestFoldersByFileOneFolder(t *testing.T) {
	err := generate.FoldersByFile(filepath.Join("foo", "bar.go"))
	assert.That("err should be nil", t, err, nil)
	assert.That("folder [foo] exists", t, path.HasFolder("foo"), true)
	os.RemoveAll("foo")
}

func TestFoldersByFileTwoFolders(t *testing.T) {
	err := generate.FoldersByFile(filepath.Join("foo", "bar.go"))
	assert.That("err should be nil", t, err, nil)
	assert.That("folder [foo] exists", t, path.HasFolder("foo"), true)
	assert.That("folder [foo/bar] does not exists", t, path.HasFolder(filepath.Join("foo", "bar.go")), false)
	os.RemoveAll("foo")
}
