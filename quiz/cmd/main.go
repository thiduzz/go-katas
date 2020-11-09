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
	timer := time.NewTimer(newQuiz.Time)
	go func() {
		<-timer.C
		quiz.ShowResults(&newQuiz)
	}()
	time.Sleep(newQuiz.Time)
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
