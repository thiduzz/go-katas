package file

import (
	"encoding/csv"
	"errors"
	"github.com/thiduzz/quiz/utils"
	"io"
	"log"
	"os"
)


type QuestionItem struct {
	Question string
	Answer string
}

func ReadFile(filename string, questions int) ([]QuestionItem, error) {

	if(questions <= 0){
		return nil, errors.New("Pick at least 1 question")
	}

	var questionsArray []QuestionItem
	file, err := os.Open(filename)
	utils.HasError(err)
	rows := csv.NewReader(file)

	for i := 0; i < questions; i++ {
		record, err := rows.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		questionsArray = append(questionsArray, QuestionItem{
			Question: record[0],
			Answer: record[1],
		})
	}

	return questionsArray, nil
}

type Reader struct {

}