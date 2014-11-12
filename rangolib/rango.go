package rangolib

import (
	"bytes"
	"io"
	"io/ioutil"
	"os"
	"path"
	"strings"
	"time"

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

type DirItem struct {
	Name string `json:"name"`
	Path string `json:"path"`
}

type PageItem struct {
	Name    string    `json:"name"`
	ModTime time.Time `json:"modified_at"`
	Path    string    `json:"path"`
}

type PathList struct {
	Directories []*DirItem  `json:"directories"`
	Pages       []*PageItem `json:"pages"`
}

func NewPathList() *PathList {
	pathList := new(PathList)
	pathList.Directories = make([]*DirItem, 0)
	pathList.Pages = make([]*PageItem, 0)
	return pathList
}

func (p *PathList) AddFile(fp string, fi os.FileInfo) {
	name := fi.Name()

	if fi.IsDir() {
		p.Directories = append(p.Directories, &DirItem{
			Name: name,
			Path: fp,
		})
	} else {
		p.Pages = append(p.Pages, &PageItem{
			Name:    name,
			ModTime: fi.ModTime(),
			Path:    fp,
		})
	}
}

func Files(folder string) (*PathList, error) {
	pathList := NewPathList()
	files, err := ioutil.ReadDir(folder)

	if err != nil {
		return nil, err
	}

	for _, f := range files {
		fp := strings.TrimPrefix(path.Join(folder, f.Name()), "content/")
		pathList.AddFile(fp, f)
	}

	return pathList, nil
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
