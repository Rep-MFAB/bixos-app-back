package crypt

import (
	"math/rand"
	"sync"
	"time"
)

const uuidBytes = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

var src = rand.NewSource(time.Now().UnixNano())
var mutex sync.Mutex

// Ensures thread safety when calling src
func int63() int64 {
	mutex.Lock()
	v := src.Int63()
	mutex.Unlock()
	return v
}

// UUID returns a new unique id string at the specified size.
// This function is based on a stack overflow answer by user 'icza' at http://stackoverflow.com/a/31832326
func UUID(n int) string {
	b := make([]byte, n)

	for i, cache, remain := n-1, int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(uuidBytes) {
			b[i] = uuidBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return string(b)
}
