package config

import (
	"testing"

	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/rawbytes"
	"github.com/knadh/koanf/providers/structs"
)

var config = []byte(`
backend:
  port: "8081"
  users:
    openrfsense: openrfsense
`)

func TestLoad(t *testing.T) {
	cfg := struct {
		Port  int
		Users map[string]string
	}{
		Port: 8081,
		Users: map[string]string{
			"openrfsense": "openrfsense",
		},
	}

	conf = koanf.New(".")
	conf.Load(structs.Provider(cfg, ""), nil)

	err := conf.Load(rawbytes.Provider(config), yaml.Parser())
	if err != nil {
		t.Fatal(err)
	}

	err = conf.Unmarshal("", &cfg)
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetWeak(t *testing.T) {
	conf = koanf.New(".")
	err := conf.Load(rawbytes.Provider(config), yaml.Parser())
	if err != nil {
		t.Fatal(err)
	}

	portStr := GetWeakString("backend.port")
	if portStr != "8081" {
		t.Logf("Got port %v (%T), expected '8081' (string)", portStr, portStr)
		t.Fail()
	}

	portInt := GetWeakInt("backend.port")
	if portInt != 8081 {
		t.Logf("Got port %v (%T), expected 8081 (int)", portInt, portInt)
		t.Fail()
	}
}
