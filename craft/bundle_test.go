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
	os.RemoveAll("testdata")
	c := craft.NewBundle("b", "g", "foo", "github.com/andygeiss/create-go-app/craft/testdata/foo", "v")
	err := c.Craft()
	os.Chdir("testdata")
	assert.That("err should be nil", t, err, nil)
	assert.That("file   [foo/web/src/flat-element.js] exists", t, path.HasFile(filepath.Join("foo", "web", "src", "flat-element.js")), true)
	assert.That("file   [foo/web/src/flat-mixins.scss] exists", t, path.HasFile(filepath.Join("foo", "web", "src", "flat-mixins.scss")), true)
	assert.That("file   [foo/web/src/flat-reset.scss] exists", t, path.HasFile(filepath.Join("foo", "web", "src", "flat-reset.scss")), true)
	assert.That("file   [foo/web/static/api.http] exists", t, path.HasFile(filepath.Join("foo", "web", "static", "api.http")), true)
	assert.That("file   [foo/web/static/bundle.min.css] exists", t, path.HasFile(filepath.Join("foo", "web", "static", "bundle.min.css")), true)
	assert.That("file   [foo/web/static/bundle.min.js] exists", t, path.HasFile(filepath.Join("foo", "web", "static", "bundle.min.js")), true)
	assert.That("file   [foo/web/static/bundle.css] not exists", t, path.HasFile(filepath.Join("foo", "web", "static", "bundle.css")), false)
	assert.That("file   [foo/web/static/bundle.js] not exists", t, path.HasFile(filepath.Join("foo", "web", "static", "bundle.js")), false)
	assert.That("file   [foo/web/static/index.html] exists", t, path.HasFile(filepath.Join("foo", "web", "static", "index.html")), true)
	os.RemoveAll("foo")
}
