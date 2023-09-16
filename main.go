package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func main() {
	csvFileName := flag.String("csv", "problems.csv", "A csv file with problems to solve")
	timeLimit := flag.Int("time limit", 10, "Timelimit in seconds")
	flag.Parse()

	var questions []string
	var answers []string
	dat, err := os.Open(*csvFileName)
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

	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)
	score := 0
	fmt.Printf("Timelimit set to %d seconds. Good luck.\n", *timeLimit)
	for index, question := range questions {
		fmt.Printf("%s \n", question)
		answerCh := make(chan string)
		go func() {
			var input string
			fmt.Scanln(&input)
			answerCh <- input
		}()
		select {
		case <-timer.C:
			fmt.Printf("\nYou got %d points of %d", score, len(answers))
			return
		case input := <-answerCh:
			if input == answers[index] {
				score++
			}
		}
	}
	fmt.Printf("You got %d points of %d", score, len(answers))
}
