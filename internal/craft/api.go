package craft

import (
	"bytes"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/andygeiss/create-go-app/internal/craft/templates"
	"github.com/andygeiss/create-go-app/pkg/generate"
	"github.com/andygeiss/create-go-app/pkg/parse"
)

// API ...
type API struct {
	Build      string   `json:"build"`
	Generator  string   `json:"generator"`
	GitBuild   string   `json:"git_build"`
	GitVersion string   `json:"git_version"`
	Name       string   `json:"name"`
	Path       string   `json:"path"`
	Services   []string `json:"services"`
	Version    string   `json:"version"`
}

// Craft ...
func (a *API) Craft() error {
	// Get the Git build
	buf := bytes.Buffer{}
	cmd := exec.Command("git", "rev-parse", "--short", "HEAD")
	cmd.Stdout = &buf
	cmd.Run()
	a.GitBuild = buf.String()
	a.GitBuild = strings.ReplaceAll(a.GitBuild, "\r", "")
	a.GitBuild = strings.ReplaceAll(a.GitBuild, "\n", "")
	// Get the Git version
	buf = bytes.Buffer{}
	cmd = exec.Command("git", "describe", "--tags")
	cmd.Stdout = &buf
	cmd.Run()
	a.GitVersion = buf.String()
	a.GitVersion = strings.ReplaceAll(a.GitVersion, "\r", "")
	a.GitVersion = strings.ReplaceAll(a.GitVersion, "\n", "")
	// Parse the API and get the services.
	parser := parse.NewParser()
	if err := parser.Parse(); err != nil {
		return err
	}
	a.Services = parser.Services
	// Update files.
	baseDir := filepath.Join(os.Getenv("GOPATH"), "src", a.Path)
	files := map[string]string{
		filepath.Join(baseDir, "make.bat"):                     templates.MakeBat,
		filepath.Join(baseDir, "build", "Dockerfile"):          templates.Dockerfile,
		filepath.Join(baseDir, "pkg", "server", "handlers.go"): templates.Handlers2Go,
		filepath.Join(baseDir, "pkg", "server", "routes.go"):   templates.Routes2Go,
		filepath.Join(baseDir, "pkg", "server", "server.go"):   templates.Server2Go,
		filepath.Join(baseDir, "web", "src", "api_client.js"):  templates.APIClient,
		filepath.Join(baseDir, "web", "static", "api.http"):    templates.APIHttp,
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
