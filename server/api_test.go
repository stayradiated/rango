package server

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stayradiated/afero"
	"github.com/stayradiated/rango/rangofs"
	"github.com/stretchr/testify/assert"
)

var (
	server *httptest.Server
	reader io.Reader
	dirUrl string
)

func init() {
	contentDir = "content"
	rangofs.Fs = new(afero.MemMapFs)

	rangofs.Fs.Mkdir("content", 0755)

	server = httptest.NewServer(Start())
	dirUrl = fmt.Sprintf("%s/api/dir", server.URL)
}

func TestReadDir(t *testing.T) {
	assert := assert.New(t)

	url := dirUrl + "/"
	reader = strings.NewReader("")

	req, _ := http.NewRequest("POST", url, reader)
	res, err := http.DefaultClient.Do(req)
	assert.Nil(err)

	assert.Equal(res.StatusCode, http.StatusOK)

	// body := new(handleReadDirResponse)
	fmt.Println(res.Body)

	// err = json.NewDecoder(res.Body).Decode(body)
	// assert.Nil(err)

	// fmt.Println(body.Data[0])
}
