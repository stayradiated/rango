package server

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var (
	server *httptest.Server
	reader io.Reader
	dirUrl string
)

func init() {
	contentDir = "test_content"

	server = httptest.NewServer(Start())
	dirUrl = fmt.Sprintf("%s/api/dir", server.URL)
}

func TestReadDir(t *testing.T) {
	reader = strings.NewReader("")

	url := dirUrl + "/"

	req, err := http.NewRequest("GET", url, reader)
	if err != nil {
		t.Error(err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Error(err)
	}

	if res.StatusCode != http.StatusOK {
		t.Errorf("Success expected: %d", res.StatusCode)
	}

	body := new(handleReadDirResponse)
	err = json.NewDecoder(res.Body).Decode(body)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(body.Data[0])
}
