package services

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const TestPassword = "test_password+1234"

func TestHashPassword(t *testing.T) {
	assert := assert.New(t)
	auth := Auth{}

	result, err := auth.HashPassword(TestPassword)
	assert.Nil(err)
	assert.NotEmpty(result)

	t.Log(result)
}

func TestCompareHashAndPass(t *testing.T) {
	assert := assert.New(t)
	auth := Auth{}
	result, err := auth.HashPassword(TestPassword)

	if err != nil {
		logged := auth.ValidatePassword(result, TestPassword)
		assert.Nil(logged)
	}
}
