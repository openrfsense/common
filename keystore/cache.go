// Copyright (C) 2022 OpenRFSense
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as
// published by the Free Software Foundation, either version 3 of the
// License, or (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package keystore

import (
	"errors"
	"fmt"
	"time"

	"github.com/dgraph-io/ristretto"
)

var ErrNoKeyFound = errors.New("no key found in keystore")

var (
	keyRetriever Retriever

	cache *ristretto.Cache

	DefaultTTL = 0 * time.Second
)

// The key retriever function will be passed the requested channel name and access string.
// It is expected to return the correct key for the given parameters, or an empty string and
// an error. An example implementation of a key retriever would consist in a function
// which requests a key using a secret from a Emitter broker, or a request to a web API.
type Retriever func(string, string) (string, error)

// Initializes the internal keystore and sets a Retriever function.
func Init(retriever Retriever, ttl time.Duration) error {
	var err error
	cache, err = ristretto.NewCache(&ristretto.Config{
		NumCounters: 10000,
		MaxCost:     1000,
		BufferItems: 64,
	})
	if err != nil {
		return err
	}

	DefaultTTL = ttl
	keyRetriever = retriever
	return nil
}

// Should be safe enough to avoid collisions
func hashKey(channel string, access string) string {
	return fmt.Sprintf("%d%s%d%s", len(channel), channel, len(access), access)
}

// Arbitrarily set a key in the keystore.
func Set(channel string, access string, newKey string) {
	cache.SetWithTTL(hashKey(channel, access), newKey, 1, DefaultTTL)
	cache.Wait()
}

// Tries retrieving a key from the keystore. If a key for the specified channel and access mode is not found,
// and the retriever also returns an empty string, the retriever error is wrapped in ErrNoKeyFound and returned.
func Must(channel string, access string) (string, error) {
	if key, found := cache.Get(hashKey(channel, access)); found {
		if keyStr, ok := key.(string); ok {
			return keyStr, nil
		}
		return "", fmt.Errorf("wrong type for key: expected string, found %T", key)
	}

	if keyRetriever == nil {
		return "", fmt.Errorf("%w, key not in cache and no retriever function was set", ErrNoKeyFound)
	}

	key, err := keyRetriever(channel, access)
	if err != nil {
		return "", fmt.Errorf("%w, retriever returned error %v", ErrNoKeyFound, err)
	}

	Set(channel, access, key)
	return key, nil
}

// Tries retrieving a key from the keystore. If a key for the specified channel and access mode is not found,
// a new key is requested using the key retriever and saved in the keystore. If the key retriever function
// fails with an error, an empty string is returned and nothing is saved in the keystore.
func Get(channel string, access string) string {
	key, err := Must(channel, access)
	if err != nil {
		return ""
	}

	return key
}
