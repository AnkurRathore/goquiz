package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	csvFilename := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")

	flag.Parse()
	_ = csvFilename

	file, err := os.Open(*csvFilename)

	if err != nil {
		exit(fmt.Sprintf("Failed to open the CSV file:%s\n", *csvFilename))

	}
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit("Failed to parse the provided CSV file")
	}
	correct := 0
	problems := parseLines(lines)
	for i, p := range problems {
		fmt.Printf("Problem #%d:%s = \n", i+1, p.question)

		var answer string
		fmt.Scanf("%s\n", &answer)
		if answer == p.answer {
			correct++
		}
	}

	fmt.Printf("You Scored %d out of %d.\n", correct, len(problems))
}

// reading the lines and creating and array of the problem struct
func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))
	for i, line := range lines {
		ret[i] = problem{
			question: line[0],
			answer:   strings.TrimSpace(line[1]),
		}
	}

	return ret
}

// defining a a type to hold the question and answer
type problem struct {
	question string
	answer   string
}

// common function for handling errors
func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
