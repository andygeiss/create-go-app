package path

import "os"

// HasFolder ...
func HasFolder(name string) bool {
	stat, _ := os.Stat(name)
	dirExists := false
	if stat != nil {
		dirExists = stat.IsDir()
	}
	return dirExists
}
