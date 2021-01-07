package templates

// ServiceGo ...
var ServiceGo = `package status

import (
	"context"

	"{{ .Path }}/pkg/api"
	"{{ .Path }}/pkg/repository"
)

// Service ...
func Service(repo repository.Repository) api.StatusService {
	repo.Add("status", "OK")
	return func(ctx context.Context, req *api.StatusRequest) (res *api.StatusResponse, err error) {
		return &api.StatusResponse{
			Text: repo.FindByID("status").(string),
		}, nil
	}
}`
