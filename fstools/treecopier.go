package fstools

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

func (c *treeCopier) visitFile(src string, info os.FileInfo) (int64, error) {
	dest := c.convertPath(src)

	sf, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer sf.Close()

	df, err := os.Create(dest)
	if err != nil {
		return 0, err
	}
	defer df.Close()

	return io.Copy(df, sf)
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
