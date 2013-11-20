package config

import (
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"log"
	"strings"
)

type Config struct {
	values  map[string]interface{}
	Version string
}

type Value interface{}

func findVal(pathRemaining, key string) Value {
	path := strings.Split(pathRemaining, ".")
	if key == "" {
		key = path[len(pathRemaining)-1]
		pathRemaining = strings.Join(path[0:len(pathRemaining)-1], ".")
	}
	return findVal(pathRemaining, key)
}

// Reads a value out of a key path
// Returns error if value at path is absent
func (c *Config) GetKey(path string) (error, Value) {
	val, ok := c.values[path]
	if !ok {
		return errors.New("Could not find config value for path: " + path), nil
	}
	return nil, val
}

// Sets a key path in the JSON to the value given
//
// Returns errors when file is unsavable or key is unJSONable
func (c *Config) SetKey(path string, value Value) error {
	return nil
}

// Reads new JSON config from reader interface
// Errors obliterate program.
func ReadConfig(cFile io.Reader) *Config {
	log.Println("Reading config file")
	file, e := ioutil.ReadAll(cFile)
	if e != nil {
		log.Fatal("Could not read config file")
	}

	var c Config
	err := json.Unmarshal(file, &c.values)
	if err != nil {
		log.Fatal(err)
	}

	version, ok := c.values["version"]
	if !ok {
		log.Fatal("No version defined in config file")
	}
	c.Version = version.(string)

	log.Printf("AlePale version %s read config\n", c.Version)

	return &c
}
