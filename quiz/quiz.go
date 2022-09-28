package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func parseArgs() string {
	cwd, err := os.Getwd()
	check(err)
	defaultFile := filepath.Join(cwd, "problems.csv")
	path := flag.String("file", defaultFile, "Path to CSV containing quiz problems.")
	flag.Parse()
	return *path
}

func readCSV(path string) []problem {
	f, err := os.Open(path)
	check(err)
	csv := csv.NewReader(f)
	records, err := csv.ReadAll()
	check(err)

	problems := make([]problem, len(records))
	for i, row := range records {
		problems[i] = problem{
			question: row[0],
			answer:   row[1],
		}
	}
	return problems
}

type problem struct {
	question string
	answer   string
}

func askProblem(p problem, n int, correct chan int) {
	fmt.Printf("Question #%d: %s\n", n+1, p.question)

	var submission string
	_, err := fmt.Scanf("%s\n", &submission)
	if err != nil && err.Error() == "unexpected newline" {
		fmt.Println("Skipped!")
		correct <- 0
		return
	} else {
		check(err)
	}

	if submission == p.answer {
		fmt.Println("Correct!")
		correct <- 1
		return
	} else {
		fmt.Println("Wrong!")
		correct <- 0
		return
	}
}

func runQuiz(problems []problem) int {
	correct := make(chan int)
	quit := make(chan bool)

	numCorrect := 0
	for i, row := range problems {
		go askProblem(row, i, correct)
		select {
		case msg := <-correct:
			numCorrect += msg
		case <-quit:
			return numCorrect
		}
	}
	return numCorrect
}

func main() {
	path := parseArgs()
	problems := readCSV(path)
	numCorrect := runQuiz(problems)
	fmt.Printf("Finished! You got %d/%d correct!\n", numCorrect, len(problems))
}
