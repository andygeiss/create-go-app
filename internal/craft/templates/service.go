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
	// Save an encrypted password and key
	key := secure.NewKey256()
	ciphertext, _ := secure.Encrypt([]byte("secret_data"), key)
	repo.Add("key", hex.EncodeToString(key[:]))
	repo.Add("secret_data", hex.EncodeToString(ciphertext))
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
