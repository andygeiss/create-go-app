package craft

import (
	"os"
	"os/exec"
	"path/filepath"

	"github.com/andygeiss/create-go-app/craft/templates"
	"github.com/andygeiss/create-go-app/pkg/generate"
)

// Bin ...
type Bin struct {
	Build      string `json:"build"`
	Generator  string `json:"generator"`
	GitBuild   string `json:"git_build"`
	GitVersion string `json:"git_version"`
	Name       string `json:"name"`
	Path       string `json:"path"`
	Version    string `json:"version"`
}

// Craft ...
func (b *Bin) Craft() error {
	// Create project dir.
	os.MkdirAll(filepath.Join(b.Name), 0755)
	// Set the Git version
	b.GitBuild = "initial"
	b.GitVersion = "v0.1.0"
	// Enable Go modules.
	wd, _ := os.Getwd()
	os.Chdir(b.Name)
	exec.Command("go", "mod", "init").Run()
	os.Chdir(wd)
	b.Path = readProjectPathFromModFile(b.Name)
	// Add files.
	files := map[string]string{
		filepath.Join(b.Name, "main.go"):                               templates.MainGo4App,
		filepath.Join(b.Name, "make.bat"):                              templates.MakeBat,
		filepath.Join(b.Name, "Makefile"):                              templates.Makefile,
		filepath.Join(b.Name, "build", "Dockerfile"):                   templates.Dockerfile,
		filepath.Join(b.Name, "internal", "status", "service.go"):      templates.ServiceGo,
		filepath.Join(b.Name, "internal", "status", "service_test.go"): templates.ServiceTestGo,
		filepath.Join(b.Name, "pkg", "api", "api.go"):                  templates.APIGo,
		filepath.Join(b.Name, "pkg", "assert", "assert.go"):            templates.AssertGo,
		filepath.Join(b.Name, "pkg", "event", "bus.go"):                templates.BusGo,
		filepath.Join(b.Name, "pkg", "event", "bus_test.go"):           templates.BusTest,
		filepath.Join(b.Name, "pkg", "server", "handlers.go"):          templates.HandlersGo,
		filepath.Join(b.Name, "pkg", "server", "middleware.go"):        templates.MiddlewareGo,
		filepath.Join(b.Name, "pkg", "server", "routes.go"):            templates.RoutesGo,
		filepath.Join(b.Name, "pkg", "server", "server.go"):            templates.ServerGo,
		filepath.Join(b.Name, "web", "src", "app.js"):                  templates.BundleAppJs,
		filepath.Join(b.Name, "web", "src", "app.css"):                 templates.BundleAppCSS,
		filepath.Join(b.Name, "web", "src", "index.html"):              templates.BundleIndexHTML,
	}
	if err := generate.FilesByData(files, b); err != nil {
		return err
	}
	// Format the files.
	exec.Command("gofmt", "-w", b.Name).Run()
	// Init Git repository.
	os.Chdir(b.Name)
	exec.Command("git", "init").Run()
	exec.Command("git", "add", ".").Run()
	exec.Command("git", "commit", "-m", "'initial commit'").Run()
	exec.Command("git", "tag", "-a", "v0.1.0", "-m", "'initial version'").Run()
	os.Chdir(wd)
	return nil
}

// NewBin ...
func NewBin(build, generator, name, version string) *Bin {
	return &Bin{
		Build:     build,
		Generator: generator,
		Name:      name,
		Version:   version,
	}
}
