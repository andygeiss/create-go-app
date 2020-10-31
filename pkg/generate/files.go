package generate

import (
	"bytes"
	"html/template"
	"io/ioutil"
	"regexp"
	"strings"
	"unicode"
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

var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")
var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchLink = regexp.MustCompile("(^[A-Za-z])|_([A-Za-z])")

func toSnakeCase(in string) (out string) {
	out = matchFirstCap.ReplaceAllString(in, "${1}_${2}")
	out = matchAllCap.ReplaceAllString(out, "${1}_${2}")
	return strings.ToLower(out)
}

func toCamelCase(in string) string {
	runes := []rune(in)
	runes[0] = unicode.ToLower(runes[0])
	in = string(runes)
	return matchLink.ReplaceAllStringFunc(in, func(s string) string {
		return strings.ToUpper(strings.Replace(s, "_", "", -1))
	})
}

func writeTemplateToFile(src, dst string, data interface{}) error {
	tmpl, err := template.New("t").Funcs(map[string]interface{}{
		"sc": func(in string) string {
			return toSnakeCase(in)
		},
		"lc": func(in string) string {
			return toCamelCase(toSnakeCase(in))
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
