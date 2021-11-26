package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

type questionsAndAnswers struct {
	Question string
	Answer   string
}

func main() {
	//flag where the user can input a csv file with thier own problems
	csvfile := flag.String("csv", "problems.csv", "add in csv file in the format of `question,answer`")
	//timelimit flag, this will allow users to customize the timelimit for the quiz
	timeLimit := flag.Int("limit", 30, "time limit for the quiz in seconds")
	//shuffle flag that where the user can configure if they want to shuffle the problems or not
	shuffleQuiz := flag.Bool("shuffle", false, "Shuffles the quiz problems")
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

	problems := parseQAndA(reader, *shuffleQuiz)
	//using the time package, this sets up the timer that will run the test for n seconds
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	correctAnswers := 0

	for i, problem := range problems {
		fmt.Printf("Problem # %d: %s = ", i+1, problem.Question)
		//channel that will transfer data to be used later is the select statement
		answerChannel := make(chan string)
		//goroutine that take the user's input, and add their answer to the answer channel
		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerChannel <- answer
		}()
		select {
		//if we get a response from the timer channel, then end the loop and return the score
		case <-timer.C:
			fmt.Printf("\nyour score is %d/%d\n", correctAnswers, len(problems))
			return
		//if we get a response from the answer channel, run the check to see if the answer the user input matches that of the answer to the problem
		case answer := <-answerChannel:
			if answer == problem.Answer {
				correctAnswers += 1
			}
		}
	}
	fmt.Printf("your score is %d/%d\n", correctAnswers, len(problems))
}

//function to parse the questions and answers into a slice that will hold each question and answer struct
func parseQAndA(lines [][]string, shuffleQuiz bool) []questionsAndAnswers {
	//create a slice that will create new instances of the questionsAndAnswers struct that will be the length of the lines
	parsedProblems := make([]questionsAndAnswers, len(lines))
	for i, line := range lines {
		parsedProblems[i] = questionsAndAnswers{
			Question: line[0],
			Answer:   strings.TrimSpace(line[1]),
		}
	}
	if shuffleQuiz != false {
		rand.Shuffle(len(parsedProblems), func(i, j int) {
			parsedProblems[i], parsedProblems[j] = parsedProblems[j], parsedProblems[i]
		})
	}
	return parsedProblems
}

//If there is an error, log it and return an exist status of 1
func returnError(err error) {
	log.Fatal(err)
	os.Exit(1)
}
