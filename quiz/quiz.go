package quiz

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)


type QuestionItem struct {
	Question string
	Answer string
}


type Quiz struct {
	Questions []QuestionItem
	Correct int
	Time time.Duration
}

func Start(currentQuiz *Quiz) {
	reader := bufio.NewReader(os.Stdin)
	for index, question := range currentQuiz.Questions {
		if printAndReadAnswer(reader, index, question) {
			currentQuiz.Correct++
		}
	}
	ShowResults(currentQuiz)
}

func ShowResults(currentQuiz *Quiz) {
	fmt.Println(fmt.Sprintf("Correct answers: %d", currentQuiz.Correct))
}

func printAndReadAnswer(reader *bufio.Reader, index int, question QuestionItem) bool {
	fmt.Println(fmt.Sprintf("Question %d: %s", index, question.Question))
	text, _ := reader.ReadString('\n')
	if strings.TrimSuffix(text,"\n") == question.Answer {
		return true
	}
	return false
}