package rangolib

import (
	"bytes"
	"io/ioutil"

	"github.com/BurntSushi/toml"
)

func ReadConfig() (map[string]interface{}, error) {
	config := map[string]interface{}{}
	datum, err := ioutil.ReadFile("config.toml")
	if err != nil {
		return config, err
	}
	if _, err := toml.Decode(string(datum), &config); err != nil {
		return config, err
	}
	return config, nil
}

func SaveConfig(config map[string]interface{}) error {
	buf := new(bytes.Buffer)
	if err := toml.NewEncoder(buf).Encode(config); err != nil {
		return err
	}
	if err := ioutil.WriteFile("config.toml", buf.Bytes(), 0644); err != nil {
		return err
	}
	return nil
}
