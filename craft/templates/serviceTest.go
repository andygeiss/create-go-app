package templates

// ServiceTestGo ...
var ServiceTestGo = `package status_test

import (
	"context"
	"testing"

	"{{ .Path }}/internal/status"
	"{{ .Path }}/pkg/api"
	"{{ .Path }}/pkg/assert"
)

func TestService(t *testing.T) {
	_, err := status.Service()(context.Background(), &api.StatusRequest{})
	assert.That("Service should return without an error", t, err, nil)
}`
