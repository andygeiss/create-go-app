package secure_test

import (
	"testing"

	"github.com/andygeiss/create-go-app/pkg/assert"
	"github.com/andygeiss/create-go-app/pkg/secure"
)

func TestDecrypt(t *testing.T) {
	plaintext := []byte("asdf 1234")
	key := secure.NewKey256()
	ciphertext, _ := secure.Encrypt(plaintext, key)
	decrypted, err := secure.Decrypt(ciphertext, key)
	assert.That("Decrypt should return without an error", t, err, nil)
	assert.That("Decrypted ciphertext should be equal with plaintext", t, decrypted, plaintext)
}

func TestEncrypt(t *testing.T) {
	plaintext := []byte("asdf 1234")
	key := secure.NewKey256()
	ciphertext, err := secure.Encrypt(plaintext, key)
	assert.That("Encrypt should return without an error", t, err, nil)
	assert.That("Ciphertext should be not empty", t, len(ciphertext) > 0, true)
}
