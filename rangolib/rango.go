package rangolib

import (
	"bytes"
	"io"
	"io/ioutil"
	"os"
	"path"
	"strings"

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

type File struct {
	Name    string `json:"name"`
	Path    string `json:"path"`
	IsDir   bool   `json:"isDir"`
	Size    int64  `json:"size"`
	ModTime int64  `json:"mtime"`
}

func (f *File) Load(info os.FileInfo) {
	f.Name = info.Name()
	f.IsDir = info.IsDir()
	f.Size = info.Size()
	f.ModTime = info.ModTime().Unix()
}

func NewFile(dirname string, info os.FileInfo) *File {
	path := path.Join(dirname, info.Name())
	path = strings.TrimPrefix(path, "content")

	file := &File{Path: path}
	file.Load(info)
	return file
}

func DirContents(dirname string) ([]*File, error) {
	files := make([]*File, 0)
	contents, err := ioutil.ReadDir(dirname)
	if err != nil {
		return nil, err
	}

	for _, info := range contents {
		file := NewFile(dirname, info)
		files = append(files, file)
	}

	return files, nil
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
