package crypt

import (
	"crypto/sha256"
	"errors"
	"io"
)

// Enrypts the password with a sha256 hash and returns it with a salt
func Compare(password, hash, salt []byte) (bool, error) {
	if len(hash) != sha256.Size {
		return false, errors.New("Wrong sized hash received.")
	}

	// Combines salt with possible password
	combination := string(salt) + string(password)

	// Hashes salted password
	newHash := sha256.New()
	io.WriteString(newHash, combination)

	hashed := newHash.Sum(nil)

	// Slow compares the two hashes to prevent time attacks(that would never even occur, i think)
	retval := hashed[0] ^ hash[0]
	for i := range hashed {
		retval |= hashed[i] ^ hash[i]
	}

	// returns comparsion
	return retval == 0, nil
}
