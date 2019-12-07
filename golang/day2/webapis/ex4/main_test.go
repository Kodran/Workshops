package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	myTestFunc := FunctionToTest()
	assert.Equal(t, myTestFunc, "my test")
}
