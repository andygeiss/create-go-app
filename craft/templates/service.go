package templates

// ServiceGo ...
var ServiceGo = `package status

import (
	"context"
	"errors"

	"{{ .Path }}/pkg/api"
)

// Service ...
func Service() api.StatusService {
	// Initialize something ...
	return func(ctx context.Context, req *api.StatusRequest) (res *api.StatusResponse, err error) {
    	// Then do great stuff ...
		return nil, errors.New("not implemented")
	}
}`
