package config

import (
	"bytes"
	"testing"
)

var configStr = `{"version":"bogus"}`

func getConfig() *Config {
	buf := bytes.NewBufferString(configStr)
	config := ReadConfig(buf)

	return config
}

func TestReadConfig(t *testing.T) {
	config := getConfig()

	if config.Version != "bogus" {
		t.Error("Version mismatch! bogus !=", config.Version)
	}
}

func TestGetKey(t *testing.T) {
	config := getConfig()

	err, version := config.GetKey("version")
	if err != nil {
		t.Error(err)
	}

	if version != "bogus" {
		t.Error("Incorrect version")
	}
}

func TestSetKey(t *testing.T) {
	//t.Error("Not implemented")
}

func TestFindVal(t *testing.T) {
	findVal("my.dotted.path", "thing")
}
