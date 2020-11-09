package file

import (
	csv "encoding/csv"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func mockFile(entries int) *csv.Reader {
	in := `
		2+2,4
		3+3,6
		4*4,16
`
	return csv.NewReader(strings.NewReader(in))
}



func TestShouldPanicWhenOpeningThatIsIsNotAvailable(t *testing.T) {
	assert.Panics(t, func() {
		OpenFile("dontexist.csv")
	})
}

func TestShouldReturnAReaderWhenIsValidFile(t *testing.T) {
	reader := OpenFile("../problems.csv")
	assert.IsType(t, &csv.Reader{}, reader)
}

func TestShouldThrowExceptionWhenQuestionsAreLowerThanZero(t *testing.T) {
	_, err := ReadFile(mockFile(3), -1)
	assert.EqualError(t, err, "Pick at least 1 question","Should be greater than 0")
}

func TestShouldPassWhenValidNumberIsProvided(t *testing.T) {
	res, err := ReadFile(mockFile(3), 3)
	assert.NoError(t, err)
	assert.IsType(t, []QuestionItem{}, res)
}

func TestShouldReturnCorrectAmountOfQuestionsSpecified(t *testing.T) {
	res, _ := ReadFile(mockFile(3), 3)
	assert.Len(t, res, 3)
}