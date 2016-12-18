package passhash

import (
	"encoding/hex"
	"os"

	"golang.org/x/crypto/scrypt"
)

var salt = []byte(os.Getenv("PASS_SALT"))

func HashString(pass string) (string, error) {
	key, err := scrypt.Key([]byte(pass), salt, 16384, 8, 1, 32)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(key[:]), nil
}
