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

// Frontmatter stores encodeable data
type Frontmatter map[string]interface{}

// Page represents a markdown file
type Page struct {
	Path     string      `json:"path"`
	Metadata Frontmatter `json:"metadata"`
	Content  string      `json:"content"`
}

// ReadPage reads a page from disk
func ReadPage(fp string) (*Page, error) {
	// open the file for reading
	file, err := os.Open(fp)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// use the Hugo parser lib to read the contents
	parser, err := parser.ReadFrom(file)
	if err != nil {
		return nil, err
	}

	// get the metadata
	rawdata, err := parser.Metadata()
	if err != nil {
		return nil, err
	}

	// convert the interface{} into map[string]interface{}
	metadata, err := cast.ToStringMapE(rawdata)
	if err != nil {
		return nil, err
	}

	// assemble a new Page instance
	return &Page{
		Path:     fp,
		Metadata: metadata,
		Content:  string(parser.Content()),
	}, nil
}

func CreatePage(dirname string, metadata Frontmatter, content []byte) (*Page, error) {

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

func UpdatePage(fp string, metadata Frontmatter, content []byte) (*Page, error) {
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

	// delete existing page
	err := DeletePage(fp)
	if err != nil {
		return nil, err
	}

	// the filepath for the page
	// var fp string
	dirname := filepath.Dir(fp)
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

func DeletePage(fp string) error {

	// remove the directory
	return os.Remove(fp)
}
