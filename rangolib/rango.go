package rangolib

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/spf13/cast"
	"github.com/spf13/hugo/parser"
)

const TOML = '+'

type Page struct {
	Metadata map[string]interface{}
	Content  string
}

type DirItem struct {
	Name     string    `json:"name"`
	Contents *PathList `json:"contents"`
}

type FileItem struct {
	Name    string    `json:"name"`
	ModTime time.Time `json:"modified_at"`
}

type PathList struct {
	Directories []*DirItem  `json:"directories"`
	Files       []*FileItem `json:"files"`
}

func NewPathList() *PathList {
	pathList := new(PathList)
	pathList.Directories = make([]*DirItem, 0)
	pathList.Files = make([]*FileItem, 0)
	return pathList
}

func (p *PathList) Add(filePath string, fi os.FileInfo) {
	sections := make([]string, 0)
	name := ""

	for {
		filePath, name = filepath.Split(filePath)
		sections = append([]string{name}, sections...)
		filePath = strings.Trim(filePath, "/")
		if len(filePath) == 0 {
			break
		}
	}

	pathList := p
	length := len(sections)
	isFile := !fi.IsDir()

	for i, path := range sections {

		if isFile && i == length-1 {
			pathList.Files = append(pathList.Files, &FileItem{
				Name:    path,
				ModTime: fi.ModTime(),
			})

		} else {

			found := false

			for _, dirItem := range pathList.Directories {
				if dirItem.Name == path {
					pathList = dirItem.Contents
					found = true
					break
				}
			}

			if found {
				continue
			}

			newPathList := NewPathList()
			newDirItem := &DirItem{
				Name:     path,
				Contents: newPathList,
			}

			pathList.Directories = append(pathList.Directories, newDirItem)
			pathList = newPathList
		}
	}
}

func Files() *PathList {
	root := NewPathList()
	baseDir := "content"

	filepath.Walk(baseDir, func(filePath string, fi os.FileInfo, err error) error {
		if err != nil {
			return nil
		}

		if fi.IsDir() && isNonProcessablePath(filePath) {
			return filepath.SkipDir
		} else if isNonProcessablePath(filePath) {
			return nil
		}

		root.Add(filePath, fi)
		return nil
	})

	// return the contents of the baseDir directory
	return root.Directories[0].Contents
}

func isNonProcessablePath(filePath string) bool {
	base := filepath.Base(filePath)
	fmt.Println(base)
	if base[0] == '.' {
		return true
	}

	if base[0] == '#' {
		return true
	}

	if base[len(base)-1] == '~' {
		return true
	}

	return false
}

func Read(file io.Reader) (page *Page, err error) {
	psr, err := parser.ReadFrom(file)
	if err != nil {
		return page, err
	}

	rawdata, err := psr.Metadata()
	if err != nil {
		return page, err
	}

	metadata, err := cast.ToStringMapE(rawdata)
	if err != nil {
		return page, err
	}

	return &Page{
		Metadata: metadata,
		Content:  string(psr.Content()),
	}, nil
}

// func main() {
//
// 	files := Files()
//
// 	for _, file := range files {
//
// 		fmt.Println("Modifying: " + file.LogicalName())
//
// 		/* READING METADATA */
//
// 		psr, err := parser.ReadFrom(file.Contents)
// 		if err != nil {
// 			panic(err)
// 		}
//
// 		metadata, err := psr.Metadata()
// 		if err != nil {
// 			panic(err)
// 		}
//
// 		metadata, err = cast.ToStringMapE(metadata)
// 		if err != nil {
// 			panic(err)
// 		}
//
// 		/* WRITING METADATA */
//
// 		page, err := hugolib.NewPage(file.LogicalName())
// 		if err != nil {
// 			panic(err)
// 		}
//
// 		// page.Dir = file.Dir
// 		page.SetSourceContent(psr.Content())
// 		page.SetSourceMetaData(metadata, TOML)
//
// 		page.SaveSourceAs(path.Join("content", page.FullFilePath()))
//
// 	}
//
// 	/* CONFIG */
//
// 	datum, err := ioutil.ReadFile("config.toml")
// 	if err != nil {
// 		panic(err)
// 	}
// 	config := map[string]interface{}{}
// 	if _, err := toml.Decode(string(datum), &config); err != nil {
// 		panic(err)
// 	}
//
// 	// editing the config
// 	config["random"] = "Something silly"
//
// 	buf := new(bytes.Buffer)
// 	if err := toml.NewEncoder(buf).Encode(config); err != nil {
// 		panic(err)
// 	}
// 	ioutil.WriteFile("config.toml", buf.Bytes(), 0644)
// }
