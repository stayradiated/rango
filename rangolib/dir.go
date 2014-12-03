package rangolib

import (
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
)

type DirManager interface {
	Read(string) (Files, error)
	Create(string) (*File, error)
	Update(string, string) (*File, error)
	Destroy(string) error
}

type Dir struct{}

func NewDir() *Dir {
	return &Dir{}
}

// Read lists the contents of a directory
func (d Dir) Read(dirname string) (Files, error) {
	contents, err := ioutil.ReadDir(dirname)
	if err != nil {
		return nil, err
	}

	// make a new slice of File's to hold the dir contents
	files := make(Files, len(contents))

	// convert os.FileInfo into Files
	for i, info := range contents {
		files[i] = NewFile(filepath.Join(dirname, info.Name()), info)
	}

	return files, nil
}

// Create creates a new directory
func (d Dir) Create(dirname string) (*File, error) {

	// make directory
	if err := os.Mkdir(dirname, 0755); err != nil {
		return nil, err
	}

	// check that directory was created
	info, err := os.Stat(dirname)
	if err != nil {
		return nil, err
	}

	// convert fileinfo into something we can print
	return NewFile(dirname, info), nil
}

// Update renames an existing directory
func (d Dir) Update(src string, dest string) (*File, error) {

	// check that destination doesn't exist
	info, err := os.Stat(dest)
	if info != nil {
		return nil, errors.New("Cannot overwrite destination")
	}

	// move directory including it's contents
	if err := moveDir(src, dest); err != nil {
		return nil, err
	}

	// check that directory was created
	info, err = os.Stat(dest)
	if err != nil {
		return nil, err
	}

	// convert fileinfo into something we can print
	return NewFile(dest, info), nil
}

// Destroy will delete a directory and it's contents
func (d Dir) Destroy(dirname string) error {

	// check that directory exists
	dir, err := os.Stat(dirname)
	if err != nil {
		return err
	}

	// check that directory is a directory
	if dir.IsDir() == false {
		return errors.New("DeleteDir can only delete directories")
	}

	// remove the directory
	return os.RemoveAll(dirname)
}
