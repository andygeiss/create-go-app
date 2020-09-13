package craft_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/andygeiss/create-go-app/craft"
	"github.com/andygeiss/create-go-app/pkg/assert"
	"github.com/andygeiss/create-go-app/pkg/path"
)

func TestBundleCraft(t *testing.T) {
	os.RemoveAll("foo")
	c := craft.NewBundle("b", "g", "foo", "v")
	err := c.Craft()
	assert.That("err should be nil", t, err, nil)
	assert.That("folder [foo/web] exists", t, path.HasFolder(filepath.Join("foo", "web")), true)
	assert.That("folder [foo/web/src] exists", t, path.HasFolder(filepath.Join("foo", "web", "src")), true)
	assert.That("folder [foo/web/static] exists", t, path.HasFolder(filepath.Join("foo", "web", "static")), true)
	assert.That("file   [foo/web/src/api_client.js] exists", t, path.HasFile(filepath.Join("foo", "web", "src", "api_client.js")), true)
	assert.That("file   [foo/web/src/app.js] exists", t, path.HasFile(filepath.Join("foo", "web", "src", "app.js")), true)
	assert.That("file   [foo/web/src/app.scss] exists", t, path.HasFile(filepath.Join("foo", "web", "src", "app.scss")), true)
	assert.That("file   [foo/web/src/flat-element.js] exists", t, path.HasFile(filepath.Join("foo", "web", "src", "flat-element.js")), true)
	assert.That("file   [foo/web/src/flat-mixins.scss] exists", t, path.HasFile(filepath.Join("foo", "web", "src", "flat-mixins.scss")), true)
	assert.That("file   [foo/web/src/flat-reset.scss] exists", t, path.HasFile(filepath.Join("foo", "web", "src", "flat-reset.scss")), true)
	assert.That("file   [foo/web/src/index.html] exists", t, path.HasFile(filepath.Join("foo", "web", "src", "index.html")), true)
	assert.That("file   [foo/web/static/.gitkeep] exists", t, path.HasFile(filepath.Join("foo", "web", "static", ".gitkeep")), true)
	assert.That("file   [foo/web/static/api.http] exists", t, path.HasFile(filepath.Join("foo", "web", "static", "api.http")), true)
	os.RemoveAll("foo")
}
