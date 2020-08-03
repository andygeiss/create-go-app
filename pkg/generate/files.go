package generate

import (
	"bytes"
	"html/template"
	"io/ioutil"
	"strings"
)

// FilesByData ...
func FilesByData(files map[string]string, data interface{}) error {
	for file, src := range files {
		FoldersByFile(file)
		if err := writeTemplateToFile(src, file, data); err != nil {
			return err
		}
	}
	return nil
}

func writeTemplateToFile(src, dst string, data interface{}) error {
	tmpl, err := template.New("t").Funcs(map[string]interface{}{
		"lc": func(in string) string {
			return strings.ToLower(in)
		},
	}).Parse(src)
	if err != nil {
		return err
	}
	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return err
	}
	if err := ioutil.WriteFile(dst, buf.Bytes(), 0644); err != nil {
		return err
	}
	return nil
}
