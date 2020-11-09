package main

import (
	"flag"
	"github.com/thiduzz/quiz"
	"github.com/thiduzz/quiz/file"
	"time"
)

func main(){
	var parsedQuestions []quiz.QuestionItem
	questions, seconds, filename := readParams()
	parsedQuestions, _ = file.ReadFile(file.OpenFile(filename), questions)

	newQuiz :=  quiz.Quiz{
		Questions: parsedQuestions,
		Correct: 0,
		Time: time.Second * time.Duration(seconds),
	}
	quiz.Start(&newQuiz)
	timer := time.AfterFunc(newQuiz.Time, func() {
		quiz.ShowResults(&newQuiz)
	})
	defer timer.Stop()

	//TODO: kick-start a timer
	//TODO: observe the timer
	//TODO: for each question
	//	- Print
	//	- Expect answer
	//	- Compare answer - store corrects in a separate array
	//TODO: no more questions to show? Show results
	//TODO: no more time? Show results
}

func readParams() (int, int, string) {

	var questions int
	var seconds int
	var filename string
	flag.IntVar(&questions, "questions",20,"Define number of questions")
	flag.IntVar(&seconds, "seconds",60, "Define test duration")
	flag.StringVar(&filename,"filename","problems.csv","The file that contains the questions")
	flag.Parse()
	return questions, seconds, filename
}
