package rangolib

import (
	"io"
	"os"
	"path/filepath"
)

type treeCopier struct {
	srcRoot, destRoot string
}

func (c *treeCopier) convertPath(src string) string {
	return filepath.Join(c.destRoot, src[len(c.srcRoot):])
}

func (c *treeCopier) visitDir(src string, info os.FileInfo) error {
	return os.Mkdir(c.convertPath(src), info.Mode())
}

// based on https://gist.github.com/elazarl/5507969
func (c *treeCopier) visitFile(src string, info os.FileInfo) error {
	dest := c.convertPath(src)

	sf, err := os.Open(src)
	if err != nil {
		return err
	}
	// no need to check errors on read only file, we already got everything
	// we need from the filesystem, so nothing can go wrong now.
	defer sf.Close()

	df, err := os.Create(dest)
	if err != nil {
		return err
	}

	if _, err := io.Copy(df, sf); err != nil {
		df.Close()
		return err
	}

	return df.Close()
}

func (c *treeCopier) Walk(path string, info os.FileInfo, err error) error {
	if err != nil {
		return nil
	}

	if info.IsDir() {
		err = c.visitDir(path, info)
		if err != nil {
			return filepath.SkipDir
		}
	} else {
		c.visitFile(path, info)
	}

	return nil
}

// copyDir copies a directory to another location (including sub-directories)
func copyDir(srcRoot, destRoot string) error {
	c := &treeCopier{srcRoot: srcRoot, destRoot: destRoot}
	return filepath.Walk(srcRoot, c.Walk)
}

// moveDir moves a directory to another location (include sub-directories)
func moveDir(srcRoot, destRoot string) error {
	err := copyDir(srcRoot, destRoot)
	if err != nil {
		return err
	}
	return os.RemoveAll(srcRoot)
}
