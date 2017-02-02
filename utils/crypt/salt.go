package crypt

import (
	"crypto/rand"
	"crypto/sha1"
	"crypto/sha256"
	"errors"
	"fmt"
	"io"
)

const SaltSize = 16

func generateSalt(secret []byte) ([]byte, error) {
	buf := make([]byte, SaltSize, SaltSize+sha1.Size)
	_, err := io.ReadFull(rand.Reader, buf)

	if err != nil {
		return nil, errors.New(fmt.Sprintf("Random read failed: %v", err))
	}

	hash := sha256.New()
	hash.Write(buf)
	hash.Write(secret)
	return hash.Sum(buf), nil
}
