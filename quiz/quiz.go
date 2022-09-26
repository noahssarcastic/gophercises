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

func main() {
	path := parseArgs()
	problems := readCSV(path)

	numCorrect := 0
	for i, row := range problems {
		fmt.Printf("Question #%d: %s\n", i+1, row.question)

		var submission string
		_, err := fmt.Scanf("%s\n", &submission)
		if err != nil && err.Error() == "unexpected newline" {
			fmt.Println("Skipped!")
			continue
		} else {
			check(err)
		}

		if submission == row.answer {
			fmt.Println("Correct!")
			numCorrect++
		} else {
			fmt.Println("Wrong!")
		}
	}

	fmt.Printf("Finished! You got %d/%d correct!\n", numCorrect, len(problems))
}
