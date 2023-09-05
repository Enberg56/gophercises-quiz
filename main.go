package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	var questions []string
	var answers []string
	dat, err := os.Open("problems.csv")
	check(err)
	defer dat.Close()

	scanner := bufio.NewScanner(dat)

	for scanner.Scan() {
		line := scanner.Text()
		qna := strings.Split(line, ",")
		if len(qna) != 2 {
			continue
		}
		question := string(qna[0])
		answer := string(qna[1])
		questions = append(questions, question)
		answers = append(answers, answer)
	}
	fmt.Println("Questions:", questions)
	fmt.Println("Answers:", answers)
}
