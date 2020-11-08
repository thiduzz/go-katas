package file

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestShouldThrowExceptionWhenQuestionsAreLowerThanZero(t *testing.T) {
	_, error := ReadFile(-1)
	assert.EqualError(t, error, "Pick at least 1 question","Should be greater than 0")
}

func TestShouldPassWhenValidNumberIsProvided(t *testing.T) {
	res, error := ReadFile(3)
	assert.NoError(t, error)
	assert.IsType(t, []string{}, res)
}