package craft

import (
	"os"
	"os/exec"
	"path/filepath"

	"github.com/andygeiss/create-go-app/craft/templates"
	"github.com/andygeiss/create-go-app/pkg/generate"
	"github.com/andygeiss/create-go-app/pkg/parse"
)

// API ...
type API struct {
	Build     string   `json:"build"`
	Generator string   `json:"generator"`
	Name      string   `json:"name"`
	Path      string   `json:"path"`
	Services  []string `json:"services"`
	Version   string   `json:"version"`
}

// Craft ...
func (a *API) Craft() error {
	// Parse the API and get the services.
	parser := parse.NewParser()
	if err := parser.Parse(); err != nil {
		return err
	}
	a.Services = parser.Services
	// Update files.
	baseDir := filepath.Join(os.Getenv("GOPATH"), "src", a.Path)
	files := map[string]string{
		filepath.Join(baseDir, "pkg", "server", "handlers.go"): templates.Handlers2Go,
		filepath.Join(baseDir, "pkg", "server", "routes.go"):   templates.Routes2Go,
		filepath.Join(baseDir, "pkg", "server", "server.go"):   templates.Server2Go,
	}
	if err := generate.FilesByData(files, a); err != nil {
		return err
	}
	// Format the files.
	exec.Command("gofmt", "-w", baseDir).Run()
	return nil
}

// NewAPI ...
func NewAPI(build, generator, name, path, version string) *API {
	return &API{
		Build:     build,
		Generator: generator,
		Name:      name,
		Path:      path,
		Version:   version,
	}
}
