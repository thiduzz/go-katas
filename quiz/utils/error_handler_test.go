package utils

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"

)

func TestShouldPanicIfError(t *testing.T) {

	assert.Panics(t, func() {
		HasError(errors.New("This is an error"))
	}, "Panics when error is validates")
}

func TestShouldNotPanicIfNotError(t *testing.T) {

	assert.NotPanicsf(t, func() {
		HasError(nil)
	}, "Not panics if error is not really an error")
}
