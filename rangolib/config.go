package rangolib

import (
	"bytes"
	"os"

	"github.com/BurntSushi/toml"
)

type ConfigManager interface {
	Parse() (*ConfigMap, error)
	Save(config *ConfigMap) error
}
type ConfigMap map[string]interface{}

type Config struct {
	path string
}

func NewConfig(path string) *Config {
	return &Config{
		path: path,
	}
}

// Open returns a fd to the config file
func (c Config) Open() (*os.File, error) {
	return os.Open(c.path)
}

// Create empties the config file and returns the fd to it
func (c Config) Create() (*os.File, error) {
	return os.Create(c.path)
}

// Parse converts the config file into a readable map
func (c Config) Parse() (*ConfigMap, error) {
	config := &ConfigMap{}
	file, err := c.Open()
	if err != nil {
		return nil, err
	}

	_, err = toml.DecodeReader(file, config)
	if err != nil {
		return nil, err
	}

	return config, nil
}

// Save saves the config to disk
func (c Config) Save(config *ConfigMap) error {

	// convert config into a string
	buf := new(bytes.Buffer)
	if err := toml.NewEncoder(buf).Encode(config); err != nil {
		return err
	}

	// write config to disk
	file, err := c.Create()
	if err != nil {
		return err
	}

	_, err = buf.WriteTo(file)
	return err
}
