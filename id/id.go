package id

import (
	"math/rand"
	"unsafe"
)

// Generates an arbitrarily long random string using an byte array as the seed.
// Only lowercase modern English alphabet are used.
func Generate(seed []byte, length int) string {
	const letterBytes = "abcdefghijklmnopqrstuvwxyz"
	const (
		letterIdxBits = 6                    // 6 bits to represent a letter index
		letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
		letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
	)

	intSeed := int64(0)
	for _, b := range seed {
		intSeed += int64(b)
	}

	src := rand.NewSource(intSeed)
	b := make([]byte, length)
	for i, cache, remain := length-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return *(*string)(unsafe.Pointer(&b))
}
