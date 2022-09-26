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

func readCSV(path string) [][]string {
	f, err := os.Open(path)
	check(err)
	csv := csv.NewReader(f)
	records, err := csv.ReadAll()
	check(err)

	return records
}

func main() {
	path := parseArgs()
	records := readCSV(path)

	numCorrect := 0
	for i, row := range records {
		question := row[0]
		answer := row[1]

		fmt.Printf("Question #%d: %s\n", i+1, question)

		var submission string
		_, err := fmt.Scanln(&submission)
		check(err)

		if submission == answer {
			fmt.Println("Correct!")
			numCorrect++
		} else {
			fmt.Println("Wrong!")
		}
	}

	fmt.Printf("Finished! You got %d/%d correct!\n", numCorrect, len(records))
}
