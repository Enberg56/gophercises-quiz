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
		questions = append(questions, qna[0])
		answers = append(answers, qna[1])
	}

	score := 0
	for index, question := range questions {
		fmt.Printf("%s \n", question)
		var input string
		fmt.Scanln(&input)
		if input == answers[index] {
			println("Correct")
			score++
		} else {
			println("Wrong")
		}
	}
	fmt.Printf("You got %d points of %d", score, len(answers))
}
