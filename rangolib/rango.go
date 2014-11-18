package rangolib

import (
	"bytes"
	"io"
	"io/ioutil"

	"github.com/BurntSushi/toml"
	"github.com/spf13/cast"
	"github.com/spf13/hugo/hugolib"
	"github.com/spf13/hugo/parser"
)

const TOML = '+'

type Page struct {
	Metadata map[string]interface{} `json:"metadata"`
	Content  string                 `json:"content"`
}

func Read(file io.Reader) (page *Page, err error) {
	psr, err := parser.ReadFrom(file)
	if err != nil {
		return page, err
	}

	rawdata, err := psr.Metadata()
	if err != nil {
		return page, err
	}

	metadata, err := cast.ToStringMapE(rawdata)
	if err != nil {
		return page, err
	}

	return &Page{
		Metadata: metadata,
		Content:  string(psr.Content()),
	}, nil
}

func Save(name string, metadata map[string]interface{}, content []byte) error {
	page, err := hugolib.NewPage(name)
	if err != nil {
		return err
	}

	page.SetSourceMetaData(metadata, TOML)
	page.SetSourceContent(content)

	if err = page.SafeSaveSourceAs(name); err != nil {
		return err
	}

	return nil
}

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
