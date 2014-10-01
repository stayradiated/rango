package rangolib

import (
	"bytes"
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/spf13/cast"
	"github.com/spf13/hugo/hugolib"
	"github.com/spf13/hugo/parser"
	"github.com/spf13/hugo/source"
	"io"
	"io/ioutil"
	"path"
)

const TOML = '+'

type Page struct {
	Metadata map[string]interface{}
	Content  string
}

func Files() []*source.File {
	fs := source.Filesystem{
		Base: "content",
	}

	files := fs.Files()

	return files
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

func main() {

	files := Files()

	for _, file := range files {

		fmt.Println("Modifying: " + file.LogicalName)

		/* READING METADATA */

		psr, err := parser.ReadFrom(file.Contents)
		if err != nil {
			panic(err)
		}

		metadata, err := psr.Metadata()
		if err != nil {
			panic(err)
		}

		metadata, err = cast.ToStringMapE(metadata)
		if err != nil {
			panic(err)
		}

		/* WRITING METADATA */

		page, err := hugolib.NewPage(file.LogicalName)
		if err != nil {
			panic(err)
		}

		page.Dir = file.Dir
		page.SetSourceContent(psr.Content())
		page.SetSourceMetaData(metadata, TOML)

		page.SaveSourceAs(path.Join("content", page.FullFilePath()))

	}

	/* CONFIG */

	datum, err := ioutil.ReadFile("config.toml")
	if err != nil {
		panic(err)
	}
	config := map[string]interface{}{}
	if _, err := toml.Decode(string(datum), &config); err != nil {
		panic(err)
	}

	// editing the config
	config["random"] = "Something silly"

	buf := new(bytes.Buffer)
	if err := toml.NewEncoder(buf).Encode(config); err != nil {
		panic(err)
	}
	ioutil.WriteFile("config.toml", buf.Bytes(), 0644)
}
