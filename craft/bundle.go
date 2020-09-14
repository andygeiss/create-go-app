package craft

import (
	"os"
	"os/exec"
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
	baseDir := filepath.Join(os.Getenv("GOPATH"), "src", b.Path)
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
		filepath.Join(baseDir, "web", "src", "api_client.js"),
		filepath.Join(baseDir, "web", "src", "app.js"),
	); err != nil {
		return err
	}
	// Compile the SASS.
	cmd := exec.Command("sassc",
		"-t", "compressed",
		filepath.Join(baseDir, "web", "src", "app.scss"),
		filepath.Join(baseDir, "web", "static", "bundle.css"),
	)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	if err := cmd.Run(); err != nil {
		return err
	}
	// Compress the CSS.
	exec.Command("postcss",
		"-u", "autoprefixer", "--autoprefixer.overrideBrowserslist", "'defaults, ie 10'",
		"-o", filepath.Join(baseDir, "web", "static", "bundle.min.css"),
		filepath.Join(baseDir, "web", "static", "bundle.css"),
	).Run()
	// Compress the JavaScript.
	exec.Command("java", "-jar",
		os.Getenv("HOME")+"/bin/closure-compiler.jar",
		"--compilation_level", "SIMPLE_OPTIMIZATIONS",
		"--language_out", "ECMASCRIPT_2015",
		"--js", filepath.Join(baseDir, "web", "static", "bundle.js"),
		"--js_output_file", filepath.Join(baseDir, "web", "static", "bundle.min.js"),
	).Run()
	// Cleanup.
	os.Remove(filepath.Join(baseDir, "web", "static", "bundle.css"))
	os.Remove(filepath.Join(baseDir, "web", "static", "bundle.js"))
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
