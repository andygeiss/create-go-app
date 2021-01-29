package craft

import (
	"os"
	"path/filepath"

	"github.com/andygeiss/create-go-app/internal/craft/templates"
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
	baseDir := filepath.Join(os.Getenv("GOPATH"), "src", b.Path)
	// Add files.
	files := map[string]string{
		filepath.Join(baseDir, "web", "src", "component.js"): templates.BundleComponentJs,
	}
	if err := generate.FilesByData(files, b); err != nil {
		return err
	}
	// Merge JavaScript files into one file named bundle.js.
	if err := merge.Files(
		filepath.Join(baseDir, "web", "static", "bundle.js"),
		filepath.Join(baseDir, "web", "src", "component.js"),
		filepath.Join(baseDir, "web", "src", "api_client.js"),
		filepath.Join(baseDir, "web", "src", "app.js"),
	); err != nil {
		return err
	}
	// Copy ServiceWorker ...
	if err := merge.Files(
		filepath.Join(baseDir, "web", "static", "service_worker.js"),
		filepath.Join(baseDir, "web", "src", "service_worker.js"),
	); err != nil {
		return err
	}
	// Merge Stylesheet files into one file named bundle.css.
	if err := merge.Files(
		filepath.Join(baseDir, "web", "static", "bundle.css"),
		filepath.Join(baseDir, "web", "src", "app.css"),
	); err != nil {
		return err
	}
	// Copy index.html
	merge.Files(
		filepath.Join(baseDir, "web", "static", "index.html"),
		filepath.Join(baseDir, "web", "src", "index.html"),
	)
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
