package generate

import (
	"os"
	"path/filepath"
)

// Folders ...
func Folders(folders ...string) error {
	for _, folder := range folders {
		if err := os.MkdirAll(folder, 0755); err != nil {
			return err
		}
	}
	return nil
}

// FoldersByFile ...
func FoldersByFile(files ...string) error {
	for _, file := range files {
		path := filepath.Dir(file)
		if err := os.MkdirAll(path, 0755); err != nil {
			return err
		}
	}
	return nil
}
