package rangolib

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func setup() {
	os.Mkdir("content", 0755)
}

func teardown() {
	os.RemoveAll("content")
}

func TestReadDir(t *testing.T) {
	setup()
	assert := assert.New(t)

	files, err := ReadDir("content/")
	assert.Nil(err)

	assert.Equal(len(files), 0)

	os.Mkdir("content/foo", 0755)
	files, err = ReadDir("content/")
	assert.Nil(err)

	assert.Equal(len(files), 1)
	teardown()
}
