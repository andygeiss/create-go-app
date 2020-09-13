package merge

import (
	"bytes"
	"io/ioutil"
)

// Files ...
func Files(into string, from ...string) error {
	var buf bytes.Buffer
	for _, file := range from {
		b, err := ioutil.ReadFile(file)
		if err != nil {
			return err
		}
		buf.Write(b)
	}
	err := ioutil.WriteFile(into, buf.Bytes(), 0644)
	if err != nil {
		return err
	}
	return nil
}
