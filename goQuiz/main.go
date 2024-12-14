package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	csvFileName := flag.String("csv", "problems.csv", "a file with qna")
	timeLimit := flag.Int("limit", 30, "the time limit in secs")

	flag.Parse()

	file, err := os.Open(*csvFileName)

	if err != nil {
		exit(fmt.Sprintf("failed to open the CSV file %s", *csvFileName))

	}

	r := csv.NewReader(file)
	lines, err := r.ReadAll()

	if err != nil {
		exit("failed to parse the provided csv file")
	}

	problems := parseLines(lines)
	fmt.Println(problems)

	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	var count int = 0
 
	for i, p := range problems {

		fmt.Printf("problem #%d :  %s = \n", i+1, p.q)
		answerCh := make(chan string)

		go func() {
			var answer string
			fmt.Scanf("%s", &answer)
			answerCh <- answer

		}()

		select {
		case <-timer.C:
			fmt.Printf("\n correct answers are : %d/%d", count, len(problems))
			return
		case answer := <-answerCh:
			if answer == p.a {

				count += 1
			}
		default:

		}
	}
	fmt.Printf("\n correct answers are : %d/%d", count, len(problems))
}

func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))

	for i, line := range lines {
		ret[i] = problem{
			q: strings.TrimSpace(line[0]),
			a: strings.TrimSpace(line[1]),
		}
	}

	return ret
}

type problem struct {
	q string
	a string
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
