package config

import (
	"encoding/json"
	"io/ioutil"
)

// LoadDefault reads and unmarshals config file using precedence order from Init()
func LoadDefault() (map[string][]string, error) {
	bytes, err := ioutil.ReadFile(ConfigFile)
	if err != nil {
		return nil, err
	}

	return Load(bytes)
}

// Load reads and unmarshals given config content
func Load(bytes []byte) (map[string][]string, error) {
	var items []string
	if err := json.Unmarshal(bytes, &items); err == nil {
		return map[string][]string{
			"default": items,
		}, nil
	}

	var data map[string][]string

	if err := json.Unmarshal(bytes, &data); err != nil {
		return nil, err
	}

	return data, nil
}

// MustLoad is like Load but panics in case of error
func MustLoad(bytes []byte) map[string][]string {
	data, err := Load(bytes)
	if err != nil {
		panic("can't load config: " + err.Error())
	}
	return data
}
