package rangolib

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/suite"
)

const SIMPLE_CONFIG = `title = "Something Neat"
`

var simpleConfigData = &Frontmatter{
	"title": "Something Neat",
}

type ConfigTestSuite struct {
	suite.Suite
}

func (assert *ConfigTestSuite) SetupTest() {
	configPath = "content/config.toml"
	os.Mkdir("content", 0755)
}

func (assert *ConfigTestSuite) TearDownTest() {
	os.RemoveAll("content")
}

// test ReadConfig on a simple config
func (assert *ConfigTestSuite) TestReadConfig() {
	ioutil.WriteFile("content/config.toml", []byte(SIMPLE_CONFIG), 0644)

	config, err := ReadConfig()
	assert.Nil(err)
	assert.Equal(config, simpleConfigData)
}

// test SaveConfig on a simple config
func (assert *ConfigTestSuite) TestSaveConfig() {
	err := SaveConfig(simpleConfigData)
	assert.Nil(err)

	data, err := ioutil.ReadFile("content/config.toml")
	assert.Nil(err)
	assert.Equal(string(data), SIMPLE_CONFIG)
}

// run config tests
func TestConfigTestSuite(t *testing.T) {
	suite.Run(t, new(ConfigTestSuite))
}
