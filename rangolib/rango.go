package rangolib

import (
	"io"
	"io/ioutil"
	"os"
	"time"

	"github.com/spf13/cast"
	"github.com/spf13/hugo/parser"
)

const TOML = '+'

type Page struct {
	Metadata map[string]interface{} `json:"metadata"`
	Content  string                 `json:"content"`
}

type DirItem struct {
	Name string `json:"name"`
}

type PageItem struct {
	Name    string    `json:"name"`
	ModTime time.Time `json:"modified_at"`
}

type PathList struct {
	Directories []*DirItem  `json:"directories"`
	Pages       []*PageItem `json:"pages"`
}

func NewPathList() *PathList {
	pathList := new(PathList)
	pathList.Directories = make([]*DirItem, 0)
	pathList.Pages = make([]*PageItem, 0)
	return pathList
}

func (p *PathList) AddFile(fi os.FileInfo) {
	name := fi.Name()
	if fi.IsDir() {
		p.Directories = append(p.Directories, &DirItem{
			Name: name,
		})
	} else {
		p.Pages = append(p.Pages, &PageItem{
			Name:    name,
			ModTime: fi.ModTime(),
		})
	}
}

func Files(path string) (*PathList, error) {
	pathList := NewPathList()
	files, err := ioutil.ReadDir(path)

	if err != nil {
		return nil, err
	}

	for _, f := range files {
		pathList.AddFile(f)
	}

	return pathList, nil
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
