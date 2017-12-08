package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReflectionBullshit(t *testing.T) {

	// test #1 -- sample non-present model
	res, err := NewConfig("meowmeow__asdasdas a")
	assert.NotNil(t, err)
	assert.Nil(t, res)

	// test #2 -- use sample model for testing
	res, err = NewConfig("child")
	assert.NotNil(t, res)
	assert.Nil(t, err)
}
