package rangolib

import "os"

type File struct {
	Name    string `json:"name"`
	Path    string `json:"path"`
	IsDir   bool   `json:"isDir"`
	Size    int64  `json:"size"`
	ModTime int64  `json:"modTime"`
}

type Files []*File

// NewFile constructs a new File based on a path and file info
func NewFile(path string, info os.FileInfo) *File {
	file := &File{Path: path}
	file.Load(info)
	return file
}
func (f *File) Load(info os.FileInfo) {
	f.Name = info.Name()
	f.IsDir = info.IsDir()
	f.Size = info.Size()
	f.ModTime = info.ModTime().Unix()
}
