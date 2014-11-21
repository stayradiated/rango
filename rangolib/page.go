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

//  ┌┬┐┬ ┬┌─┐┌─┐┌─┐
//   │ └┬┘├─┘├┤ └─┐
//   ┴  ┴ ┴  └─┘└─┘

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

func (p *Page) Save() error {
	// create new hugo page
	page, err := hugolib.NewPage(p.Path)
	if err != nil {
		return err
	}

	// set attributes
	page.SetSourceMetaData(p.Metadata, TOML)
	page.SetSourceContent([]byte(p.Content))

	// save page
	return page.SafeSaveSourceAs(p.Path)
}

//  ┌─┐┬ ┬┌┐┌┌─┐┌┬┐┬┌─┐┌┐┌┌─┐
//  ├┤ │ │││││   │ ││ ││││└─┐
//  └  └─┘┘└┘└─┘ ┴ ┴└─┘┘└┘└─┘

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

// CreatePage creates a new file and saves page content to it
func CreatePage(dirname string, metadata Frontmatter, content []byte) (*Page, error) {

	// get title from metadata
	title, err := getTitle(metadata)
	if err != nil {
		return nil, err
	}

	// the filepath for the page
	fp := generateFilePath(dirname, title)

	// create a new page
	page := &Page{
		Path:     fp,
		Metadata: metadata,
		Content:  string(content),
	}

	// save page to disk
	err = page.Save()
	if err != nil {
		return nil, err
	}

	return page, nil
}

// UpdatePage changes the content of an existing page
func UpdatePage(fp string, metadata Frontmatter, content []byte) (*Page, error) {

	// get title from metadata
	title, err := getTitle(metadata)
	if err != nil {
		return nil, err
	}

	// delete existing page
	err = DeletePage(fp)
	if err != nil {
		return nil, err
	}

	// the filepath for the page
	dirname := filepath.Dir(fp)
	fp = generateFilePath(dirname, title)

	// create a new page
	page := &Page{
		Path:     fp,
		Metadata: metadata,
		Content:  string(content),
	}

	// save page to disk
	err = page.Save()
	if err != nil {
		return nil, err
	}

	return page, nil
}

// DeletePage deletes a page
func DeletePage(fp string) error {

	// check that file exists
	info, err := os.Stat(fp)
	if err != nil {
		return err
	}

	// that file is a directory
	if info.IsDir() {
		return errors.New("DeletePage cannot delete directories")
	}

	// remove the directory
	return os.Remove(fp)
}

//  ┬ ┬┌─┐┬  ┌─┐┌─┐┬─┐┌─┐
//  ├─┤├┤ │  ├─┘├┤ ├┬┘└─┐
//  ┴ ┴└─┘┴─┘┴  └─┘┴└─└─┘

// generateFilePath generates a filepath based on a page title
// if the filename already exists, add a number on the end
// if that exists, increment the number by one until we find a filename
// that doesn't exist
func generateFilePath(dirname, title string) (fp string) {
	count := 0

	for {

		// combine title with count
		name := title
		if count != 0 {
			name += " " + strconv.Itoa(count)
		}

		// join filename with dirname
		filename := sanitize.Path(name + ".md")
		fp = filepath.Join(dirname, filename)

		// only stop looping when file doesn't already exist
		if _, err := os.Stat(fp); err != nil {
			break
		}

		// try again with a different number
		count += 1
	}

	return fp
}

func getTitle(metadata Frontmatter) (string, error) {

	// check that title has been specified
	t, ok := metadata["title"]
	if ok == false {
		return "", errors.New("page[meta].title must be specified")
	}

	// check that title is a string
	title, ok := t.(string)
	if ok == false {
		return "", errors.New("page[meta].title must be a string")
	}

	return title, nil
}
