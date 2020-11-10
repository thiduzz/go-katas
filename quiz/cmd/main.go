package main

import (
	"flag"
	"fmt"
	"github.com/thiduzz/quiz"
	"github.com/thiduzz/quiz/file"
	"github.com/thiduzz/quiz/utils"
	"os"
	"time"
)

func main(){
	var parsedQuestions []quiz.QuestionItem
	questions, seconds, filename := readParams()
	parsedQuestions, err := file.ReadFile(file.OpenFile(filename), questions)
	utils.HasError(err)
	newQuiz :=  quiz.Quiz{
		Questions: parsedQuestions,
		Correct: 0,
		Time: time.Second * time.Duration(seconds),
	}
	quiz.Start(&newQuiz)
	timer := time.NewTimer(newQuiz.Time)
	go func() {
		<-timer.C
		fmt.Println("Time finished")
		quiz.ShowResults(&newQuiz)
		os.Exit(0)
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
