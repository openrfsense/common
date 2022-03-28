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
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/providers/structs"
)

var conf *koanf.Koanf

// Formats enviroment variables: ORFS_SECTION_SUBSECTION_KEY becomes
// (as a path) section.subsection.key
func formatEnv(s string) string {
	rawPath := strings.ToLower(strings.TrimPrefix(s, "ORFS_"))
	return strings.Replace(rawPath, "_", ".", -1)
}

// Loads a YAML configuration file from the given path and overrides
// it with environment variables. If the file cannot be loaded or
// parsed as YAML, an error is returned. Requires a default config of any kind,
// will try to serialize the configuration to outConfig if present (needs to
// be a pointer to a struct).
func Load(path string, defaultConfig interface{}, outConfig ...interface{}) error {
	conf = koanf.New(".")

	if strings.TrimSpace(path) == "" {
		return fmt.Errorf("configuration file path cannot be empty")
	}

	conf.Load(structs.Provider(defaultConfig, ""), nil)

	if err := conf.Load(file.Provider(path), yaml.Parser()); err != nil {
		return fmt.Errorf("error loading configuration file: %v (%T)", err, err)
	}

	if len(outConfig) > 0 {
		err := conf.Unmarshal("", outConfig[0])
		if err != nil {
			return err
		}
	}

	conf.Load(env.Provider("ORFS_", ".", formatEnv), nil)
	return nil
}

// Returns value associated with path and nil if no value is found at path or value
// cannot be cast to type T.
func Get[T comparable](path string) T {
	var void T

	if conf == nil {
		return void
	}

	value := conf.Get(path)
	if value == nil {
		return void
	}

	if v, ok := value.(T); ok {
		return T(v)
	}

	return void
}

// Returns integer value associated with path and nil if no value is found at path or value
// cannot be cast to int. Will try converting a string found at path to integer.
func GetWeakInt(path string) int {
	var voidInt int
	var voidStr string

	if v := Get[int](path); v != voidInt {
		return v
	}

	if vStr := Get[string](path); vStr != voidStr {
		v, err := strconv.Atoi(vStr)
		if err != nil {
			return voidInt
		}
		return v
	}

	return voidInt
}

// Returns string value associated with path and nil if no value is found at path or value
// cannot be converted to string. Will convert an integer found at path to string.
func GetWeakString(path string) string {
	var voidInt int
	var voidStr string

	if v := Get[string](path); v != voidStr {
		return v
	}

	if vInt := Get[int](path); vInt != voidInt {
		return strconv.Itoa(vInt)
	}

	return voidStr
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

// Returns value found at path or calls log.Fatalf if the path does not exist or value
// cannot be cast to type T.
func Must[T comparable](path string) T {
	var void T

	if conf == nil {
		log.Fatalf("configuration was not initialized")
	}

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

// Returns map value associated with path and an calls log.Fatalf if no value is found
// at path, value cannot be cast to map[K]V or any value in the map cannot be cast to type V.
func MustMap[K comparable, V any](path string) map[K]V {
	var void map[K]V
	var voidValue V

	if conf == nil {
		log.Fatalf("configuration was not initialized")
	}

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
