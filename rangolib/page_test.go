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

var simplePageData = &Page{
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
}

func (assert *PageTestSuite) SetupTest() {
	os.Mkdir("content", 0755)
}

func (assert *PageTestSuite) TearDownTest() {
	os.RemoveAll("content")
}

// test ReadPage on a simple page
func (assert *PageTestSuite) TestReadPageOnSimplePage() {
	ioutil.WriteFile("content/simple-page.md", []byte(SIMPLE_PAGE), 0644)

	page, err := ReadPage("content/simple-page.md")
	assert.Nil(err)
	assert.Equal(page, simplePageData)
}

// test ReadPage on a missing page
func (assert *PageTestSuite) TestReadPageOnMissingPage() {
	page, err := ReadPage("content/simple-page.md")
	assert.NotNil(err)
	assert.Nil(page)
}

// test CreatePage on a simple page
func (assert *PageTestSuite) TestCreatePageOnSimplePage() {
	metadata := simplePageData.Metadata
	content := []byte(simplePageData.Content)

	page, err := CreatePage("content/", metadata, content)
	assert.Nil(err)
	assert.Equal(page, simplePageData)

	data, err := ioutil.ReadFile("content/simple-page.md")
	assert.Nil(err)
	assert.Equal(string(data), SIMPLE_PAGE)
}

// test UpdatePage
func (assert *PageTestSuite) TestUpdatePage() {
	os.Create("content/old-page.md")
	metadata := simplePageData.Metadata
	content := []byte(simplePageData.Content)

	page, err := UpdatePage("content/old-page.md", metadata, content)
	assert.Nil(err)
	assert.Equal(page, simplePageData)

	data, err := ioutil.ReadFile("content/simple-page.md")
	assert.Nil(err)
	assert.Equal(string(data), SIMPLE_PAGE)
}

// test DeletePage on a simple page
func (assert *PageTestSuite) TestDeletePageOnSimplePage() {
	os.Create("content/stuff.md")
	err := DeletePage("content/stuff.md")
	assert.Nil(err)
}

// test DeletePage on a missing page
func (assert *PageTestSuite) TestDeletePageOnMissingPage() {
	err := DeletePage("content/stuff.md")
	assert.NotNil(err)
}

// test generateFilePath against multiple cases
func (assert *PageTestSuite) TestGenerateFilePath() {
	var path string

	path = generateFilePath("content/", "Super Simple")
	assert.Equal(path, "content/super-simple.md")

	os.Create("content/super-simple.md")
	path = generateFilePath("content/", "Super Simple")
	assert.Equal(path, "content/super-simple-1.md")

	os.Create("content/super-simple-1.md")
	path = generateFilePath("content/", "Super Simple")
	assert.Equal(path, "content/super-simple-2.md")
}

// Run tests
func TestPageTestSuite(t *testing.T) {
	suite.Run(t, new(PageTestSuite))
}
