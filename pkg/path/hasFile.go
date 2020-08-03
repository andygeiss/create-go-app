package path

import "os"

// HasFile ...
func HasFile(filename string) bool {
	stat, _ := os.Stat(filename)
	fileExists := false
	if stat != nil {
		fileExists = !stat.IsDir()
	}
	return fileExists
}
