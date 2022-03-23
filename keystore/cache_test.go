package keystore

import "testing"

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
}
