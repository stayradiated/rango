package rangolib

import (
	"errors"
	"os"
	"path/filepath"
	"strconv"

	"github.com/kennygrant/sanitize"
	"github.com/spf13/cast"
	"github.com/spf13/hugo/hugolib"
	"github.com/spf13/hugo/parser"
)

const TOML = '+'
const YAML = '-'

type Page struct {
	Path     string                 `json:"path"`
	Metadata map[string]interface{} `json:"metadata"`
	Content  string                 `json:"content"`
}

func ReadPage(fp string) (*Page, error) {
	file, err := os.Open(fp)
	if err != nil {
		return nil, err
	}

	parser, err := parser.ReadFrom(file)
	if err != nil {
		return nil, err
	}

	rawdata, err := parser.Metadata()
	if err != nil {
		return nil, err
	}

	metadata, err := cast.ToStringMapE(rawdata)
	if err != nil {
		return nil, err
	}

	return &Page{
		Path:     fp,
		Metadata: metadata,
		Content:  string(parser.Content()),
	}, nil
}

func CreatePage(dirname string, metadata map[string]interface{}, content []byte) (*Page, error) {

	// check that title has been specified
	t, ok := metadata["title"]
	if ok == false {
		return nil, errors.New("page[meta].title must be specified")
	}

	// check that title is a string
	title, ok := t.(string)
	if ok == false {
		return nil, errors.New("page[meta].title must be a string")
	}

	// the filepath for the page
	var fp string
	count := 0

	// generate filename based on title
	// if the filename already exists, add a number on the end
	// if that exists, increment the number by one until we find a filename
	// that doesn't exist
	for {

		// combine title with count
		tmpTitle := title
		if count != 0 {
			tmpTitle += " " + strconv.Itoa(count)
		}

		filename := sanitize.Path(tmpTitle + ".md")
		fp = filepath.Join(dirname, filename)

		// only stop looping when file doesn't already exist
		if _, err := os.Stat(fp); err != nil {
			break
		}

		// add 1 to title
		count += 1
	}

	// create new hugo page
	page, err := hugolib.NewPage(fp)
	if err != nil {
		return nil, err
	}

	// set attributes
	page.SetSourceMetaData(metadata, TOML)
	page.SetSourceContent(content)

	// save page
	if err = page.SafeSaveSourceAs(fp); err != nil {
		return nil, err
	}

	// return page info
	return &Page{
		Path:     fp,
		Metadata: metadata,
		Content:  string(content),
	}, nil
}

func UpdatePage(name string, metadata map[string]interface{}, content []byte) error {
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

func DeletePage(fp string) error {

	// remove the directory
	return os.Remove(fp)
}
