package craft

import (
	"path/filepath"

	"github.com/andygeiss/create-go-app/craft/templates"
	"github.com/andygeiss/create-go-app/pkg/generate"
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
	files := map[string]string{
		filepath.Join(b.Name, "web", "src", "api_client.js"):    templates.APIClient,
		filepath.Join(b.Name, "web", "src", "app.js"):           templates.BundleAppJs,
		filepath.Join(b.Name, "web", "src", "app.scss"):         templates.BundleAppScss,
		filepath.Join(b.Name, "web", "src", "flat-element.js"):  templates.BundleFlatElementJs,
		filepath.Join(b.Name, "web", "src", "flat-mixins.scss"): templates.BundleFlatMixinsScss,
		filepath.Join(b.Name, "web", "src", "flat-reset.scss"):  templates.BundleFlatResetScss,
		filepath.Join(b.Name, "web", "src", "index.html"):       templates.BundleIndexHTML,
		filepath.Join(b.Name, "web", "static", "api.http"):      templates.APIHttp,
		filepath.Join(b.Name, "web", "static", ".gitkeep"):      templates.Gitkeep,
	}
	if err := generate.FilesByData(files, b); err != nil {
		return err
	}
	return nil
}

// NewBundle ...
func NewBundle(build, generator, name, version string) *Bundle {
	return &Bundle{
		Build:     build,
		Generator: generator,
		Name:      name,
		Version:   version,
	}
}
