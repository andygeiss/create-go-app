package craft

import (
	"io/ioutil"
	"path/filepath"
	"strings"
)

func readProjectPathFromModFile(name string) string {
	data, _ := ioutil.ReadFile(filepath.Join(name, "go.mod"))
	lines := strings.Split(string(data), "\n")
	parts := strings.Split(lines[0], " ")
	return parts[1]
}
