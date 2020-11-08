package main

import (
	"flag"
	"fmt"
)

func main(){
	//var parsedQuestions []string
	questions, _, _ := readParams()
	//parsedQuestions, _ = file.ReadFile(questions)
	fmt.Println(questions)
	//TODO: open csv
	//TODO: pick only the amount of questions specified (if there is none, consider default)
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
