package keystore

import (
	"errors"
	"fmt"
)

var ErrNoKeyFound = errors.New("no key found in keystore")

var keyMap = make(map[string]map[string]string)

// Arbitrarily set a key in the keystore.
func Set(channel string, access string, newKey string) {
	if _, ok := keyMap[channel]; !ok {
		keyMap[channel] = make(map[string]string)
	}
	keyMap[channel][access] = newKey
}

// Tries retrieving a key from the keystore. If a key for the specified channel and access mode is not found,
// an error wrapping ErrNoKeyFound is returned.
func Must(channel string, access string) (string, error) {
	if acc, ok := keyMap[channel]; ok {
		if key, kOk := acc[access]; kOk {
			return key, nil
		}

		return "", fmt.Errorf("%w: no key registered for channel %s and access %s", ErrNoKeyFound, channel, access)
	}

	return "", fmt.Errorf("%w: keystore does not contain channel %s", ErrNoKeyFound, channel)
}

// Tries retrieving a key from the keystore. If a key for the specified channel and access mode is not found,
// a new key is requested to the broker and saved in the keystore.
func Get(channel string, access string) string {
	if acc, ok := keyMap[channel]; ok {
		if key, kOk := acc[access]; kOk {
			return key
		}
	} else {
		keyMap[channel] = make(map[string]string)
	}

	key := ""
	keyMap[channel][access] = key
	return key
}
