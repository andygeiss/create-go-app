package templates

// ServiceTestGo ...
var ServiceTestGo = `package status_test

import (
	"context"
	"path/filepath"
	"testing"

	"{{ .Path }}/internal/status"
	"{{ .Path }}/pkg/api"
	"{{ .Path }}/pkg/assert"
	"{{ .Path }}/pkg/repository"
)

func TestService(t *testing.T) {
	path := filepath.Join("testdata", "repository.json")
	repo := repository.NewFileRepository(path)
	_, err := status.Service(repo)(context.Background(), &api.StatusRequest{})
	assert.That("Service should return without an error", t, err, nil)
}
`
