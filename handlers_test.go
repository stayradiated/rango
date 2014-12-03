package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/stayradiated/rango/rangolib"
	"github.com/stretchr/testify/suite"
)

var (
	server     *httptest.Server
	contentDir string = "__tmp__"
)

type HandlersTestSuite struct {
	suite.Suite
}

func (assert *HandlersTestSuite) SetupTest() {
	os.Mkdir(contentDir, 0755)

	server = httptest.NewServer(NewRouter(&RouterConfig{
		Handlers: &Handlers{
			Config:     rangolib.NewConfig("config.toml"),
			Dir:        rangolib.NewDir(),
			Page:       rangolib.NewPage(),
			ContentDir: contentDir,
		},
		AdminDir: "./admin/dist",
	}))
}

func (assert *HandlersTestSuite) TearDownTest() {
	os.RemoveAll(contentDir)
	server.Close()
}

func (assert *HandlersTestSuite) TestReadDir() {
	url := server.URL + "/api/dir/"
	reader := strings.NewReader("")

	req, _ := http.NewRequest("GET", url, reader)
	res, err := http.DefaultClient.Do(req)
	assert.Nil(err)

	assert.Equal(res.StatusCode, http.StatusOK)

	var body readDirResponse
	err = json.NewDecoder(res.Body).Decode(&body)
	assert.Nil(err)
}

func TestHandlers(t *testing.T) {
	suite.Run(t, new(HandlersTestSuite))
}
