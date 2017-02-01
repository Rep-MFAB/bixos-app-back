package crypt

import (
	"crypto/sha256"
	"io"
)

// Enrypts the password with a sha256 hash and returns it with a salt
func Encrypt(password []byte) ([]byte, []byte, error) {
	salt, err := generateSalt(password)
	if err != nil {
		return nil, nil, err
	}

	// Combines salt with password
	combination := string(salt) + string(password)

	// Hashes salted passwordand returns
	passwordHash := sha256.New()
	io.WriteString(passwordHash, combination)
	return passwordHash.Sum(nil), salt, nil
}
