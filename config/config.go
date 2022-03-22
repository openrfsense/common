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

package config

import (
	"errors"
	"fmt"
	"io/fs"
	"log"
	"strings"

	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/providers/file"
)

var conf *koanf.Koanf

// Formats enviroment variables: ORFS_SECTION_SUBSECTION_KEY becomes
// (as a path) section.subsection.key
func formatEnv(s string) string {
	rawPath := strings.ToLower(strings.TrimPrefix(s, "ORFS_"))
	return strings.Replace(rawPath, "_", ".", -1)
}

// Loads YAML configuration files sequentially from given paths and overrides
// everything with environment variables. The paths are loaded in order of appearance,
// so later files override earlier ones.
// TODO: document errors
func LoadConfig(paths ...string) error {
	conf = koanf.New(".")

	for _, p := range paths {
		if err := conf.Load(file.Provider(p), yaml.Parser()); err != nil {
			if errors.Is(err, fs.ErrNotExist) {
				log.Printf("configuration file %s not found, skipping", p)
			} else {
				log.Fatalf("error loading configuration file: %v (%T)", err, err)
			}
			// TODO: validate
		}
	}

	if len(conf.Keys()) == 0 {
		return fmt.Errorf("no configuration file was loaded")
	}

	conf.Load(env.Provider("ORFS_", ".", formatEnv), nil)
	return nil
}

// Returns value associated with path and nil if no value is found at path or value
// cannot be cast to type T.
func Get[T comparable](path string) T {
	var void T

	value := conf.Get(path)
	if value == nil {
		return void
	}

	if v, ok := value.(T); ok {
		return T(v)
	}

	return void
}

// Returns map value associated with path and an empty map if no value is found at path,
// value cannot be cast to map[K]V or any value in the map cannot be cast to type V.
func GetMap[K comparable, V any](path string) map[K]V {
	var void map[K]V

	obj := conf.Get(path)
	if obj == nil {
		return void
	}

	if m, ok := obj.(map[K]interface{}); ok {
		ret := make(map[K]V)
		for k, v := range m {
			if vCast, vOk := v.(V); vOk {
				ret[k] = vCast
			} else {
				return void
			}
		}
		return ret
	}

	return void
}

// Returns value associated with path and fallback if no value is found at path or value
// cannot be cast to type T.
func GetOrDefault[T comparable](path string, fallback T) T {
	var void T

	v := Get[T](path)
	if v == void {
		return fallback
	}

	return v
}

// Returns value found at path or panics if the path does not exist or value cannot be
// cast to type T.
func Must[T comparable](path string) T {
	var void T

	value := conf.Get(path)
	if value == nil {
		log.Fatalf("no value found for path %s", path)
	}

	if v, ok := value.(T); ok {
		return T(v)
	}

	log.Fatalf("invalid value %#v (%T) for type %T", conf.Get(path), conf.Get(path), void)
	return void
}

func MustMap[K comparable, V any](path string) map[K]V {
	var void map[K]V
	var voidValue V

	obj := conf.Get(path)
	if obj == nil {
		log.Fatalf("no value found for path %s", path)
	}

	if m, ok := obj.(map[K]interface{}); ok {
		ret := make(map[K]V)
		for k, v := range m {
			if vCast, vOk := v.(V); vOk {
				ret[k] = vCast
			} else {
				log.Fatalf(
					"invalid value %#v (%T) for value type %T",
					conf.Get(path), conf.Get(path), voidValue,
				)
			}
		}
		return ret
	}

	log.Fatalf("invalid value %#v (%T) for map type %T", conf.Get(path), conf.Get(path), void)
	return void
}
