package rangolib

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/suite"
)

const SIMPLE_CONFIG = `title = "Something Neat"
`

var simpleConfigData = &ConfigMap{
	"title": "Something Neat",
}

type ConfigTestSuite struct {
	suite.Suite
	Config Config
}

func (t *ConfigTestSuite) SetupTest() {
	t.Config = Config{
		path: "./content/config.toml",
	}

	os.Mkdir("content", 0755)
}

func (t *ConfigTestSuite) TearDownTest() {
	os.RemoveAll("content")
}

// test ReadConfig on a simple config
func (t *ConfigTestSuite) TestReadConfig() {
	file, _ := t.Config.Create()
	file.Write([]byte(SIMPLE_CONFIG))

	config, err := t.Config.Parse()
	t.Nil(err)
	t.Equal(config, simpleConfigData)
}

// test SaveConfig on a simple config
func (t *ConfigTestSuite) TestSaveConfig() {
	err := t.Config.Save(simpleConfigData)
	t.Nil(err)

	file, _ := t.Config.Open()
	data, err := ioutil.ReadAll(file)
	t.Nil(err)
	t.Equal(string(data), SIMPLE_CONFIG)
}

// run config tests
func TestConfigTestSuite(t *testing.T) {
	suite.Run(t, new(ConfigTestSuite))
}
