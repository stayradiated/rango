package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	server *httptest.Server
	reader io.Reader
	dirUrl string
)

func init() {
	contentDir = "test_content"

	os.Mkdir("test_content", 0755)

	server = httptest.NewServer(NewRouter())
	dirUrl = fmt.Sprintf("%s/api/dir", server.URL)
}

func TestReadDir(t *testing.T) {
	assert := assert.New(t)

	url := dirUrl + "/"
	reader = strings.NewReader("")

	req, _ := http.NewRequest("GET", url, reader)
	res, err := http.DefaultClient.Do(req)
	assert.Nil(err)

	assert.Equal(res.StatusCode, http.StatusOK)

	// body := new(handleReadDirResponse)
	fmt.Println(res.Body)

	// err = json.NewDecoder(res.Body).Decode(body)
	// assert.Nil(err)

	// fmt.Println(body.Data[0])
}
