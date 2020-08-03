package parse_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/andygeiss/create-go-app/pkg/assert"
	"github.com/andygeiss/create-go-app/pkg/parse"
)

func TestParseOneService(t *testing.T) {
	setFilename(filepath.Join("testdata", "parse_one_service.go"))
	p := parse.NewParser()
	err := p.Parse()
	assert.That("error should be nil", t, err, nil)
	assert.That("parse should return one service", t, len(p.Services), 1)
}

func TestParseTwoServices(t *testing.T) {
	setFilename(filepath.Join("testdata", "parse_two_services.go"))
	p := parse.NewParser()
	err := p.Parse()
	assert.That("error should be nil", t, err, nil)
	assert.That("parse should return two services", t, len(p.Services), 2)
}

func setFilename(filename string) error {
	os.Setenv("GOFILE", filename)
	return nil
}
