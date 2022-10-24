package id

import (
	"math/rand"
	"time"
	"unsafe"
)

// Generates and arbitrarily long random string using the current time as the seed.
func Generate(length int) string {
	src := rand.NewSource(time.Now().Unix())

	return generateWithSource(src, length)
}

// Generates an arbitrarily long random string using an byte array as the seed.
// Only lowercase modern English alphabet are used.
func GenerateFromBytes(seed []byte, length int) string {
	intSeed := int64(0)
	for _, b := range seed {
		intSeed += int64(b)
	}
	src := rand.NewSource(intSeed)

	return generateWithSource(src, length)
}

// Generates a random string given its length and a generator function returning
// a number (int64) on call.
func generateWithSource(src rand.Source, length int) string {
	const letterBytes = "abcdefghijklmnopqrstuvwxyz"
	const (
		letterIdxBits = 6                    // 6 bits to represent a letter index
		letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
		letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
	)

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
