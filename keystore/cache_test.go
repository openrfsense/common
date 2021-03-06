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
	"testing"
	"time"
)

func TestMust(t *testing.T) {
	err := Init(nil, DefaultTTL)
	if err != nil {
		t.Fatal(err)
	}

	t.Run("set a key then get it", func(t *testing.T) {
		Set("channel", "r", "key")
		key, err := Must("channel", "r")
		if err != nil {
			t.Fatal(err)
		}
		if key != "key" {
			t.Fail()
		}
	})

	cache.Clear()

	t.Run("set a key then get it with wrong access", func(t *testing.T) {
		Set("channel", "r", "key")
		key, err := Must("channel", "w")
		if err == nil {
			t.Fatal(err)
		}
		if key != "" {
			t.Fail()
		}
	})

	cache.Clear()

	t.Run("get non existing key", func(t *testing.T) {
		_, err := Must("channel", "r")
		if err == nil {
			t.Fail()
		}
	})

	cache.Clear()

	t.Run("get key after timeout expired", func(t *testing.T) {
		err = Init(nil, 100*time.Millisecond)
		if err != nil {
			t.Fatal(err)
		}

		Set("channel", "r", "key")
		time.Sleep(200 * time.Millisecond)
		key, err := Must("channel", "r")
		if err == nil {
			duration, fresh := cache.GetTTL(hashKey("channel", "r"))
			t.Logf("expected error, found key '%s' with duration %v and fresh: %v", key, duration, fresh)
			t.Fail()
		}
	})
}
