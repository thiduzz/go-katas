package file

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestShouldThrowExceptionWhenQuestionsAreLowerThanZero(t *testing.T) {
	_, err := ReadFile("test", -1)
	assert.EqualError(t, err, "Pick at least 1 question","Should be greater than 0")
}

func TestShouldPassWhenValidNumberIsProvided(t *testing.T) {
	res, err := ReadFile("test", 3)
	assert.NoError(t, err)
	assert.IsType(t, []string{}, res)
}