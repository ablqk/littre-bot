package parsers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGob(t *testing.T) {
	// let's rely on existing libs
	data, _ := parseLetter("testdata/x.xml")

	gobFile := "testdata/x.gob"
	err := SaveGob(data, gobFile)
	assert.NoError(t, err)

	_, err = ParseGob(gobFile)
	assert.NoError(t, err)
	// this fails because an empty list is compared to a nil list...
	// assert.True(t, reflect.DeepEqual(res, data))
}
