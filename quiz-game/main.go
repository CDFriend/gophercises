package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

type Question struct {
	Question string
	Answer   string
}

var inputFile string
var curQuestion Question

// Checks for an error condition, exits with a message if one is found.
func check(e error) {
	if e != nil {
		panic(e)
	}
}

// NextQuestion gets the next question from a csv file.
// If no more questions are available, it returns true.
// If more questions are available, it returns false and writes question data to q.
func NextQuestion(r *csv.Reader, q *Question) bool {
	line, err := r.Read()
	if err == io.EOF {
		return true
	}
	check(err)

	q.Question = line[0]
	q.Answer = line[1]

	return false
}

// CompareAnswers performs a case-insensitive comparison of two string answers.
// It ignores any extra whitespace at the beginning and end of an answer.
// It returns true if answer1 is the same as answer2.
func CompareAnswers(answer1 string, answer2 string) bool {
	ans1Normalized := strings.ToLower(answer1)
	ans1Normalized = strings.TrimSpace(ans1Normalized)

	ans2Normalized := strings.ToLower(answer2)
	ans2Normalized = strings.TrimSpace(ans2Normalized)

	compareResult := strings.Compare(ans1Normalized, ans2Normalized)
	return compareResult == 0
}

func main() {
	flag.StringVar(&inputFile, "p", "problems.csv", "Comma-separated problems file")
	flag.Parse()

	reader, err := os.Open(inputFile)
	check(err)
	csvReader := csv.NewReader(reader)

	questionsCorrect := 0
	questionsIncorrect := 0

	stdinReader := bufio.NewReader(os.Stdin)
	for NextQuestion(csvReader, &curQuestion) != true {

		fmt.Print(curQuestion.Question + "? ")

		answerIn, err := stdinReader.ReadString('\n')
		check(err)

		if CompareAnswers(answerIn, curQuestion.Answer) {
			questionsCorrect++
			fmt.Printf("Correct! You answered %d questions correctly, %d questions incorrectly\n",
				questionsCorrect, questionsIncorrect)
		} else {
			questionsIncorrect++
			fmt.Printf("Wrong! You answered %d questions correctly, %d questions incorrectly\n",
				questionsCorrect, questionsIncorrect)

			fmt.Printf("Correct answer was: %s\n", curQuestion.Answer)
		}
	}
}
