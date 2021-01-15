package secure

const (
	"golang.org/x/crypto/bcrypt"
)

// HashPassword ...
func HashPassword(password []byte) (hash []byte) {
	return bcrypt.GenerateFromPassword(password, 14)
}

// IsPasswordHashValid ...
func IsPasswordHashValid(password, hash []byte) (ok bool) {
	return bcrypt.CompareHashAndPassword(hash, password)
}
