package templates

// ServiceGo ...
var ServiceGo = `package status

import (
	"context"
	"encoding/hex"

	"{{ .Path }}/pkg/api"
	"{{ .Path }}/pkg/repository"
	"{{ .Path }}/pkg/secure"
)

// Service ...
func Service(repo repository.Repository) api.StatusService {
	// Save a plaintext value
	repo.Add("status", "OK")
	// Save an encrypted password
	secret := hex.EncodeToString(secure.HashPassword([]byte("password")))
	repo.Add("secret", secret)
	return func(ctx context.Context, req *api.StatusRequest) (res *api.StatusResponse, err error) {
		// Create a response
		response := &api.StatusResponse{
			// Only read the plaintext value as an example
			Text: repo.FindByID("status").(string),
		}
		return response, nil
	}
}
`
