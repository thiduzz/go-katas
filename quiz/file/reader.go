package file

import (
	"encoding/csv"
	"errors"
	"github.com/thiduzz/quiz"
	"github.com/thiduzz/quiz/utils"
	"io"
	"os"
	"strings"
)

func OpenFile(filename string) *csv.Reader {
	file, err := os.Open(filename)
	utils.HasError(err)

	return csv.NewReader(file)
}

func ReadFile(reader *csv.Reader, questions int) ([]quiz.QuestionItem, error) {

	var questionsArray []quiz.QuestionItem
	if(questions <= 0){
		return nil, errors.New("Pick at least 1 question")
	}

	for i := 0; i < questions; i++ {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		utils.HasError(err)
		questionsArray = append(questionsArray, quiz.QuestionItem{
			Question: record[0],
			Answer: strings.TrimSpace(record[1]),
		})
	}

	return questionsArray, nil
}