package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func parseArgs() (string, int) {
	cwd, err := os.Getwd()
	check(err)
	defaultFile := filepath.Join(cwd, "problems.csv")

	path := flag.String("file", defaultFile, "Path to CSV containing quiz problems.")
	timeLimit := flag.Int("time", 30, "Time limit for the quiz.")

	flag.Parse()
	return *path, *timeLimit
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

func runQuiz(problems []problem, timeLimit int) int {
	timer := time.NewTimer(time.Duration(timeLimit) * time.Second)

	correct := make(chan int)
	numCorrect := 0
	for i, row := range problems {
		go askProblem(row, i, correct)
		select {
		case msg := <-correct:
			numCorrect += msg
		case <-timer.C:
			return numCorrect
		}
	}
	return numCorrect
}

func main() {
	path, timeLimit := parseArgs()
	problems := readCSV(path)

	fmt.Print("Press [ENTER] to start!")
	fmt.Scanln()
	numCorrect := runQuiz(problems, timeLimit)

	fmt.Printf("Finished! You got %d/%d correct!\n", numCorrect, len(problems))
}
