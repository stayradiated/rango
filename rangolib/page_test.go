package rangolib

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/suite"
)

const SIMPLE_PAGE = `
+++
title = "Simple Page"
+++
# Hello World
`

type PageTestSuite struct {
	suite.Suite
}

func (assert *PageTestSuite) SetupTest() {
	os.Mkdir("content", 0755)
}

func (assert *PageTestSuite) TearDownTest() {
	os.RemoveAll("content")
}

// test readpage on a simple page
func (assert *PageTestSuite) TestReadPageOnSimplePage() {
	ioutil.WriteFile("content/bar.md", []byte(SIMPLE_PAGE), 0644)

	page, err := ReadPage("content/bar.md")
	assert.Nil(err)
	assert.Equal(page, &Page{
		Path: "content/bar.md",
		Metadata: Frontmatter{
			"title": "Simple Page",
		},
		Content: "# Hello World\n",
	})
}

func TestPageTestSuite(t *testing.T) {
	suite.Run(t, new(PageTestSuite))
}
