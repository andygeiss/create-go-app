package secure_test

import (
	"encoding/hex"
	"testing"

	"github.com/andygeiss/create-go-app/pkg/assert"
	"github.com/andygeiss/create-go-app/pkg/secure"
)

func TestHash(t *testing.T) {
	salt := "salt"
	data := []byte("data")
	digest := secure.Hash(salt, data)
	encoded := hex.EncodeToString(digest)
	assert.That("Encoded data should be equal ...", t, encoded, "37aede553b5b27f72a9bd9c48b04d270b42ab4cc0415b8c17c6fee53375a4452")
}
