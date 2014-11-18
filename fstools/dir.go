package fstools

import (
	"os"
	"path/filepath"
)

// CopyDir copies a directory to another location (including sub-directories)
func CopyDir(srcRoot, destRoot string) error {
	c := &treeCopier{srcRoot: srcRoot, destRoot: destRoot}
	return filepath.Walk(srcRoot, c.Walk)
}

// MoveDir moves a directory to another location (include sub-directories)
func MoveDir(srcRoot, destRoot string) error {
	err := CopyDir(srcRoot, destRoot)
	if err != nil {
		return err
	}
	return os.RemoveAll(srcRoot)
}
