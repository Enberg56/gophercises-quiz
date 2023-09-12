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

	for _, question := range questions {
		fmt.Print(question)
		var input string
		fmt.Scan(&input)
		fmt.Print("You answered %s\n", &input)

	}
}
