package secure_test

import (
	"testing"

	"github.com/andygeiss/create-go-app/pkg/assert"
	"github.com/andygeiss/create-go-app/pkg/secure"
)

func TestCreatePasswordHash(t *testing.T) {
	password := "asdf 1234"
	hash := secure.CreatePasswordHash(password)
	assert.That("Password Hash should be valid", t, secure.IsPasswordHashValid(password, hash))
}
