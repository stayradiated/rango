package rangolib

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/suite"
)

type DirTestSuite struct {
	suite.Suite
	Dir Dir
}

func (t *DirTestSuite) SetupTest() {
	t.Dir = Dir{}
	os.Mkdir("content", 0755)
}

func (t *DirTestSuite) TearDownTest() {
	os.RemoveAll("content")
}

// test readdir on an empty directory
func (t *DirTestSuite) TestReadDirWithEmptyDir() {
	files, err := t.Dir.Read("content/")
	t.Nil(err)
	t.Equal(len(files), 0)
}

// test readdir on a single directory
func (t *DirTestSuite) TestReadDirWithSingleDir() {
	os.Mkdir("content/foo", 0755)
	dirInfo, _ := os.Stat("content/foo")

	files, err := t.Dir.Read("content/")
	t.Nil(err)
	t.Equal(len(files), 1)

	t.Equal(files[0], &File{
		Name:    "foo",
		Path:    "content/foo",
		IsDir:   true,
		Size:    dirInfo.Size(),
		ModTime: dirInfo.ModTime().Unix(),
	})
}

// test readdir on a single file
func (t *DirTestSuite) TestReadDirWithSingleFile() {
	ioutil.WriteFile("content/bar.md", []byte(SIMPLE_PAGE), 0644)
	fileInfo, err := os.Stat("content/bar.md")

	files, err := t.Dir.Read("content/")
	t.Nil(err)
	t.Equal(len(files), 1)

	t.Equal(files[0], &File{
		Name:    "bar.md",
		Path:    "content/bar.md",
		IsDir:   false,
		Size:    fileInfo.Size(),
		ModTime: fileInfo.ModTime().Unix(),
	})
}

// test readdir on a non-existant directory
func (t *DirTestSuite) TestReadDirWithMissingDir() {
	files, err := t.Dir.Read("does_not_exist/")
	t.Nil(files)
	t.NotNil(err)
}

// test createdir
func (t *DirTestSuite) TestCreateDir() {
	dir, err := t.Dir.Create("content/foo")
	t.Nil(err)

	dirInfo, _ := os.Stat("content/foo")

	t.Equal(dir, &File{
		Name:    "foo",
		Path:    "content/foo",
		IsDir:   true,
		Size:    dirInfo.Size(),
		ModTime: dirInfo.ModTime().Unix(),
	})
}

// test createdir on existing directory
func (t *DirTestSuite) TestCreateDirWithExistingDir() {
	os.Mkdir("content/foo", 0755)

	dir, err := t.Dir.Create("content/foo")
	t.NotNil(err)
	t.Nil(dir)
}

// test updatedir on an existing directory with contents
func (t *DirTestSuite) TestUpdateDirWithExistingDir() {
	os.Mkdir("content/foo", 0755)
	ioutil.WriteFile("content/foo/bar.md", []byte(SIMPLE_PAGE), 0644)

	dir, err := t.Dir.Update("content/foo", "content/bar")
	t.Nil(err)

	_, err = os.Stat("content/foo")
	t.NotNil(err)

	dirInfo, err := os.Stat("content/bar")
	t.Nil(err)

	_, err = os.Stat("content/bar/bar.md")
	t.Nil(err)

	bytes, err := ioutil.ReadFile("content/bar/bar.md")
	t.Nil(err)
	t.Equal(string(bytes), SIMPLE_PAGE)

	t.Equal(dir, &File{
		Name:    "bar",
		Path:    "content/bar",
		IsDir:   true,
		Size:    dirInfo.Size(),
		ModTime: dirInfo.ModTime().Unix(),
	})
}

// test updatedir on a non-existant directory
func (t *DirTestSuite) TestUpdateDirOnMissingDir() {
	dir, err := t.Dir.Update("content/foo", "content/bar")
	t.NotNil(err)
	t.Nil(dir)
}

// test updatedir on a conflicting directory
func (t *DirTestSuite) TestUpdateDirOnConflictingDir() {
	os.Mkdir("content/foo", 0755)
	os.Mkdir("content/bar", 0755)

	dir, err := t.Dir.Update("content/foo", "content/bar")
	t.NotNil(err)
	t.Nil(dir)
}

// test deletedir on an existing directory
func (t *DirTestSuite) TestDestroyDirWithExistingDir() {
	os.Mkdir("content/foo", 0755)

	err := t.Dir.Destroy("content/foo")
	t.Nil(err)

	_, err = os.Stat("content/foo")
	t.NotNil(err)
}

// test deletedir on a non-directory
func (t *DirTestSuite) TestDestroyDirWithMissingDir() {
	err := t.Dir.Destroy("content/foo")
	t.NotNil(err)

	ioutil.WriteFile("content/bar.md", []byte(SIMPLE_PAGE), 0644)
	err = t.Dir.Destroy("content/bar.md")
	t.NotNil(err)
}

func TestDirTestSuite(t *testing.T) {
	suite.Run(t, new(DirTestSuite))
}
