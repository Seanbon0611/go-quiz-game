package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
)

type questionsAndAnswers struct {
	Question string
	Answer   string
}

func main() {
	csvfile := flag.String("csv", "problems.csv", "A quiz game CLI application that takes in csv file in the format of `question,answer`")

	flag.Parse()

	//opens the csv file to be able to read
	file, err := os.Open(*csvfile)
	if err != nil {
		returnError(err)
	}
	//what gets opened must be closed, don't forget!
	defer file.Close()
	//turns contents of the csv file into something readable rather than just a pointer to the csv file contents
	reader, _ := csv.NewReader(file).ReadAll()
	problems := parseQAndA(reader)
	//keeps track of how many answers the user gets correct
	correctAnswers := 0
	//loop through each problem within the slice of problems, give the question and have the customer submit the answer
	for i, problem := range problems {
		fmt.Printf("Problem # %d: %s = \n", i+1, problem.Question)
		var answer string
		fmt.Scanf("%s\n", &answer)
		if answer == problem.Answer {
			correctAnswers += 1
		}
	}
	fmt.Printf("your score is %d/%d\n", correctAnswers, len(problems))
}

//function to parse the questions and answers into a slice that will hold each question and answer struct
func parseQAndA(lines [][]string) []questionsAndAnswers {
	//create a slice that will create new instances of the questionsAndAnswers struct that will be the length of the lines
	parsedProblems := make([]questionsAndAnswers, len(lines))
	for i, line := range lines {
		parsedProblems[i] = questionsAndAnswers{
			Question: line[0],
			Answer:   line[1],
		}
	}
	return parsedProblems
}

//If there is an error, log it and return an exist status of 1
func returnError(err error) {
	log.Fatal(err)
	os.Exit(1)
}
