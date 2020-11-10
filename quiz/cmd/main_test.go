package main

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestArgumentsAreIntegers(t *testing.T) {

	os.Args = append(os.Args, "-questions=1")
	os.Args = append(os.Args, "-seconds=3")
	os.Args = append(os.Args, "-filename=oioioi.csv")

	num1, num2, filename := readParams()
	assert.Equal(t, num1, 1)
	assert.Equal(t, num2, 3)
	assert.Equal(t, filename, "oioioi.csv")
}

func TestArgumentsHaveDefaults(t *testing.T) {

	questions, seconds, filename := readParams()
	assert.Equal(t, questions, 20)
	assert.Equal(t, seconds, 60)
	assert.Equal(t, filename, "problems.csv")
}