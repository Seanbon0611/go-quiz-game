package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

type questionsAndAnswers struct {
	Question string
	Answer   string
}

func main() {
	//opens the csv file to be able to read
	file, err := os.Open("./problems.csv")
	if err != nil {
		log.Fatal(err)
	}
	//what gets opened must be closed, don't forget!
	defer file.Close()
	//turns contents of the csv file into somethijng readable rather than just a pointer to the csv file contents
	reader, _ := csv.NewReader(file).ReadAll()
	//loop through the reader and create new instances of the questionsAndAnswers struct and print the results
	for _, line := range reader {
		qAndA := questionsAndAnswers{
			Question: line[0],
			Answer:   line[1],
		}
		fmt.Println(qAndA.Question + " " + qAndA.Answer)
	}
}
