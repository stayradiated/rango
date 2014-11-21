package rangolib

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/suite"
)

type DirTestSuite struct {
	suite.Suite
}

func (assert *DirTestSuite) SetupTest() {
	os.Mkdir("content", 0755)
}

func (assert *DirTestSuite) TearDownTest() {
	os.RemoveAll("content")
}

// test readdir on an empty directory
func (assert *DirTestSuite) TestReadDirWithEmptyDir() {
	files, err := ReadDir("content/")
	assert.Nil(err)
	assert.Equal(len(files), 0)
}

// test readdir on a single directory
func (assert *DirTestSuite) TestReadDirWithSingleDir() {
	os.Mkdir("content/foo", 0755)
	dirInfo, _ := os.Stat("content/foo")

	files, err := ReadDir("content/")
	assert.Nil(err)
	assert.Equal(len(files), 1)

	assert.Equal(files[0], &File{
		Name:    "foo",
		Path:    "content/foo",
		IsDir:   true,
		Size:    dirInfo.Size(),
		ModTime: dirInfo.ModTime().Unix(),
	})
}

// test readdir on a single file
func (assert *DirTestSuite) TestReadDirWithSingleFile() {
	ioutil.WriteFile("content/bar.md", []byte(SIMPLE_PAGE), 0644)
	fileInfo, err := os.Stat("content/bar.md")

	files, err := ReadDir("content/")
	assert.Nil(err)
	assert.Equal(len(files), 1)

	assert.Equal(files[0], &File{
		Name:    "bar.md",
		Path:    "content/bar.md",
		IsDir:   false,
		Size:    fileInfo.Size(),
		ModTime: fileInfo.ModTime().Unix(),
	})
}

// test readdir on a non-existant directory
func (assert *DirTestSuite) TestReadDirWithMissingDir() {
	files, err := ReadDir("does_not_exist/")
	assert.Nil(files)
	assert.NotNil(err)
}

// test createdir
func (assert *DirTestSuite) TestCreateDir() {
	dir, err := CreateDir("content/foo")
	assert.Nil(err)

	dirInfo, _ := os.Stat("content/foo")

	assert.Equal(dir, &File{
		Name:    "foo",
		Path:    "content/foo",
		IsDir:   true,
		Size:    dirInfo.Size(),
		ModTime: dirInfo.ModTime().Unix(),
	})
}

// test createdir on existing directory
func (assert *DirTestSuite) TestCreateDirWithExistingDir() {
	os.Mkdir("content/foo", 0755)

	dir, err := CreateDir("content/foo")
	assert.Nil(err)
	assert.NotNil(dir)
}

// test updatedir on an existing directory with contents
func (assert *DirTestSuite) TestUpdateDirWithExistingDir() {
	os.Mkdir("content/foo", 0755)
	ioutil.WriteFile("content/foo/bar.md", []byte(SIMPLE_PAGE), 0644)

	dir, err := UpdateDir("content/foo", "content/bar")
	assert.Nil(err)

	_, err = os.Stat("content/foo")
	assert.NotNil(err)

	dirInfo, err := os.Stat("content/bar")
	assert.Nil(err)

	_, err = os.Stat("content/bar/bar.md")
	assert.Nil(err)

	bytes, err := ioutil.ReadFile("content/bar/bar.md")
	assert.Nil(err)
	assert.Equal(string(bytes), SIMPLE_PAGE)

	assert.Equal(dir, &File{
		Name:    "bar",
		Path:    "content/bar",
		IsDir:   true,
		Size:    dirInfo.Size(),
		ModTime: dirInfo.ModTime().Unix(),
	})
}

// test updatedir on a non-existant directory
func (assert *DirTestSuite) TestUpdateDirOnMissingDir() {
	dir, err := UpdateDir("content/foo", "content/bar")
	assert.NotNil(err)
	assert.Nil(dir)
}

// test updatedir on a conflicting directory
func (assert *DirTestSuite) TestUpdateDirOnConflictingDir() {
	os.Mkdir("content/foo", 0755)
	os.Mkdir("content/bar", 0755)

	dir, err := UpdateDir("content/foo", "content/bar")
	assert.NotNil(err)
	assert.Nil(dir)
}

// test deletedir on an existing directory
func (assert *DirTestSuite) TestDeleteDirWithExistingDir() {
	os.Mkdir("content/foo", 0755)

	err := DeleteDir("content/foo")
	assert.Nil(err)

	_, err = os.Stat("content/foo")
	assert.NotNil(err)
}

// test deletedir on a non-directory
func (assert *DirTestSuite) TestDeleteDirWithMissingDir() {
	err := DeleteDir("content/foo")
	assert.NotNil(err)

	ioutil.WriteFile("content/bar.md", []byte(SIMPLE_PAGE), 0644)
	err = DeleteDir("content/bar.md")
	assert.NotNil(err)
}

func TestDirTestSuite(t *testing.T) {
	suite.Run(t, new(DirTestSuite))
}
