package rangolib

import (
	"bytes"
	"io/ioutil"

	"github.com/BurntSushi/toml"
)

var configPath = "config.toml"

// ReadConfig reads the config from disk
func ReadConfig() (*Frontmatter, error) {

	// read data from config file
	data, err := ioutil.ReadFile(configPath)
	if err != nil {
		return nil, err
	}

	// decode toml
	config := &Frontmatter{}
	_, err = toml.Decode(string(data), config)
	if err != nil {
		return nil, err
	}

	return config, nil
}

// SaveConfig saves the config to disk
func SaveConfig(config *Frontmatter) error {

	// convert config into a string
	buf := new(bytes.Buffer)
	if err := toml.NewEncoder(buf).Encode(config); err != nil {
		return err
	}

	// write config to disk
	err := ioutil.WriteFile(configPath, buf.Bytes(), 0644)
	if err != nil {
		return err
	}

	return nil
}
