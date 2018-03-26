package auth

import (
	"crypto/rand"
)

const (
	RANDOM_SALT_LEN = 32
)

// GenerateSalt returns a random string suitable for use within the database.
func GenerateSalt() (string, error) {

	buf := make([]byte, RANDOM_SALT_LEN)
	_, err := rand.Read(buf)

	if err != nil {
		return "", err
	}

	return string(buf), nil
}
