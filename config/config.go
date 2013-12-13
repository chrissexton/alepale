package config

import (
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"log"
	"reflect"
	"strings"
)

type Config struct {
	values  map[string]interface{}
	Version string
}

type Value interface{}

// Run action on a path, gives the action the closest level of the struct
func pathOp(path string, action func(string, map[string]interface{}) (error, Value), values map[string]interface{}) (error, Value) {
	parts := strings.Split(path, ".")
	key := parts[0]
	remain := parts[1:]

	if len(remain) == 0 {
		return action(key, values)
	}

	if reflect.TypeOf(values[key]) == reflect.TypeOf(map[string]interface{}{}) {
		casted := values[key].(map[string]interface{})
		return pathOp(strings.Join(remain, "."), action, casted)
	}
	return errors.New("Could not find key at path: " + path), nil
}

func setVal(path string, val interface{}, values map[string]interface{}) error {
	err, _ := pathOp(path, func(key string, value map[string]interface{}) (error, Value) {
		value[key] = val
		return nil, nil
	}, values)
	return err
}

func findVal(path string, values map[string]interface{}) (error, Value) {
	return pathOp(path, func(key string, value map[string]interface{}) (error, Value) {
		return nil, value[key]
	}, values)
}

// Reads a value out of a key path
// Returns error if value at path is absent
func (c *Config) GetKey(path string) (error, Value) {
	err, val := findVal(path, c.values)
	if err != nil {
		log.Println(err)
		return err, nil
	}
	return nil, val
}

// Sets a key path in the JSON to the value given
//
// Returns errors when file is unsavable or key is unJSONable
func (c *Config) SetKey(path string, value Value) error {
	err := setVal(path, value, c.values)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

// Reads new JSON config from reader interface
// Errors obliterate program.
func ReadConfig(cFile io.Reader) *Config {
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

	return &c
}
