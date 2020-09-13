package craft

import (
	"os"
	"path/filepath"

	"github.com/andygeiss/create-go-app/craft/templates"
	"github.com/andygeiss/create-go-app/pkg/generate"
	"github.com/andygeiss/create-go-app/pkg/merge"
)

// Bundle ...
type Bundle struct {
	Build     string   `json:"build"`
	Generator string   `json:"generator"`
	Name      string   `json:"name"`
	Path      string   `json:"path"`
	Services  []string `json:"services"`
	Version   string   `json:"version"`
}

// Craft ...
func (b *Bundle) Craft() error { // Add files.
	// Set basedir.
	baseDir := filepath.Join(os.Getenv("GOPATH"), "src", b.Path, b.Name)
	// Add files.
	files := map[string]string{
		filepath.Join(baseDir, "web", "src", "api_client.js"):    templates.APIClient,
		filepath.Join(baseDir, "web", "src", "app.js"):           templates.BundleAppJs,
		filepath.Join(baseDir, "web", "src", "app.scss"):         templates.BundleAppScss,
		filepath.Join(baseDir, "web", "src", "flat-element.js"):  templates.BundleFlatElementJs,
		filepath.Join(baseDir, "web", "src", "flat-mixins.scss"): templates.BundleFlatMixinsScss,
		filepath.Join(baseDir, "web", "src", "flat-reset.scss"):  templates.BundleFlatResetScss,
		filepath.Join(baseDir, "web", "src", "index.html"):       templates.BundleIndexHTML,
		filepath.Join(baseDir, "web", "static", "api.http"):      templates.APIHttp,
		filepath.Join(baseDir, "web", "static", ".gitkeep"):      templates.Gitkeep,
	}
	if err := generate.FilesByData(files, b); err != nil {
		return err
	}
	// Merge JavaScript files into one file named bundle.js.
	if err := merge.Files(
		filepath.Join(baseDir, "web", "static", "bundle.js"),
		filepath.Join(baseDir, "web", "src", "flat-element.js"),
		filepath.Join(baseDir, "web", "src", "app_client.js"),
		filepath.Join(baseDir, "web", "src", "app.js"),
	); err != nil {
		return err
	}
	// Merge JavaScript files into one file named bundle.js.
	if err := merge.Files(
		filepath.Join(baseDir, "web", "static", "bundle.scss"),
		filepath.Join(baseDir, "web", "src", "app.scss"),
	); err != nil {
		return err
	}
	return nil
}

// NewBundle ...
func NewBundle(build, generator, name, path, version string) *Bundle {
	return &Bundle{
		Build:     build,
		Generator: generator,
		Name:      name,
		Path:      path,
		Version:   version,
	}
}
