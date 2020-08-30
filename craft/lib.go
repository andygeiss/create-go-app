package craft

import (
	"os"
	"os/exec"
	"path/filepath"

	"github.com/andygeiss/create-go-app/craft/templates"
	"github.com/andygeiss/create-go-app/pkg/generate"
)

// Lib ...
type Lib struct {
	Build     string `json:"build"`
	Generator string `json:"generator"`
	Name      string `json:"name"`
	Path      string `json:"path"`
	Version   string `json:"version"`
}

// Craft ...
func (b *Lib) Craft() error {
	// Create project dir.
	os.MkdirAll(filepath.Join(b.Name), 0755)
	// Enable Go modules.
	wd, _ := os.Getwd()
	os.Chdir(b.Name)
	exec.Command("go", "mod", "init").Run()
	os.Chdir(wd)
	b.Path = readProjectPathFromModFile(b.Name)
	// Add files.
	files := map[string]string{
		filepath.Join(b.Name, "pkg", "assert", "assert.go"): templates.AssertGo,
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

// NewLib ...
func NewLib(build, generator, name, version string) *Lib {
	return &Lib{
		Build:     build,
		Generator: generator,
		Name:      name,
		Version:   version,
	}
}
