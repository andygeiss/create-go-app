package craft

import (
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/andygeiss/create-go-app/craft/templates"
	"github.com/andygeiss/create-go-app/pkg/generate"
)

// Project ...
type Project struct {
	Build     string `json:"build"`
	Generator string `json:"generator"`
	Name      string `json:"name"`
	Path      string `json:"path"`
	Version   string `json:"version"`
}

// Craft ...
func (p *Project) Craft() error {
	// Create project dir.
	os.MkdirAll(filepath.Join(p.Name), 0755)
	// Enable Go modules.
	wd, _ := os.Getwd()
	os.Chdir(p.Name)
	exec.Command("go", "mod", "init").Run()
	os.Chdir(wd)
	p.Path = readProjectPathFromModFile(p.Name)
	// Add files.
	files := map[string]string{
		filepath.Join(p.Name, "main.go"):                               templates.MainGo,
		filepath.Join(p.Name, "Makefile"):                              templates.Makefile,
		filepath.Join(p.Name, "internal", "status", "service.go"):      templates.ServiceGo,
		filepath.Join(p.Name, "internal", "status", "service_test.go"): templates.ServiceTestGo,
		filepath.Join(p.Name, "pkg", "api", "api.http"):                templates.APIHttp,
		filepath.Join(p.Name, "pkg", "api", "api.go"):                  templates.APIGo,
		filepath.Join(p.Name, "pkg", "assert", "assert.go"):            templates.AssertGo,
		filepath.Join(p.Name, "pkg", "server", "handlers.go"):          templates.HandlersGo,
		filepath.Join(p.Name, "pkg", "server", "middleware.go"):        templates.MiddlewareGo,
		filepath.Join(p.Name, "pkg", "server", "routes.go"):            templates.RoutesGo,
		filepath.Join(p.Name, "pkg", "server", "server.go"):            templates.ServerGo,
	}
	if err := generate.FilesByData(files, p); err != nil {
		return err
	}
	// Format the files.
	exec.Command("gofmt", "-w", p.Name).Run()
	// Init Git repository.
	os.Chdir(p.Name)
	exec.Command("git", "init").Run()
	exec.Command("git", "add", ".").Run()
	exec.Command("git", "commit", "-m", "'initial commit'").Run()
	exec.Command("git", "tag", "-a", "v0.1.0", "-m", "'initial version'").Run()
	os.Chdir(wd)
	return nil
}

// NewProject ...
func NewProject(build, generator, name, version string) *Project {
	return &Project{
		Build:     build,
		Generator: generator,
		Name:      name,
		Version:   version,
	}
}

func readContentFromFile(filename string) string {
	content, _ := ioutil.ReadFile(filename)
	return string(content)
}
