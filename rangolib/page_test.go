package rangolib

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/suite"
)

const SIMPLE_PAGE = `+++
draft = true
keywords = ["test", "awesome"]
title = "Simple Page"

+++

Hello World
===========

How are you today?
`

var simplePageFile = &PageFile{
	Path: "content/simple-page.md",
	Metadata: Frontmatter{
		"title": "Simple Page",
		"draft": true,
		"keywords": []interface{}{
			"test",
			"awesome",
		},
	},
	Content: "Hello World\n===========\n\nHow are you today?\n",
}

type PageTestSuite struct {
	suite.Suite
	Page Page
}

func (t *PageTestSuite) SetupTest() {
	t.Page = Page{}
	os.Mkdir("content", 0755)
}

func (t *PageTestSuite) TearDownTest() {
	os.RemoveAll("content")
}

// test ReadPage on a simple page
func (t *PageTestSuite) TestReadPageOnSimplePage() {
	ioutil.WriteFile("content/simple-page.md", []byte(SIMPLE_PAGE), 0644)

	page, err := t.Page.Read("content/simple-page.md")
	t.Nil(err)
	t.Equal(page, simplePageFile)
}

// test ReadPage on a missing page
func (t *PageTestSuite) TestReadPageOnMissingPage() {
	page, err := t.Page.Read("content/simple-page.md")
	t.NotNil(err)
	t.Nil(page)
}

// test CreatePage on a simple page
func (t *PageTestSuite) TestCreatePageOnSimplePage() {
	metadata := simplePageFile.Metadata
	content := []byte(simplePageFile.Content)

	page, err := t.Page.Create("content/", metadata, content)
	t.Nil(err)
	t.Equal(page, simplePageFile)

	data, err := ioutil.ReadFile("content/simple-page.md")
	t.Nil(err)
	t.Equal(string(data), SIMPLE_PAGE)
}

// test UpdatePage
func (t *PageTestSuite) TestUpdatePage() {
	os.Create("content/old-page.md")
	metadata := simplePageFile.Metadata
	content := []byte(simplePageFile.Content)

	page, err := t.Page.Update("content/old-page.md", metadata, content)
	t.Nil(err)
	t.Equal(page, simplePageFile)

	data, err := ioutil.ReadFile("content/simple-page.md")
	t.Nil(err)
	t.Equal(string(data), SIMPLE_PAGE)
}

// test DestroyPage on a simple page
func (t *PageTestSuite) TestDestroyPageOnSimplePage() {
	os.Create("content/stuff.md")
	err := t.Page.Destroy("content/stuff.md")
	t.Nil(err)
}

// test DestroyPage on a missing page
func (t *PageTestSuite) TestDestroyPageOnMissingPage() {
	err := t.Page.Destroy("content/stuff.md")
	t.NotNil(err)
}

// test generateFilePath against multiple cases
func (t *PageTestSuite) TestGenerateFilePath() {
	var path string

	path = generateFilePath("content/", "Super Simple")
	t.Equal(path, "content/super-simple.md")

	os.Create("content/super-simple.md")
	path = generateFilePath("content/", "Super Simple")
	t.Equal(path, "content/super-simple-1.md")

	os.Create("content/super-simple-1.md")
	path = generateFilePath("content/", "Super Simple")
	t.Equal(path, "content/super-simple-2.md")
}

// Run tests
func TestPageTestSuite(t *testing.T) {
	suite.Run(t, new(PageTestSuite))
}
