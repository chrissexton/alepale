package config

import (
	"bytes"
	"testing"
)

var configStr = `{
	"version":"bogus",
	"thing":{
		"this": 6
	}
}`

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
	config := getConfig()

	err := config.SetKey("thing.that", "no")
	if err != nil {
		t.Error(err)
	}

	err, val := config.GetKey("thing.that")
	if err != nil {
		t.Error(err)
	}
	if val != "no" {
		t.Error("Key set incorrectly. Value:", val)
	}

	// [todo] - Set key for non-extant path should work
}

func TestFindVal(t *testing.T) {
	config := getConfig()
	err, val := findVal("thing.this", config.values)
	if err != nil {
		t.Error(err)
	}
	if val != 6.0 {
		t.Error("Did not find correct path. Got:", val)
	}

	err, val = findVal("thing.this.doesnt.exist", config.values)
	if err == nil {
		t.Error("Should not be able to find non-extant value:", val)
	}
}

func TestSetVal(t *testing.T) {
	config := getConfig()
	err := setVal("thing.that", 10, config.values)
	if err != nil {
		t.Error(err)
	}

	err, val := findVal("thing.that", config.values)
	if err != nil {
		t.Error(err)
	}

	if val != 10 {
		t.Error("Did not find correct path. Got:", val)
	}

	err = setVal("thing.this.doesnt.exist", 10, config.values)
	if err == nil {
		t.Error("Should not be able to find non-extant value:")
	}
}
